package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestCreateStoryCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/story_categories" {
			t.Errorf("unexpected path: %s, want /story_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Category":{"id":"1010104801000086939","workspace_id":"10104801","name":"test111","description":"","parent_id":"0","modified":"2020-08-10 10:02:46","created":"2020-08-10 10:02:45","creator":"dev","modifier":"dev"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	cat, err := c.CreateStoryCategory(context.Background(), &model.CreateStoryCategoryRequest{
		WorkspaceID: "10104801",
		Name:        "test111",
	})
	if err != nil {
		t.Fatalf("CreateStoryCategory() unexpected error: %v", err)
	}
	if cat.ID != "1010104801000086939" {
		t.Errorf("ID = %q, want %q", cat.ID, "1010104801000086939")
	}
	if cat.Name != "test111" {
		t.Errorf("Name = %q, want %q", cat.Name, "test111")
	}
	if cat.Creator != "dev" {
		t.Errorf("Creator = %q, want %q", cat.Creator, "dev")
	}
}

func TestUpdateStoryCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/story_categories" {
			t.Errorf("unexpected path: %s, want /story_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Category":{"id":"1010104801002646384","workspace_id":"10104801","name":"test111","description":"","parent_id":"0","modified":"2025-07-08 15:22:49","created":"2025-06-16 16:20:51","creator":"v_xuanfang","modifier":"v_xuanfang"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	cat, err := c.UpdateStoryCategory(context.Background(), &model.UpdateStoryCategoryRequest{
		WorkspaceID: "10104801",
		ID:          "1010104801002646384",
		Name:        "test111",
	})
	if err != nil {
		t.Fatalf("UpdateStoryCategory() unexpected error: %v", err)
	}
	if cat.ID != "1010104801002646384" {
		t.Errorf("ID = %q, want %q", cat.ID, "1010104801002646384")
	}
	if cat.Name != "test111" {
		t.Errorf("Name = %q, want %q", cat.Name, "test111")
	}
	if cat.Modifier != "v_xuanfang" {
		t.Errorf("Modifier = %q, want %q", cat.Modifier, "v_xuanfang")
	}
}

func TestCountStoryCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/story_categories/count" {
			t.Errorf("unexpected path: %s, want /story_categories/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":5},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountStoryCategories(context.Background(), &model.CountStoryCategoriesRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("CountStoryCategories() unexpected error: %v", err)
	}
	if count != 5 {
		t.Errorf("count = %d, want 5", count)
	}
}

func TestAddStoryLinkRelations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/add_story_link_relations" {
			t.Errorf("unexpected path: %s, want /stories/add_story_link_relations", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":1},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ok, err := c.AddStoryLinkRelations(context.Background(), &model.AddStoryLinkRelationsRequest{
		WorkspaceID:   "10104801",
		SrcStoryID:    "100",
		TargetStoryID: "200",
	})
	if err != nil {
		t.Fatalf("AddStoryLinkRelations() unexpected error: %v", err)
	}
	if !ok {
		t.Error("expected true, got false")
	}
}

func TestGetLinkStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_link_stories" {
			t.Errorf("unexpected path: %s, want /stories/get_link_stories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"type":"derivation","id":"1010158231500691433","story_id":"1010158231500691431","workspace_id":"10158231","actas":"target","created":"2019-08-01 16:32:22","modified":"2022-03-18 18:54:25","linked_workspace_id":10158231}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	rels, err := c.GetLinkStories(context.Background(), &model.GetLinkStoriesRequest{
		WorkspaceID: "10158231",
		StoryID:     "1010158231500691431",
	})
	if err != nil {
		t.Fatalf("GetLinkStories() unexpected error: %v", err)
	}
	if len(rels) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(rels))
	}
	if rels[0].Type != "derivation" {
		t.Errorf("Type = %q, want %q", rels[0].Type, "derivation")
	}
	if rels[0].ID != "1010158231500691433" {
		t.Errorf("ID = %q, want %q", rels[0].ID, "1010158231500691433")
	}
	if rels[0].Actas != "target" {
		t.Errorf("Actas = %q, want %q", rels[0].Actas, "target")
	}
	if rels[0].LinkedWorkspaceID != 10158231 {
		t.Errorf("LinkedWorkspaceID = %d, want %d", rels[0].LinkedWorkspaceID, 10158231)
	}
}

