package requestwrapper_test

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	rw "github.com/Benjamintlj/request-wrapper"
)

func TestErrorCapture(t *testing.T) {
	var logs bytes.Buffer
	previous := slog.Default()
	slog.SetDefault(slog.New(slog.NewJSONHandler(&logs, nil)))
	t.Cleanup(func() { slog.SetDefault(previous) })

	handler := rw.New(rw.ErrorCapture).Handler(func(w rw.ResponseWriter, r *http.Request) rw.Response {
		return rw.InternalServerError(errors.New("example failure"))
	})
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", recorder.Code, http.StatusInternalServerError)
	}
	if !strings.Contains(logs.String(), `"level":"ERROR"`) || !strings.Contains(logs.String(), "example failure") {
		t.Errorf("expected error log, got %s", logs.String())
	}
}
