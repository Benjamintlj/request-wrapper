package requestwrapper_test

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	rw "github.com/Benjamintlj/request-wrapper"
)

func TestAuditRequest(t *testing.T) {
	var logs bytes.Buffer
	oldLogger := slog.Default()
	slog.SetDefault(slog.New(slog.NewJSONHandler(&logs, nil)))
	t.Cleanup(func() { slog.SetDefault(oldLogger) })

	handler := rw.New(rw.CorrelationID, rw.Audit).Handler(func(w rw.ResponseWriter, r *http.Request) rw.Response {
		if !strings.Contains(logs.String(), "HTTP request") {
			t.Error("request was not logged before the handler")
		}
		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) != "hello" {
			t.Errorf("handler body = %q, error = %v", body, err)
		}
		return rw.OK()
	})
	req := httptest.NewRequest(http.MethodPost, "/example?x=1", strings.NewReader("hello"))
	req.Header.Set(rw.CorrelationIDHeader, "test-id")
	req.Header.Set("X-Example", "value")
	handler.ServeHTTP(httptest.NewRecorder(), req)

	for _, value := range []string{"test-id", "POST", "/example?x=1", "X-Example", "value", "hello"} {
		if !strings.Contains(logs.String(), value) {
			t.Errorf("request log missing %q", value)
		}
	}
}
