package inbound

import (
	"log/slog"
	"net/http"
)

func ErrorCapture(next HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, req *http.Request) Response {
		response := next(w, req)
		if response.Err != nil {
			slog.Error("request handler error", "error", response.Err)
		}
		return response
	}
}
