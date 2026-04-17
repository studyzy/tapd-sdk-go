package tapd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestCreateRelease(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/releases" {
			t.Errorf("unexpected path: %s, want /releases", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		if bodyStr == "" {
			t.Fatal("expected non-empty body")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Release":{"id":"1111111111001000001","workspace_id":"11111111","name":"v1.0","description":"first release","startdate":"2026-01-01","enddate":"2026-02-01","status":"open","created":"2026-01-01 10:00:00","modified":"2026-01-01 10:00:00"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	release, err := c.CreateRelease(&model.CreateReleaseRequest{
		WorkspaceID: "11111111",
		Name:        "v1.0",
		Description: "first release",
		StartDate:   "2026-01-01",
		EndDate:     "2026-02-01",
	})
	if err != nil {
		t.Fatalf("CreateRelease() unexpected error: %v", err)
	}
	if release.ID != "1111111111001000001" {
		t.Errorf("release id = %q, want %q", release.ID, "1111111111001000001")
	}
	if release.Name != "v1.0" {
		t.Errorf("release name = %q, want %q", release.Name, "v1.0")
	}
	if release.Description != "first release" {
		t.Errorf("release description = %q, want %q", release.Description, "first release")
	}
	if release.Status != "open" {
		t.Errorf("release status = %q, want %q", release.Status, "open")
	}
}

func TestUpdateRelease(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/releases" {
			t.Errorf("unexpected path: %s, want /releases", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Release":{"id":"1111111111001000001","workspace_id":"11111111","name":"v1.1","status":"done"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	release, err := c.UpdateRelease(&model.UpdateReleaseRequest{
		WorkspaceID: "11111111",
		ID:          "1111111111001000001",
		Name:        "v1.1",
		Status:      "done",
	})
	if err != nil {
		t.Fatalf("UpdateRelease() unexpected error: %v", err)
	}
	if release.ID != "1111111111001000001" {
		t.Errorf("release id = %q, want %q", release.ID, "1111111111001000001")
	}
	if release.Name != "v1.1" {
		t.Errorf("release name = %q, want %q", release.Name, "v1.1")
	}
	if release.Status != "done" {
		t.Errorf("release status = %q, want %q", release.Status, "done")
	}
}

func TestCountReleases(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/releases/count" {
			t.Errorf("unexpected path: %s, want /releases/count", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "11111111" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "11111111")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":5},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountReleases(&model.WorkspaceIDRequest{WorkspaceID: "11111111"})
	if err != nil {
		t.Fatalf("CountReleases() unexpected error: %v", err)
	}
	if count != 5 {
		t.Errorf("count = %d, want %d", count, 5)
	}
}
