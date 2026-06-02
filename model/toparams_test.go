package model

import (
	"encoding/json"
	"testing"
)

// TestSetOptional 测试 setOptional 辅助函数
func TestSetOptional(t *testing.T) {
	params := map[string]string{}
	setOptional(params, "key1", "val1")
	setOptional(params, "key2", "")
	if params["key1"] != "val1" {
		t.Errorf("expected key1=val1, got %q", params["key1"])
	}
	if _, ok := params["key2"]; ok {
		t.Error("empty value should not be set")
	}
}

// === story.go ===

func TestListStoriesRequest_ToParams(t *testing.T) {
	req := &ListStoriesRequest{
		WorkspaceID:   "100",
		Name:          "test",
		PriorityLabel: "high",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "test" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["priority_label"] != "high" {
		t.Errorf("priority_label: got %q", params["priority_label"])
	}
	if _, ok := params["status"]; ok {
		t.Error("empty status should not be in params")
	}
	if _, ok := params["id"]; ok {
		t.Error("empty id should not be in params")
	}
}

func TestCreateStoryRequest_ToParams(t *testing.T) {
	req := &CreateStoryRequest{
		WorkspaceID:   "1",
		Name:          "Test Story",
		PriorityLabel: "high",
		CustomFields:  map[string]string{"custom_field_one": "cf1"},
	}
	params := req.ToParams()
	if params["workspace_id"] != "1" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "Test Story" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["priority_label"] != "high" {
		t.Errorf("priority_label: got %q", params["priority_label"])
	}
	if params["custom_field_one"] != "cf1" {
		t.Errorf("custom_field_one: got %q", params["custom_field_one"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestUpdateStoryRequest_ToParams(t *testing.T) {
	req := &UpdateStoryRequest{
		WorkspaceID:  "1",
		ID:           "99",
		Status:       "done",
		CustomFields: map[string]string{"custom_field_two": "cf2"},
	}
	params := req.ToParams()
	if params["workspace_id"] != "1" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "99" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["status"] != "done" {
		t.Errorf("status: got %q", params["status"])
	}
	if params["custom_field_two"] != "cf2" {
		t.Errorf("custom_field_two: got %q", params["custom_field_two"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

func TestCountStoriesRequest_ToParams(t *testing.T) {
	req := &CountStoriesRequest{
		WorkspaceID: "1",
		Status:      "open",
	}
	params := req.ToParams()
	if params["workspace_id"] != "1" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "open" {
		t.Errorf("status: got %q", params["status"])
	}

	// 无可选字段
	req2 := &CountStoriesRequest{WorkspaceID: "2"}
	params2 := req2.ToParams()
	if _, ok := params2["status"]; ok {
		t.Error("empty status should not be in params")
	}
}

// === bug.go ===

func TestListBugsRequest_ToParams(t *testing.T) {
	req := &ListBugsRequest{
		WorkspaceID:   "10",
		Title:         "crash",
		PriorityLabel: "urgent",
		Severity:      "fatal",
		Limit:         50,
	}
	params := req.ToParams()
	if params["workspace_id"] != "10" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["title"] != "crash" {
		t.Errorf("title: got %q", params["title"])
	}
	if params["priority_label"] != "urgent" {
		t.Errorf("priority_label: got %q", params["priority_label"])
	}
	if params["severity"] != "fatal" {
		t.Errorf("severity: got %q", params["severity"])
	}
	if params["limit"] != "50" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["status"]; ok {
		t.Error("empty status should not be in params")
	}
}

func TestCreateBugRequest_ToParams(t *testing.T) {
	req := &CreateBugRequest{
		WorkspaceID:  "10",
		Title:        "NPE",
		Severity:     "fatal",
		CurrentOwner: "alice",
		CustomFields: map[string]string{"custom_field_three": "cf3"},
	}
	params := req.ToParams()
	if params["workspace_id"] != "10" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["title"] != "NPE" {
		t.Errorf("title: got %q", params["title"])
	}
	if params["severity"] != "fatal" {
		t.Errorf("severity: got %q", params["severity"])
	}
	if params["current_owner"] != "alice" {
		t.Errorf("current_owner: got %q", params["current_owner"])
	}
	if params["custom_field_three"] != "cf3" {
		t.Errorf("custom_field_three: got %q", params["custom_field_three"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestUpdateBugRequest_ToParams(t *testing.T) {
	req := &UpdateBugRequest{
		WorkspaceID:  "10",
		ID:           "55",
		Status:       "resolved",
		CustomFields: map[string]string{"custom_field_four": "cf4"},
	}
	params := req.ToParams()
	if params["workspace_id"] != "10" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "55" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["status"] != "resolved" {
		t.Errorf("status: got %q", params["status"])
	}
	if params["custom_field_four"] != "cf4" {
		t.Errorf("custom_field_four: got %q", params["custom_field_four"])
	}
	if _, ok := params["title"]; ok {
		t.Error("empty title should not be in params")
	}
}

func TestCountBugsRequest_ToParams(t *testing.T) {
	req := &CountBugsRequest{
		WorkspaceID: "10",
		Status:      "new",
	}
	params := req.ToParams()
	if params["workspace_id"] != "10" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "new" {
		t.Errorf("status: got %q", params["status"])
	}
	if _, ok := params["title"]; ok {
		t.Error("empty title should not be in params")
	}
}

// === task.go ===

func TestListTasksRequest_ToParams(t *testing.T) {
	req := &ListTasksRequest{
		WorkspaceID: "20",
		Status:      "open",
		Owner:       "bob",
		Limit:       100,
	}
	params := req.ToParams()
	if params["workspace_id"] != "20" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "open" {
		t.Errorf("status: got %q", params["status"])
	}
	if params["owner"] != "bob" {
		t.Errorf("owner: got %q", params["owner"])
	}
	if params["limit"] != "100" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

func TestCreateTaskRequest_ToParams(t *testing.T) {
	req := &CreateTaskRequest{
		WorkspaceID:  "20",
		Name:         "Impl feature",
		Owner:        "bob",
		CustomFields: map[string]string{"custom_field_five": "cf5"},
	}
	params := req.ToParams()
	if params["workspace_id"] != "20" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "Impl feature" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["owner"] != "bob" {
		t.Errorf("owner: got %q", params["owner"])
	}
	if params["custom_field_five"] != "cf5" {
		t.Errorf("custom_field_five: got %q", params["custom_field_five"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestUpdateTaskRequest_ToParams(t *testing.T) {
	req := &UpdateTaskRequest{
		WorkspaceID:  "20",
		ID:           "77",
		Status:       "done",
		CustomFields: map[string]string{"custom_field_six": "cf6"},
	}
	params := req.ToParams()
	if params["workspace_id"] != "20" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "77" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["status"] != "done" {
		t.Errorf("status: got %q", params["status"])
	}
	if params["custom_field_six"] != "cf6" {
		t.Errorf("custom_field_six: got %q", params["custom_field_six"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

func TestCountTasksRequest_ToParams(t *testing.T) {
	req := &CountTasksRequest{
		WorkspaceID: "20",
		Status:      "progressing",
	}
	params := req.ToParams()
	if params["workspace_id"] != "20" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "progressing" {
		t.Errorf("status: got %q", params["status"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

// === iteration.go ===

func TestListIterationsRequest_ToParams(t *testing.T) {
	req := &ListIterationsRequest{
		WorkspaceID: "30",
		Status:      "open",
		Name:        "Sprint 1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "30" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "open" {
		t.Errorf("status: got %q", params["status"])
	}
	if params["name"] != "Sprint 1" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestCreateIterationRequest_ToParams(t *testing.T) {
	req := &CreateIterationRequest{
		WorkspaceID: "30",
		Name:        "Sprint 2",
		StartDate:   "2025-01-01",
		EndDate:     "2025-01-14",
		Creator:     "admin",
		Description: "test sprint",
	}
	params := req.ToParams()
	if params["workspace_id"] != "30" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "Sprint 2" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["startdate"] != "2025-01-01" {
		t.Errorf("startdate: got %q", params["startdate"])
	}
	if params["enddate"] != "2025-01-14" {
		t.Errorf("enddate: got %q", params["enddate"])
	}
	if params["creator"] != "admin" {
		t.Errorf("creator: got %q", params["creator"])
	}
	if params["description"] != "test sprint" {
		t.Errorf("description: got %q", params["description"])
	}
	if _, ok := params["status"]; ok {
		t.Error("empty status should not be in params")
	}
}

func TestUpdateIterationRequest_ToParams(t *testing.T) {
	req := &UpdateIterationRequest{
		WorkspaceID: "30",
		ID:          "88",
		CurrentUser: "admin",
		Name:        "Sprint 2 updated",
	}
	params := req.ToParams()
	if params["workspace_id"] != "30" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "88" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["current_user"] != "admin" {
		t.Errorf("current_user: got %q", params["current_user"])
	}
	if params["name"] != "Sprint 2 updated" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestCountIterationsRequest_ToParams(t *testing.T) {
	req := &CountIterationsRequest{
		WorkspaceID: "30",
		Status:      "done",
	}
	params := req.ToParams()
	if params["workspace_id"] != "30" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "done" {
		t.Errorf("status: got %q", params["status"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

// === tcase.go ===

func TestListTCasesRequest_ToParams(t *testing.T) {
	req := &ListTCasesRequest{
		WorkspaceID: "40",
		Name:        "login test",
		Priority:    "1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "40" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "login test" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["priority"] != "1" {
		t.Errorf("priority: got %q", params["priority"])
	}
	if _, ok := params["status"]; ok {
		t.Error("empty status should not be in params")
	}
}

func TestCountTCasesRequest_ToParams(t *testing.T) {
	req := &CountTCasesRequest{
		WorkspaceID: "40",
		Status:      "normal",
	}
	params := req.ToParams()
	if params["workspace_id"] != "40" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["status"] != "normal" {
		t.Errorf("status: got %q", params["status"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

func TestCreateTCaseRequest_ToParams(t *testing.T) {
	req := &CreateTCaseRequest{
		WorkspaceID: "40",
		Name:        "new case",
		Priority:    "2",
		Steps:       "step1;step2",
	}
	params := req.ToParams()
	if params["workspace_id"] != "40" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "new case" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["priority"] != "2" {
		t.Errorf("priority: got %q", params["priority"])
	}
	if params["steps"] != "step1;step2" {
		t.Errorf("steps: got %q", params["steps"])
	}
	if _, ok := params["status"]; ok {
		t.Error("empty status should not be in params")
	}
}

func TestBatchCreateTCasesRequest_ToJSON(t *testing.T) {
	req := &BatchCreateTCasesRequest{
		Items: []BatchCreateTCaseItem{
			{WorkspaceID: "40", Name: "case1"},
			{WorkspaceID: "40", Name: "case2", Priority: "high"},
		},
	}
	body, err := req.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() error: %v", err)
	}
	if len(body) == 0 {
		t.Error("expected non-empty JSON body")
	}

	// 空 items
	req2 := &BatchCreateTCasesRequest{}
	body2, err := req2.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() error: %v", err)
	}
	if string(body2) != "null" {
		t.Errorf("expected null for empty items, got %q", string(body2))
	}
}

// === comment.go ===

func TestListCommentsRequest_ToParams(t *testing.T) {
	req := &ListCommentsRequest{
		WorkspaceID: "50",
		EntryType:   "stories",
		Author:      "charlie",
	}
	params := req.ToParams()
	if params["workspace_id"] != "50" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["entry_type"] != "stories" {
		t.Errorf("entry_type: got %q", params["entry_type"])
	}
	if params["author"] != "charlie" {
		t.Errorf("author: got %q", params["author"])
	}
	if _, ok := params["title"]; ok {
		t.Error("empty title should not be in params")
	}
}

func TestAddCommentRequest_ToParams(t *testing.T) {
	req := &AddCommentRequest{
		WorkspaceID: "50",
		Description: "looks good",
		Author:      "charlie",
		EntryType:   "stories",
		EntryID:     "1001",
		RootID:      "500",
	}
	params := req.ToParams()
	if params["workspace_id"] != "50" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["description"] != "looks good" {
		t.Errorf("description: got %q", params["description"])
	}
	if params["author"] != "charlie" {
		t.Errorf("author: got %q", params["author"])
	}
	if params["entry_type"] != "stories" {
		t.Errorf("entry_type: got %q", params["entry_type"])
	}
	if params["entry_id"] != "1001" {
		t.Errorf("entry_id: got %q", params["entry_id"])
	}
	if params["root_id"] != "500" {
		t.Errorf("root_id: got %q", params["root_id"])
	}
	if _, ok := params["reply_id"]; ok {
		t.Error("empty reply_id should not be in params")
	}
}

func TestUpdateCommentRequest_ToParams(t *testing.T) {
	req := &UpdateCommentRequest{
		WorkspaceID: "50",
		ID:          "600",
		Description: "updated comment",
	}
	params := req.ToParams()
	if params["workspace_id"] != "50" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "600" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["description"] != "updated comment" {
		t.Errorf("description: got %q", params["description"])
	}
	if _, ok := params["change_creator"]; ok {
		t.Error("empty change_creator should not be in params")
	}
}

func TestCountCommentsRequest_ToParams(t *testing.T) {
	req := &CountCommentsRequest{
		WorkspaceID: "50",
		EntryType:   "bug",
	}
	params := req.ToParams()
	if params["workspace_id"] != "50" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["entry_type"] != "bug" {
		t.Errorf("entry_type: got %q", params["entry_type"])
	}
	if _, ok := params["author"]; ok {
		t.Error("empty author should not be in params")
	}
}

// === wiki.go ===

func TestListWikisRequest_ToParams(t *testing.T) {
	req := &ListWikisRequest{
		WorkspaceID: "60",
		Creator:     "dave",
		Limit:       20,
	}
	params := req.ToParams()
	if params["workspace_id"] != "60" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["creator"] != "dave" {
		t.Errorf("creator: got %q", params["creator"])
	}
	if params["limit"] != "20" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["name"]; ok {
		t.Error("empty name should not be in params")
	}
}

func TestCreateWikiRequest_ToParams(t *testing.T) {
	req := &CreateWikiRequest{
		WorkspaceID:         "60",
		Name:                "API Guide",
		Creator:             "dave",
		MarkdownDescription: "# Hello",
	}
	params := req.ToParams()
	if params["workspace_id"] != "60" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "API Guide" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["creator"] != "dave" {
		t.Errorf("creator: got %q", params["creator"])
	}
	if params["markdown_description"] != "# Hello" {
		t.Errorf("markdown_description: got %q", params["markdown_description"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestUpdateWikiRequest_ToParams(t *testing.T) {
	req := &UpdateWikiRequest{
		WorkspaceID: "60",
		ID:          "700",
		Name:        "Updated Guide",
	}
	params := req.ToParams()
	if params["workspace_id"] != "60" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "700" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "Updated Guide" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

// === timesheet.go ===

func TestListTimesheetsRequest_ToParams(t *testing.T) {
	req := &ListTimesheetsRequest{
		WorkspaceID: "70",
		EntityType:  "story",
		Owner:       "eve",
	}
	params := req.ToParams()
	if params["workspace_id"] != "70" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["entity_type"] != "story" {
		t.Errorf("entity_type: got %q", params["entity_type"])
	}
	if params["owner"] != "eve" {
		t.Errorf("owner: got %q", params["owner"])
	}
	if _, ok := params["entity_id"]; ok {
		t.Error("empty entity_id should not be in params")
	}
}

func TestAddTimesheetRequest_ToParams(t *testing.T) {
	req := &AddTimesheetRequest{
		WorkspaceID: "70",
		EntityType:  "task",
		EntityID:    "800",
		Timespent:   "3h",
		Owner:       "eve",
		Memo:        "coding",
	}
	params := req.ToParams()
	if params["workspace_id"] != "70" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["entity_type"] != "task" {
		t.Errorf("entity_type: got %q", params["entity_type"])
	}
	if params["entity_id"] != "800" {
		t.Errorf("entity_id: got %q", params["entity_id"])
	}
	if params["timespent"] != "3h" {
		t.Errorf("timespent: got %q", params["timespent"])
	}
	if params["owner"] != "eve" {
		t.Errorf("owner: got %q", params["owner"])
	}
	if params["memo"] != "coding" {
		t.Errorf("memo: got %q", params["memo"])
	}
	if _, ok := params["timeremain"]; ok {
		t.Error("empty timeremain should not be in params")
	}
}

func TestUpdateTimesheetRequest_ToParams(t *testing.T) {
	req := &UpdateTimesheetRequest{
		WorkspaceID: "70",
		ID:          "900",
		Timespent:   "5h",
	}
	params := req.ToParams()
	if params["workspace_id"] != "70" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "900" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["timespent"] != "5h" {
		t.Errorf("timespent: got %q", params["timespent"])
	}
	if _, ok := params["memo"]; ok {
		t.Error("empty memo should not be in params")
	}
}

// === attachment.go ===

func TestGetImageRequest_ToParams(t *testing.T) {
	req := &GetImageRequest{
		WorkspaceID: "80",
		ImagePath:   "/img/test.png",
	}
	params := req.ToParams()
	if params["workspace_id"] != "80" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["image_path"] != "/img/test.png" {
		t.Errorf("image_path: got %q", params["image_path"])
	}

	// 无 image_path
	req2 := &GetImageRequest{WorkspaceID: "80"}
	params2 := req2.ToParams()
	if _, ok := params2["image_path"]; ok {
		t.Error("empty image_path should not be in params")
	}
}

func TestGetAttachmentsRequest_ToParams(t *testing.T) {
	req := &GetAttachmentsRequest{
		WorkspaceID: "80",
		Type:        "story",
		EntryID:     "1001",
		Limit:       10,
	}
	params := req.ToParams()
	if params["workspace_id"] != "80" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["type"] != "story" {
		t.Errorf("type: got %q", params["type"])
	}
	if params["entry_id"] != "1001" {
		t.Errorf("entry_id: got %q", params["entry_id"])
	}
	if params["limit"] != "10" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["page"]; ok {
		t.Error("empty page should not be in params")
	}
}

// === request.go ===

func TestGetCustomFieldsRequest_ToParams(t *testing.T) {
	req := &GetCustomFieldsRequest{
		WorkspaceID: "90",
		EntityType:  "stories",
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

func TestWorkspaceIDRequest_ToParams(t *testing.T) {
	req := &WorkspaceIDRequest{
		WorkspaceID: "90",
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

func TestWorkflowRequest_ToParams(t *testing.T) {
	req := &WorkflowRequest{
		WorkspaceID:    "90",
		System:         "story",
		WorkitemTypeID: "wt1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["system"] != "story" {
		t.Errorf("system: got %q", params["system"])
	}
	if params["workitem_type_id"] != "wt1" {
		t.Errorf("workitem_type_id: got %q", params["workitem_type_id"])
	}

	// 无可选字段
	req2 := &WorkflowRequest{WorkspaceID: "90"}
	params2 := req2.ToParams()
	if _, ok := params2["system"]; ok {
		t.Error("empty system should not be in params")
	}
}

func TestGetCommitMsgRequest_ToParams(t *testing.T) {
	req := &GetCommitMsgRequest{
		WorkspaceID: "90",
		ObjectID:    "obj1",
		Type:        "story",
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["object_id"] != "obj1" {
		t.Errorf("object_id: got %q", params["object_id"])
	}
	if params["type"] != "story" {
		t.Errorf("type: got %q", params["type"])
	}
	if len(params) != 3 {
		t.Errorf("expected 3 params, got %d", len(params))
	}
}

func TestGetTodoRequest_ToParams(t *testing.T) {
	req := &GetTodoRequest{
		WorkspaceID: "90",
		EntityType:  "story",
		Limit:       50,
		Page:        2,
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["limit"] != "50" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if params["page"] != "2" {
		t.Errorf("page: got %q", params["page"])
	}
	if _, ok := params["order"]; ok {
		t.Error("empty order should not be in params")
	}
}

func TestGetRelatedBugsRequest_ToParams(t *testing.T) {
	req := &GetRelatedBugsRequest{
		WorkspaceID: "90",
		StoryID:     "s1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["story_id"] != "s1" {
		t.Errorf("story_id: got %q", params["story_id"])
	}
	if len(params) != 2 {
		t.Errorf("expected 2 params, got %d", len(params))
	}
}

func TestCreateRelationRequest_ToParams(t *testing.T) {
	req := &CreateRelationRequest{
		WorkspaceID: "90",
		SourceType:  "story",
		TargetType:  "bug",
		SourceID:    "s1",
		TargetID:    "b1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "90" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["source_type"] != "story" {
		t.Errorf("source_type: got %q", params["source_type"])
	}
	if params["target_type"] != "bug" {
		t.Errorf("target_type: got %q", params["target_type"])
	}
	if params["source_id"] != "s1" {
		t.Errorf("source_id: got %q", params["source_id"])
	}
	if params["target_id"] != "b1" {
		t.Errorf("target_id: got %q", params["target_id"])
	}
	if len(params) != 5 {
		t.Errorf("expected 5 params, got %d", len(params))
	}
}

// === change.go ===

func TestGetStoryChangesRequest_ToParams(t *testing.T) {
	req := &GetStoryChangesRequest{
		WorkspaceID: "100",
		StoryID:     "s1",
		Creator:     "alice",
		ChangeType:  "update",
		Limit:       30,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["story_id"] != "s1" {
		t.Errorf("story_id: got %q", params["story_id"])
	}
	if params["creator"] != "alice" {
		t.Errorf("creator: got %q", params["creator"])
	}
	if params["change_type"] != "update" {
		t.Errorf("change_type: got %q", params["change_type"])
	}
	if params["limit"] != "30" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
	if _, ok := params["page"]; ok {
		t.Error("empty page should not be in params")
	}
}

func TestCountStoryChangesRequest_ToParams(t *testing.T) {
	req := &CountStoryChangesRequest{
		WorkspaceID: "100",
		StoryID:     "s1",
		Creator:     "alice",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["story_id"] != "s1" {
		t.Errorf("story_id: got %q", params["story_id"])
	}
	if params["creator"] != "alice" {
		t.Errorf("creator: got %q", params["creator"])
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
}

func TestGetBugChangesRequest_ToParams(t *testing.T) {
	req := &GetBugChangesRequest{
		WorkspaceID: "200",
		BugID:       "b1",
		Author:      "bob",
		Field:       "severity",
		Limit:       50,
	}
	params := req.ToParams()
	if params["workspace_id"] != "200" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["bug_id"] != "b1" {
		t.Errorf("bug_id: got %q", params["bug_id"])
	}
	if params["author"] != "bob" {
		t.Errorf("author: got %q", params["author"])
	}
	if params["field"] != "severity" {
		t.Errorf("field: got %q", params["field"])
	}
	if params["limit"] != "50" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
	if _, ok := params["page"]; ok {
		t.Error("empty page should not be in params")
	}
}

func TestCountBugChangesRequest_ToParams(t *testing.T) {
	req := &CountBugChangesRequest{
		WorkspaceID: "200",
		BugID:       "b1",
		Field:       "status",
	}
	params := req.ToParams()
	if params["workspace_id"] != "200" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["bug_id"] != "b1" {
		t.Errorf("bug_id: got %q", params["bug_id"])
	}
	if params["field"] != "status" {
		t.Errorf("field: got %q", params["field"])
	}
	if _, ok := params["author"]; ok {
		t.Error("empty author should not be in params")
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
}

func TestGetTaskChangesRequest_ToParams(t *testing.T) {
	req := &GetTaskChangesRequest{
		WorkspaceID: "300",
		TaskID:      "t1",
		Creator:     "charlie",
		Limit:       20,
		Page:        2,
	}
	params := req.ToParams()
	if params["workspace_id"] != "300" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["task_id"] != "t1" {
		t.Errorf("task_id: got %q", params["task_id"])
	}
	if params["creator"] != "charlie" {
		t.Errorf("creator: got %q", params["creator"])
	}
	if params["limit"] != "20" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if params["page"] != "2" {
		t.Errorf("page: got %q", params["page"])
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
}

func TestCountTaskChangesRequest_ToParams(t *testing.T) {
	req := &CountTaskChangesRequest{
		WorkspaceID: "300",
		TaskID:      "t1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "300" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["task_id"] != "t1" {
		t.Errorf("task_id: got %q", params["task_id"])
	}
	if _, ok := params["creator"]; ok {
		t.Error("empty creator should not be in params")
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
}

func TestGetIterationChangesRequest_ToParams(t *testing.T) {
	req := &GetIterationChangesRequest{
		WorkspaceID: "400",
		IterationID: "i1",
		Author:      "dave",
		Field:       "name",
		Limit:       10,
	}
	params := req.ToParams()
	if params["workspace_id"] != "400" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["iteration_id"] != "i1" {
		t.Errorf("iteration_id: got %q", params["iteration_id"])
	}
	if params["author"] != "dave" {
		t.Errorf("author: got %q", params["author"])
	}
	if params["field"] != "name" {
		t.Errorf("field: got %q", params["field"])
	}
	if params["limit"] != "10" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
	if _, ok := params["page"]; ok {
		t.Error("empty page should not be in params")
	}
}

// === board.go ===

func TestCreateBoardCardRequest_ToParams(t *testing.T) {
	req := &CreateBoardCardRequest{
		WorkspaceID: "100",
		Name:        "新看板项",
		BoardID:     "2001",
		ColumnID:    "3001",
		Description: "test",
		Owner:       "admin",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "新看板项" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["b_board_id"] != "2001" {
		t.Errorf("b_board_id: got %q", params["b_board_id"])
	}
	if params["b_column_id"] != "3001" {
		t.Errorf("b_column_id: got %q", params["b_column_id"])
	}
	if params["description"] != "test" {
		t.Errorf("description: got %q", params["description"])
	}
	if params["owner"] != "admin" {
		t.Errorf("owner: got %q", params["owner"])
	}
	if _, ok := params["priority"]; ok {
		t.Error("empty priority should not be in params")
	}
}

func TestGetBoardCardsRequest_ToParams(t *testing.T) {
	req := &GetBoardCardsRequest{
		WorkspaceID: "100",
		BoardID:     "2001",
		Limit:       20,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["b_board_id"] != "2001" {
		t.Errorf("b_board_id: got %q", params["b_board_id"])
	}
	if params["limit"] != "20" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["b_column_id"]; ok {
		t.Error("empty b_column_id should not be in params")
	}
	if _, ok := params["owner"]; ok {
		t.Error("empty owner should not be in params")
	}
}

func TestUpdateBoardCardRequest_ToParams(t *testing.T) {
	req := &UpdateBoardCardRequest{
		WorkspaceID: "100",
		ID:          "1001",
		Name:        "已更新",
		ColumnID:    "3002",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "1001" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "已更新" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["b_column_id"] != "3002" {
		t.Errorf("b_column_id: got %q", params["b_column_id"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestGetBoardColumnsRequest_ToParams(t *testing.T) {
	req := &GetBoardColumnsRequest{
		WorkspaceID: "100",
		BoardID:     "2001",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["board_id"] != "2001" {
		t.Errorf("board_id: got %q", params["board_id"])
	}
	if len(params) != 2 {
		t.Errorf("expected 2 params, got %d", len(params))
	}
}

// === release.go ===

func TestCreateReleaseRequest_ToParams(t *testing.T) {
	req := &CreateReleaseRequest{
		WorkspaceID: "100",
		Name:        "v1.0",
		Description: "first release",
		StartDate:   "2026-01-01",
		EndDate:     "2026-02-01",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "v1.0" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["description"] != "first release" {
		t.Errorf("description: got %q", params["description"])
	}
	if params["startdate"] != "2026-01-01" {
		t.Errorf("startdate: got %q", params["startdate"])
	}
	if params["enddate"] != "2026-02-01" {
		t.Errorf("enddate: got %q", params["enddate"])
	}
}

func TestUpdateReleaseRequest_ToParams(t *testing.T) {
	req := &UpdateReleaseRequest{
		WorkspaceID: "100",
		ID:          "1001",
		Name:        "v1.1",
		Status:      "done",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "1001" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "v1.1" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["status"] != "done" {
		t.Errorf("status: got %q", params["status"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
	if _, ok := params["startdate"]; ok {
		t.Error("empty startdate should not be in params")
	}
}

// === request.go (new request structs) ===

func TestGetSubWorkspacesRequest_ToParams(t *testing.T) {
	req := &GetSubWorkspacesRequest{
		WorkspaceID: "100",
		TemplateID:  "tpl1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["template_id"] != "tpl1" {
		t.Errorf("template_id: got %q", params["template_id"])
	}

	req2 := &GetSubWorkspacesRequest{WorkspaceID: "100"}
	params2 := req2.ToParams()
	if _, ok := params2["template_id"]; ok {
		t.Error("empty template_id should not be in params")
	}
}

func TestListCompanyProjectsRequest_ToParams(t *testing.T) {
	req := &ListCompanyProjectsRequest{
		CompanyID:   "c1",
		Category:    "project",
		WithExtends: "1",
	}
	params := req.ToParams()
	if params["company_id"] != "c1" {
		t.Errorf("company_id: got %q", params["company_id"])
	}
	if params["category"] != "project" {
		t.Errorf("category: got %q", params["category"])
	}
	if params["with_extends"] != "1" {
		t.Errorf("with_extends: got %q", params["with_extends"])
	}
}

func TestGetWorkspaceUsersRequest_ToParams(t *testing.T) {
	req := &GetWorkspaceUsersRequest{
		WorkspaceID: "100",
		Fields:      "user,email",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["fields"] != "user,email" {
		t.Errorf("fields: got %q", params["fields"])
	}

	req2 := &GetWorkspaceUsersRequest{WorkspaceID: "100"}
	params2 := req2.ToParams()
	if _, ok := params2["fields"]; ok {
		t.Error("empty fields should not be in params")
	}
}

func TestAddWorkspaceMemberRequest_ToParams(t *testing.T) {
	req := &AddWorkspaceMemberRequest{
		WorkspaceID: "100",
		Nick:        "davidning",
		CompanyID:   "c1",
		RoleIDs:     "1000000000000000010,1000000000000000015",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["nick"] != "davidning" {
		t.Errorf("nick: got %q", params["nick"])
	}
	if params["company_id"] != "c1" {
		t.Errorf("company_id: got %q", params["company_id"])
	}
	if params["role_ids"] != "1000000000000000010,1000000000000000015" {
		t.Errorf("role_ids: got %q", params["role_ids"])
	}
}

func TestGetLifeTimesRequest_ToParams(t *testing.T) {
	req := &GetLifeTimesRequest{
		WorkspaceID: "100",
		EntityType:  "story",
		EntityID:    "2001",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["entity_type"] != "story" {
		t.Errorf("entity_type: got %q", params["entity_type"])
	}
	if params["entity_id"] != "2001" {
		t.Errorf("entity_id: got %q", params["entity_id"])
	}
	if len(params) != 3 {
		t.Errorf("expected 3 params, got %d", len(params))
	}
}

func TestAddCodeCommitInfoRequest_ToJSON(t *testing.T) {
	req := &AddCodeCommitInfoRequest{
		WorkspaceID: "100",
		Message:     "fix bug #100",
		Author:      "admin",
		CommitID:    "abc123",
		Files:       []string{"U main.go", "A new.go"},
		Repo:        "my-repo",
		RepoID:      "repo-001",
		CommitTime:  "2026-01-01 10:00:00",
		HookURL:     "https://example.com/hook",
		Ref:         "refs/heads/main",
	}
	data, err := req.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if m["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %v", m["workspace_id"])
	}
	if m["message"] != "fix bug #100" {
		t.Errorf("message: got %v", m["message"])
	}
	if m["author"] != "admin" {
		t.Errorf("author: got %v", m["author"])
	}
	if m["hook_url"] != "https://example.com/hook" {
		t.Errorf("hook_url: got %v", m["hook_url"])
	}
	if m["ref"] != "refs/heads/main" {
		t.Errorf("ref: got %v", m["ref"])
	}
	files, ok := m["files"].([]interface{})
	if !ok || len(files) != 2 {
		t.Errorf("files: got %v", m["files"])
	}
}

func TestGetCodeCommitInfosRequest_ToParams(t *testing.T) {
	req := &GetCodeCommitInfosRequest{
		WorkspaceID: "100",
		Limit:       20,
		Page:        2,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["limit"] != "20" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if params["page"] != "2" {
		t.Errorf("page: got %q", params["page"])
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
}

func TestGetOneAttachmentRequest_ToParams(t *testing.T) {
	req := &GetOneAttachmentRequest{
		WorkspaceID: "100",
		ID:          "att1",
		FileName:    "test.pdf",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "att1" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["file_name"] != "test.pdf" {
		t.Errorf("file_name: got %q", params["file_name"])
	}

	req2 := &GetOneAttachmentRequest{WorkspaceID: "100", ID: "att1"}
	params2 := req2.ToParams()
	if _, ok := params2["file_name"]; ok {
		t.Error("empty file_name should not be in params")
	}
}

func TestDownloadDocumentRequest_ToParams(t *testing.T) {
	req := &DownloadDocumentRequest{
		WorkspaceID: "100",
		ID:          "doc1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "doc1" {
		t.Errorf("id: got %q", params["id"])
	}
	if len(params) != 2 {
		t.Errorf("expected 2 params, got %d", len(params))
	}
}

// === timesheet.go (new) ===

func TestCountTimesheetsRequest_ToParams(t *testing.T) {
	req := &CountTimesheetsRequest{
		WorkspaceID: "70",
		EntityType:  "story",
		Owner:       "eve",
		Spentdate:   "2026-01-01",
	}
	params := req.ToParams()
	if params["workspace_id"] != "70" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["entity_type"] != "story" {
		t.Errorf("entity_type: got %q", params["entity_type"])
	}
	if params["owner"] != "eve" {
		t.Errorf("owner: got %q", params["owner"])
	}
	if params["spentdate"] != "2026-01-01" {
		t.Errorf("spentdate: got %q", params["spentdate"])
	}
	if _, ok := params["entity_id"]; ok {
		t.Error("empty entity_id should not be in params")
	}
	if _, ok := params["created"]; ok {
		t.Error("empty created should not be in params")
	}
}

// === wiki.go (new) ===

func TestCountWikisRequest_ToParams(t *testing.T) {
	req := &CountWikisRequest{
		WorkspaceID: "60",
	}
	params := req.ToParams()
	if params["workspace_id"] != "60" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

// === setting.go ===

func TestCreateModuleRequest_ToParams(t *testing.T) {
	req := &CreateModuleRequest{
		WorkspaceID: "100",
		Name:        "核心模块",
		Description: "核心功能模块",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "核心模块" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["description"] != "核心功能模块" {
		t.Errorf("description: got %q", params["description"])
	}
}

func TestUpdateModuleRequest_ToParams(t *testing.T) {
	req := &UpdateModuleRequest{
		WorkspaceID: "100",
		ID:          "1001",
		Name:        "更新模块",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "1001" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "更新模块" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestGetModulesRequest_ToParams(t *testing.T) {
	req := &GetModulesRequest{
		WorkspaceID: "100",
		Limit:       20,
		Page:        2,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["limit"] != "20" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if params["page"] != "2" {
		t.Errorf("page: got %q", params["page"])
	}
}

func TestCountModulesRequest_ToParams(t *testing.T) {
	req := &CountModulesRequest{WorkspaceID: "100"}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

func TestCreateVersionRequest_ToParams(t *testing.T) {
	req := &CreateVersionRequest{
		WorkspaceID: "100",
		Name:        "v1.0",
		Description: "第一版",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "v1.0" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["description"] != "第一版" {
		t.Errorf("description: got %q", params["description"])
	}
}

func TestUpdateVersionRequest_ToParams(t *testing.T) {
	req := &UpdateVersionRequest{
		WorkspaceID: "100",
		ID:          "1001",
		Name:        "v2.0",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "1001" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "v2.0" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestGetVersionsRequest_ToParams(t *testing.T) {
	req := &GetVersionsRequest{
		WorkspaceID: "100",
		Limit:       10,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["limit"] != "10" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["page"]; ok {
		t.Error("empty page should not be in params")
	}
}

func TestCountVersionsRequest_ToParams(t *testing.T) {
	req := &CountVersionsRequest{WorkspaceID: "100"}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

func TestCreateBaselineRequest_ToParams(t *testing.T) {
	req := &CreateBaselineRequest{
		WorkspaceID: "100",
		Name:        "基线1",
		Description: "测试基线",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "基线1" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["description"] != "测试基线" {
		t.Errorf("description: got %q", params["description"])
	}
}

func TestUpdateBaselineRequest_ToParams(t *testing.T) {
	req := &UpdateBaselineRequest{
		WorkspaceID: "100",
		ID:          "1001",
		Name:        "基线2",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "1001" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "基线2" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestGetBaselinesRequest_ToParams(t *testing.T) {
	req := &GetBaselinesRequest{
		WorkspaceID: "100",
		Limit:       30,
		Page:        1,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["limit"] != "30" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if params["page"] != "1" {
		t.Errorf("page: got %q", params["page"])
	}
}

func TestCountBaselinesRequest_ToParams(t *testing.T) {
	req := &CountBaselinesRequest{WorkspaceID: "100"}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

func TestCreateFeatureRequest_ToParams(t *testing.T) {
	req := &CreateFeatureRequest{
		WorkspaceID: "100",
		Name:        "特性1",
		Description: "测试特性",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["name"] != "特性1" {
		t.Errorf("name: got %q", params["name"])
	}
	if params["description"] != "测试特性" {
		t.Errorf("description: got %q", params["description"])
	}
}

func TestUpdateFeatureRequest_ToParams(t *testing.T) {
	req := &UpdateFeatureRequest{
		WorkspaceID: "100",
		ID:          "1001",
		Name:        "特性2",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "1001" {
		t.Errorf("id: got %q", params["id"])
	}
	if params["name"] != "特性2" {
		t.Errorf("name: got %q", params["name"])
	}
	if _, ok := params["description"]; ok {
		t.Error("empty description should not be in params")
	}
}

func TestGetFeaturesRequest_ToParams(t *testing.T) {
	req := &GetFeaturesRequest{
		WorkspaceID: "100",
		Limit:       50,
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["limit"] != "50" {
		t.Errorf("limit: got %q", params["limit"])
	}
	if _, ok := params["page"]; ok {
		t.Error("empty page should not be in params")
	}
}

func TestCountFeaturesRequest_ToParams(t *testing.T) {
	req := &CountFeaturesRequest{WorkspaceID: "100"}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if len(params) != 1 {
		t.Errorf("expected 1 param, got %d", len(params))
	}
}

func TestIterationLockRequest_ToParams(t *testing.T) {
	req := &IterationLockRequest{
		WorkspaceID: "100",
		ID:          "iter1",
	}
	params := req.ToParams()
	if params["workspace_id"] != "100" {
		t.Errorf("workspace_id: got %q", params["workspace_id"])
	}
	if params["id"] != "iter1" {
		t.Errorf("id: got %q", params["id"])
	}
	if len(params) != 2 {
		t.Errorf("expected 2 params, got %d", len(params))
	}
}