func TestAddStoryTcase(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/add_story_tcase" {
			t.Errorf("unexpected path: %s, want /stories/add_story_tcase", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success_id":["1010104801077291609"]},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ids, err := c.AddStoryTcase(context.Background(), &model.AddStoryTcaseRequest{
		WorkspaceID: "10104801",
		StoryID:     "100",
		TcaseID:     "1010104801077291609",
	})
	if err != nil {
		t.Fatalf("AddStoryTcase() unexpected error: %v", err)
	}
	if len(ids) != 1 {
		t.Fatalf("expected 1 id, got %d", len(ids))
	}
	if ids[0] != "1010104801077291609" {
		t.Errorf("ids[0] = %q, want %q", ids[0], "1010104801077291609")
	}
}

func TestGetStoryTcase(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_story_tcase" {
			t.Errorf("unexpected path: %s, want /stories/get_story_tcase", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"TestPlanStoryTcaseRelation":{"id":"1010104801021215005","workspace_id":"10104801","test_plan_id":"0","story_id":"1010104801866191641","tcase_id":"1010104801076110789","sort":"0","creator":"v_xuanfang","created":"2021-08-06 12:35:01"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	rels, err := c.GetStoryTcase(context.Background(), &model.GetStoryTcaseRequest{
		WorkspaceID: "10104801",
		StoryID:     "1010104801866191641",
	})
	if err != nil {
		t.Fatalf("GetStoryTcase() unexpected error: %v", err)
	}
	if len(rels) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(rels))
	}
	if rels[0].ID != "1010104801021215005" {
		t.Errorf("ID = %q, want %q", rels[0].ID, "1010104801021215005")
	}
	if rels[0].TcaseID != "1010104801076110789" {
		t.Errorf("TcaseID = %q, want %q", rels[0].TcaseID, "1010104801076110789")
	}
	if rels[0].Creator != "v_xuanfang" {
		t.Errorf("Creator = %q, want %q", rels[0].Creator, "v_xuanfang")
	}
}

func TestRemoveStoryBugRelations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/remove_story_bug_raletions" {
			t.Errorf("unexpected path: %s, want /stories/remove_story_bug_raletions", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ok, err := c.RemoveStoryBugRelations(context.Background(), &model.RemoveStoryBugRelationsRequest{
		WorkspaceID: "10104801",
		StoryID:     "100",
		BugID:       "200",
	})
	if err != nil {
		t.Fatalf("RemoveStoryBugRelations() unexpected error: %v", err)
	}
	if !ok {
		t.Error("expected true, got false")
	}
}

func TestGetTimeRelativeStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_time_relative_stories" {
			t.Errorf("unexpected path: %s, want /stories/get_time_relative_stories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemTimeRelation":{"id":"1210104801000007813","workspace_id":"10104801","workitem_type":"story","workitem_id":"1010104801854915911","src_field":"begin","dst_workspace_id":"10104801","dst_workitem_type":"story","dst_workitem_id":"1010104801854917775","dst_field":"due","relation_type":"after"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	rels, err := c.GetTimeRelativeStories(context.Background(), &model.GetTimeRelativeStoriesRequest{
		WorkspaceID: "10104801",
		StoryID:     "1010104801854915911",
	})
	if err != nil {
		t.Fatalf("GetTimeRelativeStories() unexpected error: %v", err)
	}
	if len(rels) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(rels))
	}
	if rels[0].ID != "1210104801000007813" {
		t.Errorf("ID = %q, want %q", rels[0].ID, "1210104801000007813")
	}
	if rels[0].RelationType != "after" {
		t.Errorf("RelationType = %q, want %q", rels[0].RelationType, "after")
	}
	if rels[0].SrcField != "begin" {
		t.Errorf("SrcField = %q, want %q", rels[0].SrcField, "begin")
	}
	if rels[0].DstField != "due" {
		t.Errorf("DstField = %q, want %q", rels[0].DstField, "due")
	}
}

func TestSaveTimeRelations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/save_time_relations" {
			t.Errorf("unexpected path: %s, want /stories/save_time_relations", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"result":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ok, err := c.SaveTimeRelations(context.Background(), &model.SaveTimeRelationsRequest{
		WorkspaceID: "10104801",
		Relations: []model.SaveTimeRelationItem{
			{WorkitemID: "100", DstWorkitemID: "200", SrcField: "begin", DstField: "due"},
		},
		CurrentUser: "testuser",
	})
	if err != nil {
		t.Fatalf("SaveTimeRelations() unexpected error: %v", err)
	}
	if !ok {
		t.Error("expected true, got false")
	}
}

