package requestwrapper_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	rw "github.com/Benjamintlj/request-wrapper"
)

func TestCorrelationID(t *testing.T) {
	for _, test := range []struct {
		name   string
		header string
	}{
		{name: "existing", header: "existing-correlation-id"},
		{name: "generated"},
	} {
		t.Run(test.name, func(t *testing.T) {
			var got string
			var requestID string
			handler := rw.CorrelationID(func(w rw.ResponseWriter, r *http.Request) rw.Response {
				got, _ = rw.CorrelationIDFromContext(r.Context())
				requestID = r.Header.Get(rw.CorrelationIDHeader)
				return rw.OK()
			})
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.header != "" {
				req.Header.Set(rw.CorrelationIDHeader, test.header)
			}
			recorder := httptest.NewRecorder()
			rw.New().Handler(handler).ServeHTTP(recorder, req)
			if test.header != "" && got != test.header {
				t.Fatalf("expected correlation ID %q, got %q", test.header, got)
			}
			if test.header == "" && !regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`).MatchString(got) {
				t.Fatalf("expected generated UUID, got %q", got)
			}
			if requestID != test.header {
				t.Errorf("request header = %q, want %q", requestID, test.header)
			}
			if responseID := recorder.Result().Header.Get(rw.CorrelationIDHeader); responseID != got {
				t.Errorf("response header = %q, want %q", responseID, got)
			}
		})
	}
}
