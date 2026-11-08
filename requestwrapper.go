package requestwrapper

import (
	"context"
	"errors"
	"log"
	"net/http"
	"reflect"
	"slices"
)

// HandlerFunc handles an HTTP request and returns a response for the wrapper.
type HandlerFunc func(http.ResponseWriter, *http.Request) Response

// Middleware wraps a HandlerFunc with request or response behavior.
type Middleware func(HandlerFunc) HandlerFunc

// Wrapper composes middleware and manages its HTTP server.
type Wrapper struct {
	middleware []Middleware
	server     *http.Server
	done       chan struct{}
}

// New creates a wrapper with middleware applied in the order provided.
// The first middleware is the outermost wrapper and runs first on a request.
// It exits the process when built-in middleware is supplied out of order.
func New(middleware ...Middleware) *Wrapper {
	expected := []Middleware{
		CorrelationID,
		Audit,
		ErrorCapture,
	}

	positions := make(map[uintptr]int, len(expected))
	for i, entry := range expected {
		positions[getPointer(entry)] = i
	}

	previous := -1
	for _, entry := range middleware {
		if entry == nil {
			continue
		}

		position, builtIn := positions[getPointer(entry)]
		if !builtIn {
			continue
		}

		if position <= previous {
			log.Fatal("middleware out of order")
		}

		previous = position
	}

	return &Wrapper{middleware: middleware}
}

// Handler applies the middleware and adapts a HandlerFunc to http.Handler.
func (w *Wrapper) Handler(handler HandlerFunc) http.Handler {
	wrapped := handler
	for _, v := range slices.Backward(w.middleware) {
		wrapped = v(wrapped)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()
		r = r.WithContext(ctx)
		response := wrapped(w, r)
		w.WriteHeader(response.StatusCode)
	})
}

// ListenAndServe starts an HTTP server in a goroutine and returns its error channel.
// The channel closes when the server stops; normal shutdown does not send an error.
func (w *Wrapper) ListenAndServe(address string, handler http.Handler) <-chan error {
	w.server = &http.Server{Addr: address, Handler: handler}
	errorsCh := make(chan error, 1)
	w.done = make(chan struct{}, 1)

	go func() {
		defer close(w.done)
		defer close(errorsCh)
		if err := w.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorsCh <- err
		}
	}()

	return errorsCh
}

// Close stops the HTTP server if it has been started.
func (w *Wrapper) Close() {
	if w.server != nil {
		_ = w.server.Close()
	}

	<-w.done
}

func getPointer(middleware Middleware) uintptr {
	return reflect.ValueOf(middleware).Pointer()
}