func TestDeleteTimeRelations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/delete_time_relations" {
			t.Errorf("unexpected path: %s, want /stories/delete_time_relations", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"num":1},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	num, err := c.DeleteTimeRelations(context.Background(), &model.DeleteTimeRelationsRequest{
		WorkspaceID: "10104801",
		RelationIDs: []string{"12345"},
		CurrentUser: "testuser",
	})
	if err != nil {
		t.Fatalf("DeleteTimeRelations() unexpected error: %v", err)
	}
	if num != 1 {
		t.Errorf("num = %d, want 1", num)
	}
}

func TestGetStoryTemplateList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/template_list" {
			t.Errorf("unexpected path: %s, want /stories/template_list", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemTemplate":{"id":"1010104801000002561","name":"系统默认模板","description":"系统自动创建xxx","sort":"1","default":"1","creator":"SYSTEM","editor_type":"1"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	templates, err := c.GetStoryTemplateList(context.Background(), &model.GetStoryTemplateListRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("GetStoryTemplateList() unexpected error: %v", err)
	}
	if len(templates) != 1 {
		t.Fatalf("expected 1 template, got %d", len(templates))
	}
	if templates[0].ID != "1010104801000002561" {
		t.Errorf("ID = %q, want %q", templates[0].ID, "1010104801000002561")
	}
	if templates[0].Name != "系统默认模板" {
		t.Errorf("Name = %q, want %q", templates[0].Name, "系统默认模板")
	}
	if templates[0].Default != "1" {
		t.Errorf("Default = %q, want %q", templates[0].Default, "1")
	}
}

func TestGetDefaultStoryTemplate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_default_story_template" {
			t.Errorf("unexpected path: %s, want /stories/get_default_story_template", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemTemplateField":{"id":"1010104801015287651","workspace_id":"10104801","type":"story","template_id":"1010104801000850579","field":"description","value":"","required":"1","sort":"0","linkage_rules":""}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	fields, err := c.GetDefaultStoryTemplate(context.Background(), &model.GetDefaultStoryTemplateRequest{
		WorkspaceID: "10104801",
		TemplateID:  "1010104801000850579",
	})
	if err != nil {
		t.Fatalf("GetDefaultStoryTemplate() unexpected error: %v", err)
	}
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	if fields[0].ID != "1010104801015287651" {
		t.Errorf("ID = %q, want %q", fields[0].ID, "1010104801015287651")
	}
	if fields[0].Field != "description" {
		t.Errorf("Field = %q, want %q", fields[0].Field, "description")
	}
	if fields[0].Required != "1" {
		t.Errorf("Required = %q, want %q", fields[0].Required, "1")
	}
}

func TestGetStorySteps(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_story_step_list" {
			t.Errorf("unexpected path: %s, want /stories/get_story_step_list", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemStepInfo":{"id":"1070002667000137213","workspace_id":"70002667","entity_type":"story","workitem_id":"1070002667006658827","step":"step_2970811_1","status":"0","owner":"","begin":null,"due":null,"effort":"3","iteration_id":"0","begin_time":"2026-01-04 09:37:57","complete_time":"2026-01-04 09:38:23","time_cost":"26","completer":"ocenhu"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	steps, err := c.GetStorySteps(context.Background(), &model.GetStoryStepsRequest{
		WorkspaceID: "70002667",
		StoryID:     "1070002667006658827",
	})
	if err != nil {
		t.Fatalf("GetStorySteps() unexpected error: %v", err)
	}
	if len(steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(steps))
	}
	if steps[0].ID != "1070002667000137213" {
		t.Errorf("ID = %q, want %q", steps[0].ID, "1070002667000137213")
	}
	if steps[0].Step != "step_2970811_1" {
		t.Errorf("Step = %q, want %q", steps[0].Step, "step_2970811_1")
	}
	if steps[0].Completer != "ocenhu" {
		t.Errorf("Completer = %q, want %q", steps[0].Completer, "ocenhu")
	}
	if steps[0].TimeCost != "26" {
		t.Errorf("TimeCost = %q, want %q", steps[0].TimeCost, "26")
	}
}

func TestGetRemovedStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_removed_stories" {
			t.Errorf("unexpected path: %s, want /stories/get_removed_stories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"RemovedStory":{"id":"1010104801854921589","name":"cat","creator":"tapd","created":"2021-08-25 15:37:16","operation_user":"v_xuanfang","deleted":"2021-09-13 16:48:16","is_archived":"0"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	stories, err := c.GetRemovedStories(context.Background(), &model.GetRemovedStoriesRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("GetRemovedStories() unexpected error: %v", err)
	}
	if len(stories) != 1 {
		t.Fatalf("expected 1 story, got %d", len(stories))
	}
	if stories[0].ID != "1010104801854921589" {
		t.Errorf("ID = %q, want %q", stories[0].ID, "1010104801854921589")
	}
	if stories[0].Name != "cat" {
		t.Errorf("Name = %q, want %q", stories[0].Name, "cat")
	}
	if stories[0].OperationUser != "v_xuanfang" {
		t.Errorf("OperationUser = %q, want %q", stories[0].OperationUser, "v_xuanfang")
	}
	if stories[0].IsArchived != "0" {
		t.Errorf("IsArchived = %q, want %q", stories[0].IsArchived, "0")
	}
}

func TestCopyStory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/copy_story" {
			t.Errorf("unexpected path: %s, want /stories/copy_story", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Story":{"id":"1000000755854845111","name":"bbbbbbbb","workspace_id":"755","creator":"anyechen","created":"2020-12-09 17:00:09","status":"planning"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	story, err := c.CopyStory(context.Background(), &model.CopyStoryRequest{
		WorkspaceID:    "755",
		SrcStoryID:     "100",
		DstWorkspaceID: "755",
	})
	if err != nil {
		t.Fatalf("CopyStory() unexpected error: %v", err)
	}
	if story.ID != "1000000755854845111" {
		t.Errorf("ID = %q, want %q", story.ID, "1000000755854845111")
	}
	if story.Name != "bbbbbbbb" {
		t.Errorf("Name = %q, want %q", story.Name, "bbbbbbbb")
	}
	if story.Status != "planning" {
		t.Errorf("Status = %q, want %q", story.Status, "planning")
	}
}

func TestUpdateStoryParent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/update_story_parent" {
			t.Errorf("unexpected path: %s, want /stories/update_story_parent", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Story":{"id":"1010104801871430407","name":"aatt","workspace_id":"10104801","creator":"v_xuanfang","status":"status_5","parent_id":"1010104801870636009"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	story, err := c.UpdateStoryParent(context.Background(), &model.UpdateStoryParentRequest{
		WorkspaceID: "10104801",
		StoryID:     "1010104801871430407",
		ParentID:    "1010104801870636009",
	})
	if err != nil {
		t.Fatalf("UpdateStoryParent() unexpected error: %v", err)
	}
	if story.ID != "1010104801871430407" {
		t.Errorf("ID = %q, want %q", story.ID, "1010104801871430407")
	}
	if story.ParentID != "1010104801870636009" {
		t.Errorf("ParentID = %q, want %q", story.ParentID, "1010104801870636009")
	}
	if story.Status != "status_5" {
		t.Errorf("Status = %q, want %q", story.Status, "status_5")
	}
}

func TestChangeWorkitemType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/change_workitem_type" {
			t.Errorf("unexpected path: %s, want /stories/change_workitem_type", r.URL.Path)
		}
		// data 中直接是 Story 字段，无 "Story" 信封
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"id":"1000000755854804275","name":"testsetestsetse","workspace_id":"755","workitem_type_id":"1000000755000033239","status":"resolved"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	story, err := c.ChangeWorkitemType(context.Background(), &model.ChangeWorkitemTypeRequest{
		WorkspaceID:    "755",
		StoryID:        "1000000755854804275",
		WorkitemTypeID: "1000000755000033239",
	})
	if err != nil {
		t.Fatalf("ChangeWorkitemType() unexpected error: %v", err)
	}
	if story.ID != "1000000755854804275" {
		t.Errorf("ID = %q, want %q", story.ID, "1000000755854804275")
	}
	if story.Name != "testsetestsetse" {
		t.Errorf("Name = %q, want %q", story.Name, "testsetestsetse")
	}
	if story.WorkitemTypeID != "1000000755000033239" {
		t.Errorf("WorkitemTypeID = %q, want %q", story.WorkitemTypeID, "1000000755000033239")
	}
	if story.Status != "resolved" {
		t.Errorf("Status = %q, want %q", story.Status, "resolved")
	}
}

