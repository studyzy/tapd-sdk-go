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

func TestUpdateLabel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/label" {
			t.Errorf("unexpected path: %s, want /label", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s, want POST", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("ParseForm: %v", err)
		}
		if r.Form.Get("workspace_id") != "20358527" {
			t.Errorf("workspace_id = %q, want 20358527", r.Form.Get("workspace_id"))
		}
		if r.Form.Get("id") != "1220358527000001577" {
			t.Errorf("id = %q, want 1220358527000001577", r.Form.Get("id"))
		}
		if r.Form.Get("color") != "3" {
			t.Errorf("color = %q, want 3", r.Form.Get("color"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"LabelPool":{"id":"1220358527000001577","workspace_id":"20358527","name":"创建标签","color":"3","creator":"","modifier":"","created":"2022-09-26 20:25:02","modified":"2022-09-26 20:25:02"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	label, err := c.UpdateLabel(context.Background(), &model.UpdateLabelRequest{
		WorkspaceID: "20358527",
		ID:          "1220358527000001577",
		Color:       "3",
	})
	if err != nil {
		t.Fatalf("UpdateLabel() unexpected error: %v", err)
	}
	if label.ID != "1220358527000001577" {
		t.Errorf("label id = %q, want %q", label.ID, "1220358527000001577")
	}
	if label.Color != "3" {
		t.Errorf("label color = %q, want %q", label.Color, "3")
	}
	if label.Name != "创建标签" {
		t.Errorf("label name = %q, want %q", label.Name, "创建标签")
	}
}
