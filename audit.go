package requestwrapper

import "net/http"

func Audit(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) Response {
		response := next(w, req)
		return response
	}
}
