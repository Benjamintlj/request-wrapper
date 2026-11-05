package requestwrapper

import "net/http"

func ContextID(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) Response {
		response := next(w, req)
		return response
	}
}