func TestGetStoriesByViewConfID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_stories_by_view_conf_id" {
			t.Errorf("unexpected path: %s, want /stories/get_stories_by_view_conf_id", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Story":{"id":"1010104801855983211","name":"sss1","priority":"","status":"developing","owner":"anyechen;","iteration_id":"1010104801000708781"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	stories, err := c.GetStoriesByViewConfID(context.Background(), &model.GetStoriesByViewConfIDRequest{
		WorkspaceID: "10104801",
		ViewConfID:  "view123",
	})
	if err != nil {
		t.Fatalf("GetStoriesByViewConfID() unexpected error: %v", err)
	}
	if len(stories) != 1 {
		t.Fatalf("expected 1 story, got %d", len(stories))
	}
	if stories[0].ID != "1010104801855983211" {
		t.Errorf("ID = %q, want %q", stories[0].ID, "1010104801855983211")
	}
	if stories[0].Name != "sss1" {
		t.Errorf("Name = %q, want %q", stories[0].Name, "sss1")
	}
	if stories[0].Status != "developing" {
		t.Errorf("Status = %q, want %q", stories[0].Status, "developing")
	}
}

func TestBatchUpdateStory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/batch_update_story" {
			t.Errorf("unexpected path: %s, want /stories/batch_update_story", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"msg":"batch update success"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	msg, err := c.BatchUpdateStory(context.Background(), &model.BatchUpdateStoryRequest{
		WorkspaceID: "10104801",
		Workitems: []model.BatchUpdateStoryItem{
			{ID: "100", Name: "updated"},
		},
	})
	if err != nil {
		t.Fatalf("BatchUpdateStory() unexpected error: %v", err)
	}
	if msg != "batch update success" {
		t.Errorf("msg = %q, want %q", msg, "batch update success")
	}
}

func TestGetSecretInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/get_secret_info" {
			t.Errorf("unexpected path: %s, want /stories/get_secret_info", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"creator":"testuser","allow_list":"user1;user2","secret_root_id":"100","add_participant_fields":"true","secret_scrope":"secret"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	info, err := c.GetSecretInfo(context.Background(), &model.GetSecretInfoRequest{
		WorkspaceID: "10104801",
		StoryID:     "100",
	})
	if err != nil {
		t.Fatalf("GetSecretInfo() unexpected error: %v", err)
	}
	if info.Creator != "testuser" {
		t.Errorf("Creator = %q, want %q", info.Creator, "testuser")
	}
	if info.AllowList != "user1;user2" {
		t.Errorf("AllowList = %q, want %q", info.AllowList, "user1;user2")
	}
	if info.SecretScope != "secret" {
		t.Errorf("SecretScope = %q, want %q", info.SecretScope, "secret")
	}
}

func TestGetSecretStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/secret_stories" {
			t.Errorf("unexpected path: %s, want /secret_stories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"list":["100","200","300"]},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetSecretStories(context.Background(), &model.GetSecretStoriesRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("GetSecretStories() unexpected error: %v", err)
	}
	if len(result.List) != 3 {
		t.Fatalf("expected 3 items, got %d", len(result.List))
	}
	if result.List[0] != "100" {
		t.Errorf("List[0] = %q, want %q", result.List[0], "100")
	}
}

func TestCountSecretStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/secret_stories/count" {
			t.Errorf("unexpected path: %s, want /secret_stories/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":10},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountSecretStories(context.Background(), "10104801")
	if err != nil {
		t.Fatalf("CountSecretStories() unexpected error: %v", err)
	}
	if count != 10 {
		t.Errorf("count = %d, want 10", count)
	}
}

func TestBatchUpdateSecretInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/batch_update_secret_info" {
			t.Errorf("unexpected path: %s, want /stories/batch_update_secret_info", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"code":"0","msg":"success"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	msg, err := c.BatchUpdateSecretInfo(context.Background(), &model.BatchUpdateSecretInfoRequest{
		WorkspaceID:          "10104801",
		StoryIDList:          "100|200",
		SecretScope:          "secret",
		AllowList:            "user1;user2",
		AddParticipantFields: "true",
		CurrentUser:          "testuser",
	})
	if err != nil {
		t.Fatalf("BatchUpdateSecretInfo() unexpected error: %v", err)
	}
	if msg != "success" {
		t.Errorf("msg = %q, want %q", msg, "success")
	}
}

func TestStoryFilterToQueryToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/filter_to_query_token" {
			t.Errorf("unexpected path: %s, want /stories/filter_to_query_token", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"queryToken":"abc123","href":"https://www.tapd.cn/10104801/sparrow/stories/query_token/abc123"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.StoryFilterToQueryToken(context.Background(), &model.FilterToQueryTokenRequest{
		WorkspaceID: "10104801",
		Filters:     map[string]string{"status": "open"},
	})
	if err != nil {
		t.Fatalf("StoryFilterToQueryToken() unexpected error: %v", err)
	}
	if result.QueryToken != "abc123" {
		t.Errorf("QueryToken = %q, want %q", result.QueryToken, "abc123")
	}
	if result.Href != "https://www.tapd.cn/10104801/sparrow/stories/query_token/abc123" {
		t.Errorf("Href = %q, want %q", result.Href, "https://www.tapd.cn/10104801/sparrow/stories/query_token/abc123")
	}
}

func TestStoryIDsToQueryToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/ids_to_query_token" {
			t.Errorf("unexpected path: %s, want /stories/ids_to_query_token", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"queryToken":"def456","href":"https://www.tapd.cn/10104801/sparrow/stories/query_token/def456"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.StoryIDsToQueryToken(context.Background(), &model.IDsToQueryTokenRequest{
		WorkspaceID: "10104801",
		IDs:         "100,200,300",
	})
	if err != nil {
		t.Fatalf("StoryIDsToQueryToken() unexpected error: %v", err)
	}
	if result.QueryToken != "def456" {
		t.Errorf("QueryToken = %q, want %q", result.QueryToken, "def456")
	}
	if result.Href != "https://www.tapd.cn/10104801/sparrow/stories/query_token/def456" {
		t.Errorf("Href = %q, want %q", result.Href, "https://www.tapd.cn/10104801/sparrow/stories/query_token/def456")
	}
}

func TestRemoveStoryLinkRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/remove_story_link_relation" {
			t.Errorf("unexpected path: %s, want /stories/remove_story_link_relation", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":1},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ok, err := c.RemoveStoryLinkRelation(context.Background(), &model.RemoveStoryLinkRelationRequest{
		WorkspaceID:   "10104801",
		SrcStoryID:    "100",
		TargetStoryID: "200",
	})
	if err != nil {
		t.Fatalf("RemoveStoryLinkRelation() unexpected error: %v", err)
	}
	if !ok {
		t.Error("expected true, got false")
	}
}

func TestResetWorkitemSteps(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/reset_workitem_steps" {
			t.Errorf("unexpected path: %s, want /stories/reset_workitem_steps", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Story":{"id":"100","name":"ResetStory","workspace_id":"10104801","status":"open"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	story, err := c.ResetWorkitemSteps(context.Background(), &model.ResetWorkitemStepsRequest{
		WorkspaceID: "10104801",
		StoryID:     "100",
		ResetType:   "status",
		ResetDst:    "open",
	})
	if err != nil {
		t.Fatalf("ResetWorkitemSteps() unexpected error: %v", err)
	}
	if story.ID != "100" {
		t.Errorf("ID = %q, want %q", story.ID, "100")
	}
	if story.Name != "ResetStory" {
		t.Errorf("Name = %q, want %q", story.Name, "ResetStory")
	}
	if story.Status != "open" {
		t.Errorf("Status = %q, want %q", story.Status, "open")
	}
}

func TestUpdateStoryStepStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/update_story_step_status" {
			t.Errorf("unexpected path: %s, want /stories/update_story_step_status", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Story":{"id":"100","name":"StepStory","workspace_id":"10104801","status":"developing"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	story, err := c.UpdateStoryStepStatus(context.Background(), &model.UpdateStoryStepStatusRequest{
		WorkspaceID: "10104801",
		StoryID:     "100",
		Step:        "step_1",
	})
	if err != nil {
		t.Fatalf("UpdateStoryStepStatus() unexpected error: %v", err)
	}
	if story.ID != "100" {
		t.Errorf("ID = %q, want %q", story.ID, "100")
	}
	if story.Name != "StepStory" {
		t.Errorf("Name = %q, want %q", story.Name, "StepStory")
	}
	if story.Status != "developing" {
		t.Errorf("Status = %q, want %q", story.Status, "developing")
	}
}

func TestCountStoriesByCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stories/count_by_categories" {
			t.Errorf("unexpected path: %s, want /stories/count_by_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"cat1":5,"cat2":10},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.CountStoriesByCategories(context.Background(), &model.CountStoriesByCategoriesRequest{
		WorkspaceID: "10104801",
		CategoryID:  "cat1,cat2",
	})
	if err != nil {
		t.Fatalf("CountStoriesByCategories() unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(result))
	}
	if result["cat1"] != 5 {
		t.Errorf("cat1 = %d, want 5", result["cat1"])
	}
	if result["cat2"] != 10 {
		t.Errorf("cat2 = %d, want 10", result["cat2"])
	}
}
