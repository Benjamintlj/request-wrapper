package requestwrapper

import (
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
}

// New creates a wrapper with middleware applied in the order provided.
// The first middleware is the outermost wrapper and runs first on a request.
// It exits the process when built-in middleware is supplied out of order.
func New(middleware ...Middleware) *Wrapper {
	positions := map[uintptr]int{
		reflect.ValueOf(ContextID).Pointer():    0,
		reflect.ValueOf(Audit).Pointer():        1,
		reflect.ValueOf(ErrorCapture).Pointer(): 2,
	}

	previous := -1
	for _, entry := range middleware {
		if entry == nil {
			continue
		}
		position, builtIn := positions[reflect.ValueOf(entry).Pointer()]
		if !builtIn {
			continue
		}
		if position <= previous {
			log.Fatalf("middleware out of order: expected position %d follows %d", position, previous)
		}
		previous = position
	}

	return &Wrapper{middleware: middleware}
}

// Handler applies the middleware and adapts a HandlerFunc to http.Handler.
func (wrapper *Wrapper) Handler(handler HandlerFunc) http.Handler {
	wrapped := handler
	for _, v := range slices.Backward(wrapper.middleware) {
		wrapped = v(wrapped)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := wrapped(w, r)
		w.WriteHeader(response.StatusCode)
	})
}

// ListenAndServe starts an HTTP server in a goroutine and returns its error channel.
// The channel closes when the server stops; normal shutdown does not send an error.
func (w *Wrapper) ListenAndServe(address string, handler http.Handler) <-chan error {
	w.server = &http.Server{Addr: address, Handler: handler}
	errorsCh := make(chan error, 1)

	go func() {
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
}
