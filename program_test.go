package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestBindProgramEntities(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/program/program_bind_entities" {
			t.Errorf("unexpected path: %s, want /program/program_bind_entities", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s, want POST", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() error: %v", err)
		}
		if r.FormValue("action") != "bind" {
			t.Errorf("action = %q, want bind", r.FormValue("action"))
		}
		if r.FormValue("entity_type") != "story" {
			t.Errorf("entity_type = %q, want story", r.FormValue("entity_type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.BindProgramEntities(context.Background(), &model.BindEntitiesRequest{
		WorkspaceID: "1",
		Action:      "bind",
		EntityType:  "story",
		EntityIDs:   "100001,100002",
	})
	if err != nil {
		t.Fatalf("BindProgramEntities() unexpected error: %v", err)
	}
}

func TestRelateProgramWorkspace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/program/program_relate_workspace" {
			t.Errorf("unexpected path: %s, want /program/program_relate_workspace", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s, want POST", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() error: %v", err)
		}
		if r.FormValue("action") != "unbind" {
			t.Errorf("action = %q, want unbind", r.FormValue("action"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.RelateProgramWorkspace(context.Background(), &model.RelateWorkspaceRequest{
		WorkspaceID:        "1",
		Action:             "unbind",
		RelateWorkspaceIDs: "2,3",
	})
	if err != nil {
		t.Fatalf("RelateProgramWorkspace() unexpected error: %v", err)
	}
}
