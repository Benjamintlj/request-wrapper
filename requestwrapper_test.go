package requestwrapper_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
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
	server := httptest.NewServer(mux)
	defer server.Close()

	response, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.StatusCode)
	}
}

func TestNewFatalsWhenBuiltInMiddlewareIsOutOfOrder(t *testing.T) {
	custom := func(next rw.HandlerFunc) rw.HandlerFunc { return next }
	if os.Getenv("RW_TEST_FATAL_ORDER") == "1" {
		rw.New(rw.Audit, custom, rw.ContextID)
		return
	}

	rw.New(rw.ContextID, custom, rw.Audit, rw.ErrorCapture)
	command := exec.Command(os.Args[0], "-test.run=^TestNewFatalsWhenBuiltInMiddlewareIsOutOfOrder$")
	command.Env = append(os.Environ(), "RW_TEST_FATAL_ORDER=1")
	output, err := command.CombinedOutput()
	exit, ok := err.(*exec.ExitError)
	if !ok || exit.ExitCode() != 1 {
		t.Fatalf("expected fatal exit, got %v: %s", err, output)
	}
	if !strings.Contains(string(output), "middleware out of order: expected position 0 follows 1") {
		t.Fatalf("expected out-of-order message, got %q", output)
	}
}

func TestHandlerUsesReturnedStatus(t *testing.T) {
	r := rw.New()
	server := httptest.NewServer(r.Handler(func(w http.ResponseWriter, req *http.Request) rw.Response {
		return rw.Created()
	}))
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, response.StatusCode)
	}
}
