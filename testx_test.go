package tapd

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestDoTestxRequest_HTTPError(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`internal server error`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.doTestxGet(context.Background(), "/api/testx/case/v1/test", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var tapErr *TAPDError
	if !errors.As(err, &tapErr) {
		t.Fatalf("expected *TAPDError, got %T: %v", err, err)
	}
	if tapErr.HTTPStatus != 500 {
		t.Errorf("got HTTPStatus=%d, want 500", tapErr.HTTPStatus)
	}
}

func TestDoTestxRequest_APIError(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":{"Code":"ERR","Message":"test error"},"Data":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.doTestxGet(context.Background(), "/api/testx/case/v1/test", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var tapErr *TAPDError
	if !errors.As(err, &tapErr) {
		t.Fatalf("expected *TAPDError, got %T: %v", err, err)
	}
	if tapErr.ExitCode != 4 {
		t.Errorf("got ExitCode=%d, want 4", tapErr.ExitCode)
	}
}
