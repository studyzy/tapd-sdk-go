package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestAddLabel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/label" {
			t.Errorf("unexpected path: %s, want /label", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"LabelPool":{"id":"501","workspace_id":"1","name":"backend","color":"1"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	label, err := c.AddLabel(context.Background(), &model.AddLabelRequest{
		WorkspaceID: "1",
		Name:        "backend",
		Color:       "1",
	})
	if err != nil {
		t.Fatalf("AddLabel() unexpected error: %v", err)
	}
	if label.ID != "501" {
		t.Errorf("label id = %q, want %q", label.ID, "501")
	}
	if label.Name != "backend" {
		t.Errorf("label name = %q, want %q", label.Name, "backend")
	}
}

func TestQueryLabels(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/label" {
			t.Errorf("unexpected path: %s, want /label", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"LabelPool":{"id":"501","workspace_id":"1","name":"backend","color":"1"}},{"LabelPool":{"id":"502","workspace_id":"1","name":"frontend","color":"2"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	labels, err := c.QueryLabels(context.Background(), &model.QueryLabelRequest{
		WorkspaceID: "1",
	})
	if err != nil {
		t.Fatalf("QueryLabels() unexpected error: %v", err)
	}
	if len(labels) != 2 {
		t.Fatalf("expected 2 labels, got %d", len(labels))
	}
	if labels[0].ID != "501" {
		t.Errorf("labels[0].id = %q, want %q", labels[0].ID, "501")
	}
	if labels[1].Name != "frontend" {
		t.Errorf("labels[1].name = %q, want %q", labels[1].Name, "frontend")
	}
}

func TestCountLabels(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/label/count" {
			t.Errorf("unexpected path: %s, want /label/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":5},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountLabels(context.Background(), &model.CountLabelRequest{
		WorkspaceID: "1",
	})
	if err != nil {
		t.Fatalf("CountLabels() unexpected error: %v", err)
	}
	if count != 5 {
		t.Errorf("count = %d, want 5", count)
	}
}
