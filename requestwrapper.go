package requestwrapper

import (
	"errors"
	"net/http"
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
func New(middleware ...Middleware) *Wrapper {
	return &Wrapper{middleware: middleware}
}

// Handler applies the middleware and adapts a HandlerFunc to http.Handler.
func (wrapper *Wrapper) Handler(handler HandlerFunc) http.Handler {
	wrapped := handler
	for _, v := range slices.Backward(wrapper.middleware) {
		wrapped = v(wrapped)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped(w, r)
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
