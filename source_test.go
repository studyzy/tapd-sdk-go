package tapd

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestAddCodeCommitInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/code_commit_infos" {
			t.Errorf("unexpected path: %s, want /code_commit_infos", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) == "" {
			t.Fatal("expected non-empty body")
		}
		// 验证 body 是合法的 JSON
		var m map[string]interface{}
		if err := json.Unmarshal(body, &m); err != nil {
			t.Fatalf("body is not valid JSON: %v", err)
		}
		if m["workspace_id"] != "11111111" {
			t.Errorf("workspace_id = %v, want 11111111", m["workspace_id"])
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"id":"1001","hook_user_name":"admin","commit_id":"abc123","workspace_id":"11111111","message":"fix bug #100"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.AddCodeCommitInfo(context.Background(), &model.AddCodeCommitInfoRequest{
		WorkspaceID: "11111111",
		Message:     "fix bug #100",
		Author:      "admin",
		CommitID:    "abc123",
		Files:       []string{"U main.go", "A new.go"},
		Repo:        "my-repo",
		RepoID:      "repo-001",
		CommitTime:  "2026-01-01 10:00:00",
		CommitURL:   "https://example.com/commit/abc123",
	})
	if err != nil {
		t.Fatalf("AddCodeCommitInfo() unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.ID != "1001" {
		t.Errorf("id = %q, want %q", result.ID, "1001")
	}
}

func TestGetCodeCommitInfos(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/code_commit_infos" {
			t.Errorf("unexpected path: %s, want /code_commit_infos", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "11111111" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "11111111")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"id":"1001","workspace_id":"11111111","message":"fix bug #100","user_name":"admin","commit_id":"abc123"}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetCodeCommitInfos(context.Background(), &model.GetCodeCommitInfosRequest{
		WorkspaceID: "11111111",
		Limit:       10,
		Page:        1,
	})
	if err != nil {
		t.Fatalf("GetCodeCommitInfos() unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 commit info, got %d", len(result))
	}
	if result[0].ID != "1001" {
		t.Errorf("id = %q, want %q", result[0].ID, "1001")
	}
}

func TestGetCodeCommitInfos_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":0,"data":null,"info":"permission denied"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.GetCodeCommitInfos(context.Background(), &model.GetCodeCommitInfosRequest{
		WorkspaceID: "11111111",
	})
	if err == nil {
		t.Fatal("expected error for status=0")
	}
}

func TestGetCodeCommitObjects(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/code_commit_objects/workitems" {
			t.Errorf("unexpected path: %s, want /code_commit_objects/workitems", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s, want GET", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "20355782" {
			t.Errorf("workspace_id = %q, want 20355782", r.URL.Query().Get("workspace_id"))
		}
		if r.URL.Query().Get("commit_id") != "7b0645c6a467a502fe1d3b678fea8bdf2890aa8d" {
			t.Errorf("commit_id = %q", r.URL.Query().Get("commit_id"))
		}
		if r.URL.Query().Get("entity_type") != "story" {
			t.Errorf("entity_type = %q, want story", r.URL.Query().Get("entity_type"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Task":{"id":"1020355782500602947","name":"666","workspace_id":"20355782","status":"open"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	data, err := c.GetCodeCommitObjects(context.Background(), &model.GetCodeCommitObjectsRequest{
		WorkspaceID: "20355782",
		CommitID:    "7b0645c6a467a502fe1d3b678fea8bdf2890aa8d",
		EntityType:  "story",
	})
	if err != nil {
		t.Fatalf("GetCodeCommitObjects() unexpected error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil data")
	}
	// 验证返回的 JSON 包含预期内容
	if !json.Valid(data) {
		t.Fatalf("returned data is not valid JSON: %s", string(data))
	}
	var items []map[string]interface{}
	if err := json.Unmarshal(data, &items); err != nil {
		t.Fatalf("failed to unmarshal data: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
}
