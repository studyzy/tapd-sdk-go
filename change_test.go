package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestGetStoryChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/story_changes" {
			t.Errorf("unexpected path: %s, want /story_changes", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemChange":{"id":"1010104801027730979","workspace_id":"10104801","creator":"anyechen","created":"2015-06-30 14:28:53","change_summary":"planning","comment":"","changes":"[{\"field\":\"parent_id\",\"value_before\":\"0\",\"value_after\":\"1010104801056751739\"}]","entity_type":"Story","change_type":"","story_id":"1010104801056751735"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetStoryChangesRequest{
		WorkspaceID: "10104801",
	}
	changes, err := c.GetStoryChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("GetStoryChanges() unexpected error: %v", err)
	}
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].ID != "1010104801027730979" {
		t.Errorf("id = %q, want %q", changes[0].ID, "1010104801027730979")
	}
	if changes[0].WorkspaceID != "10104801" {
		t.Errorf("workspace_id = %q, want %q", changes[0].WorkspaceID, "10104801")
	}
	if changes[0].Creator != "anyechen" {
		t.Errorf("creator = %q, want %q", changes[0].Creator, "anyechen")
	}
	if changes[0].ChangeSummary != "planning" {
		t.Errorf("change_summary = %q, want %q", changes[0].ChangeSummary, "planning")
	}
	if changes[0].EntityType != "Story" {
		t.Errorf("entity_type = %q, want %q", changes[0].EntityType, "Story")
	}
	if changes[0].StoryID != "1010104801056751735" {
		t.Errorf("story_id = %q, want %q", changes[0].StoryID, "1010104801056751735")
	}
}

func TestCountStoryChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/story_changes/count" {
			t.Errorf("unexpected path: %s, want /story_changes/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":42},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountStoryChangesRequest{
		WorkspaceID: "10104801",
	}
	count, err := c.CountStoryChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("CountStoryChanges() unexpected error: %v", err)
	}
	if count != 42 {
		t.Errorf("count = %d, want 42", count)
	}
}

func TestGetBugChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bug_changes" {
			t.Errorf("unexpected path: %s, want /bug_changes", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"BugChange":{"id":"10101582315000015921","bug_id":"1010158231500628815","author":"anyechen","field":"severity","old_value":"serious","new_value":"normal","memo":"","created":"2019-06-26 20:48:52","workspace_id":"10158231"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetBugChangesRequest{
		WorkspaceID: "10158231",
	}
	changes, err := c.GetBugChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("GetBugChanges() unexpected error: %v", err)
	}
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].ID != "10101582315000015921" {
		t.Errorf("id = %q, want %q", changes[0].ID, "10101582315000015921")
	}
	if changes[0].BugID != "1010158231500628815" {
		t.Errorf("bug_id = %q, want %q", changes[0].BugID, "1010158231500628815")
	}
	if changes[0].Author != "anyechen" {
		t.Errorf("author = %q, want %q", changes[0].Author, "anyechen")
	}
	if changes[0].Field != "severity" {
		t.Errorf("field = %q, want %q", changes[0].Field, "severity")
	}
	if changes[0].OldValue != "serious" {
		t.Errorf("old_value = %q, want %q", changes[0].OldValue, "serious")
	}
	if changes[0].NewValue != "normal" {
		t.Errorf("new_value = %q, want %q", changes[0].NewValue, "normal")
	}
}

func TestCountBugChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bug_changes/count" {
			t.Errorf("unexpected path: %s, want /bug_changes/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":7},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountBugChangesRequest{
		WorkspaceID: "10158231",
	}
	count, err := c.CountBugChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("CountBugChanges() unexpected error: %v", err)
	}
	if count != 7 {
		t.Errorf("count = %d, want 7", count)
	}
}

func TestGetTaskChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/task_changes" {
			t.Errorf("unexpected path: %s, want /task_changes", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemChange":{"id":"1010158231046789917","workspace_id":"10158231","creator":"anyechen","created":"2020-02-28 00:55:03","change_summary":"open","comment":"","changes":"[{\"field\":\"effort\",\"value_before\":\"0\",\"value_after\":\"10\"}]","entity_type":"Task","task_id":"1010158231500600411"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetTaskChangesRequest{
		WorkspaceID: "10158231",
	}
	changes, err := c.GetTaskChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTaskChanges() unexpected error: %v", err)
	}
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].ID != "1010158231046789917" {
		t.Errorf("id = %q, want %q", changes[0].ID, "1010158231046789917")
	}
	if changes[0].WorkspaceID != "10158231" {
		t.Errorf("workspace_id = %q, want %q", changes[0].WorkspaceID, "10158231")
	}
	if changes[0].Creator != "anyechen" {
		t.Errorf("creator = %q, want %q", changes[0].Creator, "anyechen")
	}
	if changes[0].ChangeSummary != "open" {
		t.Errorf("change_summary = %q, want %q", changes[0].ChangeSummary, "open")
	}
	if changes[0].EntityType != "Task" {
		t.Errorf("entity_type = %q, want %q", changes[0].EntityType, "Task")
	}
	if changes[0].TaskID != "1010158231500600411" {
		t.Errorf("task_id = %q, want %q", changes[0].TaskID, "1010158231500600411")
	}
}

func TestCountTaskChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/task_changes/count" {
			t.Errorf("unexpected path: %s, want /task_changes/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":15},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountTaskChangesRequest{
		WorkspaceID: "10158231",
	}
	count, err := c.CountTaskChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("CountTaskChanges() unexpected error: %v", err)
	}
	if count != 15 {
		t.Errorf("count = %d, want 15", count)
	}
}

func TestGetIterationChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iteration_changes" {
			t.Errorf("unexpected path: %s, want /iteration_changes", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"IterationChange":{"id":"1020355782015033213","iteration_id":"1020355782000700291","author":"v_xinyucao","field":"name","old_value":"","new_value":"对方的身份","memo":"","created":"2020-04-29 10:42:02","operater_type":"add","workspace_id":"20355782"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetIterationChangesRequest{
		WorkspaceID: "20355782",
		IterationID: "1020355782000700291",
	}
	changes, err := c.GetIterationChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("GetIterationChanges() unexpected error: %v", err)
	}
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].ID != "1020355782015033213" {
		t.Errorf("id = %q, want %q", changes[0].ID, "1020355782015033213")
	}
	if changes[0].IterationID != "1020355782000700291" {
		t.Errorf("iteration_id = %q, want %q", changes[0].IterationID, "1020355782000700291")
	}
	if changes[0].Author != "v_xinyucao" {
		t.Errorf("author = %q, want %q", changes[0].Author, "v_xinyucao")
	}
	if changes[0].Field != "name" {
		t.Errorf("field = %q, want %q", changes[0].Field, "name")
	}
	if changes[0].NewValue != "对方的身份" {
		t.Errorf("new_value = %q, want %q", changes[0].NewValue, "对方的身份")
	}
	if changes[0].OperaterType != "add" {
		t.Errorf("operater_type = %q, want %q", changes[0].OperaterType, "add")
	}
	if changes[0].WorkspaceID != "20355782" {
		t.Errorf("workspace_id = %q, want %q", changes[0].WorkspaceID, "20355782")
	}
}
