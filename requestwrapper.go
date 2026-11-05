package requestwrapper

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request) Response

type Middleware func(HandlerFunc) HandlerFunc

type Wrapper struct{}

func New() *Wrapper {
	return &Wrapper{}
}

func (w *Wrapper) Wrap(middleware ...Middleware) {
	panic("not implemented")
}

func (w *Wrapper) Handler(handler HandlerFunc) http.Handler {
	panic("not implemented")
}

func (w *Wrapper) Start(address string, handler http.Handler) {
	panic("not implemented")
}
