package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

// relatedBugsResponse 用于构造符合 API 实际格式的测试数据
var relatedBugsResponse = `{"status":1,"data":[{"workspace_id":1,"story_id":"100","bug_id":"200"},{"workspace_id":1,"story_id":"100","bug_id":"201"},{"workspace_id":1,"story_id":"100","bug_id":"202"}],"info":"success"}`

func TestGetRelatedBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_related_bugs" {
			t.Errorf("unexpected path: %s, want /stories/get_related_bugs", r.URL.Path)
		}
		if r.URL.Query().Get("story_id") != "100" {
			t.Errorf("story_id = %q, want %q", r.URL.Query().Get("story_id"), "100")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(relatedBugsResponse))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetRelatedBugs(context.Background(), &model.GetRelatedBugsRequest{
		WorkspaceID: "1",
		StoryID:     "100",
	})
	if err != nil {
		t.Fatalf("GetRelatedBugs() unexpected error: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("expected 3 relations, got %d", len(result))
	}
	if result[0].BugID != "200" {
		t.Errorf("result[0].BugID = %q, want %q", result[0].BugID, "200")
	}
	if result[0].StoryID != "100" {
		t.Errorf("result[0].StoryID = %q, want %q", result[0].StoryID, "100")
	}
}

func TestGetRelatedBugs_Empty(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetRelatedBugs(context.Background(), &model.GetRelatedBugsRequest{
		WorkspaceID: "1",
		StoryID:     "100",
	})
	if err != nil {
		t.Fatalf("GetRelatedBugs() unexpected error: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("expected 0 relations, got %d", len(result))
	}
}

func TestCreateRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/relations" {
			t.Errorf("unexpected path: %s, want /relations", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Relation":{"id":"5001","workspace_id":"1","source_type":"story","target_type":"bug","source_id":"100","target_id":"200","created":"2021-05-20 15:10:59","modified":"2021-05-20 15:10:59"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.CreateRelation(context.Background(), &model.CreateRelationRequest{
		WorkspaceID: "1",
		SourceType:  "story",
		TargetType:  "bug",
		SourceID:    "100",
		TargetID:    "200",
	})
	if err != nil {
		t.Fatalf("CreateRelation() unexpected error: %v", err)
	}

	if result.ID != "5001" {
		t.Errorf("id = %q, want %q", result.ID, "5001")
	}
	if result.SourceType != "story" {
		t.Errorf("source_type = %q, want %q", result.SourceType, "story")
	}
	if result.TargetType != "bug" {
		t.Errorf("target_type = %q, want %q", result.TargetType, "bug")
	}
}

func TestCreateRelation_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":0,"data":null,"info":"relation already exists"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.CreateRelation(context.Background(), &model.CreateRelationRequest{
		WorkspaceID: "1",
		SourceType:  "story",
		TargetType:  "bug",
		SourceID:    "100",
		TargetID:    "200",
	})
	if err == nil {
		t.Fatal("expected error for status=0")
	}
}
