package requestwrapper_test

import (
	"net/http"
	"testing"

	rw "github.com/Benjamintlj/request-wrapper"
)

func TestHandlerReturnsOK(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) rw.Response {
		return rw.OK()
	}

	r := rw.New(
		rw.ContextID,
		rw.Audit,
		rw.ErrorCapture,
	)
	mux := http.NewServeMux()
	mux.Handle("GET /", r.Handler(handler))
	r.ListenAndServe(":8080", mux)
	defer r.Close()

	response, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.StatusCode)
	}
}
