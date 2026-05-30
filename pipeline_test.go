package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestAddThirdRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/third_relations" {
			t.Errorf("unexpected path: %s, want /third_relations", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"ThirdRelations":{"id":"601","workspace_id":"1","tapd_id":"100001","tapd_type":"story","source_type":"build","source_id":"build-999"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	rel, err := c.AddThirdRelation(context.Background(), &model.AddThirdRelationRequest{
		SourceType:      "build",
		SourceProjectID: "proj-1",
		SourceID:        "build-999",
		WorkspaceID:     "1",
		TapdID:          "100001",
		TapdType:        "story",
		Operator:        "admin",
	})
	if err != nil {
		t.Fatalf("AddThirdRelation() unexpected error: %v", err)
	}
	if rel.ID != "601" {
		t.Errorf("relation id = %q, want %q", rel.ID, "601")
	}
	if rel.TapdType != "story" {
		t.Errorf("tapd_type = %q, want %q", rel.TapdType, "story")
	}
}

func TestGetThirdRelations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/third_relations" {
			t.Errorf("unexpected path: %s, want /third_relations", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"ThirdRelations":{"id":"601","workspace_id":"1","tapd_id":"100001","tapd_type":"story","source_id":"build-999"}},{"ThirdRelations":{"id":"602","workspace_id":"1","tapd_id":"100001","tapd_type":"story","source_id":"build-1000"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	rels, err := c.GetThirdRelations(context.Background(), &model.GetThirdRelationsRequest{
		SourceType:  "build",
		WorkspaceID: "1",
		TapdID:      "100001",
		TapdType:    "story",
	})
	if err != nil {
		t.Fatalf("GetThirdRelations() unexpected error: %v", err)
	}
	if len(rels) != 2 {
		t.Fatalf("expected 2 relations, got %d", len(rels))
	}
	if rels[0].SourceID != "build-999" {
		t.Errorf("rels[0].source_id = %q, want %q", rels[0].SourceID, "build-999")
	}
	if rels[1].ID != "602" {
		t.Errorf("rels[1].id = %q, want %q", rels[1].ID, "602")
	}
}

func TestDeleteThirdRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/third_relations" {
			t.Errorf("unexpected path: %s, want /third_relations", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("unexpected method: %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"result":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ok, err := c.DeleteThirdRelation(context.Background(), &model.DeleteThirdRelationRequest{
		ID:       "601",
		Operator: "admin",
	})
	if err != nil {
		t.Fatalf("DeleteThirdRelation() unexpected error: %v", err)
	}
	if !ok {
		t.Error("DeleteThirdRelation() = false, want true")
	}
}
