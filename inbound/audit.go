package inbound

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

func Audit(next HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, req *http.Request) Response {
		var body []byte
		if req.Body != nil {
			var err error
			body, err = io.ReadAll(req.Body)
			req.Body = &auditRequestBody{Reader: bytes.NewReader(body), Closer: req.Body}
			if err != nil {
				slog.Warn("read request body for audit", "error", err)
			}
		}
		slog.Info("HTTP request", "method", req.Method, "url", req.URL.String(), "headers", req.Header, "body", string(body))
		response := next(w, req)
		return response
	}
}

type auditRequestBody struct {
	io.Reader
	io.Closer
}
