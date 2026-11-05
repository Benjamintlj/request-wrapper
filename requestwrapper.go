package requestwrapper

import (
	"errors"
	"net/http"
	"slices"
	"sync"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) Response

type Middleware func(HandlerFunc) HandlerFunc

type Wrapper struct {
	middleware []Middleware
	closeCh    chan struct{}
	closeOnce  sync.Once
}

func New(middleware ...Middleware) *Wrapper {
	return &Wrapper{middleware: middleware, closeCh: make(chan struct{})}
}

func (wrapper *Wrapper) Handler(handler HandlerFunc) http.Handler {
	wrapped := handler
	for _, v := range slices.Backward(wrapper.middleware) {
		wrapped = v(wrapped)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped(w, r)
	})
}

func (w *Wrapper) ListenAndServe(address string, handler http.Handler) <-chan error {
	server := &http.Server{Addr: address, Handler: handler}
	errorsCh := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)
		defer close(errorsCh)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorsCh <- err
		}
	}()

	go func() {
		select {
		case <-w.closeCh:
			_ = server.Close()
		case <-done:
		}
	}()

	return errorsCh
}

func (w *Wrapper) Close() {
	w.closeOnce.Do(func() { close(w.closeCh) })
}
