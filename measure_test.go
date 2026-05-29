package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestGetLifeTimes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/life_times" {
			t.Errorf("unexpected path: %s, want /life_times", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "11111111" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "11111111")
		}
		if r.URL.Query().Get("entity_type") != "story" {
			t.Errorf("entity_type = %q, want %q", r.URL.Query().Get("entity_type"), "story")
		}
		if r.URL.Query().Get("entity_id") != "2001" {
			t.Errorf("entity_id = %q, want %q", r.URL.Query().Get("entity_id"), "2001")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"LifeTime":{"id":"1001","workspace_id":"11111111","entity_id":"2001","status":"open","begin_date":"2026-01-01 10:00:00","end_date":"2026-01-02 10:00:00","time_cost":"86400"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	data, err := c.GetLifeTimes(context.Background(), &model.GetLifeTimesRequest{
		WorkspaceID: "11111111",
		EntityType:  "story",
		EntityID:    "2001",
	})
	if err != nil {
		t.Fatalf("GetLifeTimes() unexpected error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil data")
	}
	if string(data) == "" {
		t.Error("expected non-empty data")
	}
}

func TestGetLifeTimes_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":0,"data":null,"info":"permission denied"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.GetLifeTimes(context.Background(), &model.GetLifeTimesRequest{
		WorkspaceID: "11111111",
		EntityType:  "story",
		EntityID:    "2001",
	})
	if err == nil {
		t.Fatal("expected error for status=0")
	}
}
