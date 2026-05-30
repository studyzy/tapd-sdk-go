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
		w.Write([]byte(`{"status":1,"data":{"msg":"success"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	data, err := c.AddCodeCommitInfo(context.Background(), &model.AddCodeCommitInfoRequest{
		WorkspaceID: "11111111",
		Message:     "fix bug #100",
		Author:      "admin",
		CommitID:    "abc123",
		Files:       []string{"U main.go", "A new.go"},
		Repo:        "my-repo",
		RepoID:      "repo-001",
		CommitTime:  "2026-01-01 10:00:00",
		HookURL:     "https://example.com/hook",
		Ref:         "refs/heads/main",
	})
	if err != nil {
		t.Fatalf("AddCodeCommitInfo() unexpected error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil data")
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
		w.Write([]byte(`{"status":1,"data":[{"CodeCommitInfo":{"id":"1001","workspace_id":"11111111","message":"fix bug #100","author":"admin","created":"2026-01-01 10:00:00"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	data, err := c.GetCodeCommitInfos(context.Background(), &model.GetCodeCommitInfosRequest{
		WorkspaceID: "11111111",
		Limit:       10,
		Page:        1,
	})
	if err != nil {
		t.Fatalf("GetCodeCommitInfos() unexpected error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil data")
	}
	if string(data) == "" {
		t.Error("expected non-empty data")
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
