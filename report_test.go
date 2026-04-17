package tapd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestGetWorkspaceReports(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/reports" {
			t.Errorf("unexpected path: %s, want /reports", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "11111111" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "11111111")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Report":{"id":"1001","workspace_id":"11111111","name":"Weekly Report","created":"2026-01-01"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	data, err := c.GetWorkspaceReports(&model.WorkspaceIDRequest{WorkspaceID: "11111111"})
	if err != nil {
		t.Fatalf("GetWorkspaceReports() unexpected error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil data")
	}
	if string(data) == "" {
		t.Error("expected non-empty data")
	}
}

func TestGetWorkspaceReports_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":0,"data":null,"info":"permission denied"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.GetWorkspaceReports(&model.WorkspaceIDRequest{WorkspaceID: "11111111"})
	if err == nil {
		t.Fatal("expected error for status=0")
	}
}
