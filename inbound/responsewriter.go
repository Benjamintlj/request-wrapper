package inbound

import (
	"bytes"
	"net/http"
	"sync"
)

// ResponseWriter provides the standard HTTP response writer methods and access
// to the response body.
type ResponseWriter interface {
	// Header returns the header map that will be sent by WriteHeader.
	// The header map also allows handlers to set HTTP trailers.
	Header() http.Header

	// Write writes data as part of an HTTP reply. If WriteHeader has not yet
	// been called, Write implicitly sets the status to http.StatusOK.
	Write([]byte) (int, error)

	// WriteHeader sets the HTTP response status code. If it is not called
	// explicitly, the first call to Write sets the status to http.StatusOK.
	WriteHeader(statusCode int)

	// Body returns the response body written so far.
	Body() []byte
}

type responseWriter struct {
	writer http.ResponseWriter
	body   *bytes.Buffer
}

var bodyPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

// newResponseWriter wraps the HTTP writer and obtains a buffer from the pool.
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{writer: w, body: bodyPool.Get().(*bytes.Buffer)}
}

func (w *responseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *responseWriter) Write(p []byte) (int, error) {
	return w.body.Write(p)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.writer.WriteHeader(statusCode)
}

func (w *responseWriter) Body() []byte {
	return w.body.Bytes()
}

// Close resets the body buffer and returns it to the pool. Body is invalid after Close.
func (w *responseWriter) Close() error {
	if w.body != nil {
		w.body.Reset()
		bodyPool.Put(w.body)
		w.body = nil
	}
	return nil
}
