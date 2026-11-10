package requestwrapper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseWriter(t *testing.T) {
	recorder := httptest.NewRecorder()
	writer := newResponseWriter(recorder)
	defer writer.Close()

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	if _, err := writer.Write([]byte("hello")); err != nil {
		t.Fatal(err)
	}
	if _, err := writer.Write([]byte(" world")); err != nil {
		t.Fatal(err)
	}

	if got := recorder.Header().Get("Content-Type"); got != "text/plain" {
		t.Fatalf("header = %q, want text/plain", got)
	}
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusCreated)
	}
	if got := string(writer.Body()); got != "hello world" {
		t.Fatalf("body = %q, want hello world", got)
	}
	if recorder.Body.Len() != 0 {
		t.Fatal("body was written before the handler finished")
	}
}

func TestResponseWriterClose(t *testing.T) {
	writer := newResponseWriter(httptest.NewRecorder())
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
}
