package tapd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestListIterations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations" {
			t.Errorf("unexpected path: %s, want /iterations", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"Iteration":{"id":"3001","name":"Sprint 1","status":"open","startdate":"2026-04-01","enddate":"2026-04-15"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	iterations, err := c.ListIterations(context.Background(), &model.ListIterationsRequest{
		WorkspaceID: "1",
	})
	if err != nil {
		t.Fatalf("ListIterations() unexpected error: %v", err)
	}
	if len(iterations) != 1 {
		t.Fatalf("expected 1 iteration, got %d", len(iterations))
	}
	iter := iterations[0]
	if iter.ID != "3001" {
		t.Errorf("iteration ID = %q, want %q", iter.ID, "3001")
	}
	if iter.Name != "Sprint 1" {
		t.Errorf("iteration Name = %q, want %q", iter.Name, "Sprint 1")
	}
	if iter.Status != "open" {
		t.Errorf("iteration Status = %q, want %q", iter.Status, "open")
	}
	if iter.StartDate != "2026-04-01" {
		t.Errorf("iteration StartDate = %q, want %q", iter.StartDate, "2026-04-01")
	}
	if iter.EndDate != "2026-04-15" {
		t.Errorf("iteration EndDate = %q, want %q", iter.EndDate, "2026-04-15")
	}
}

func TestCreateIteration(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/iterations" {
			t.Errorf("unexpected path: %s, want /iterations", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Iteration":{"id":"3002","name":"Sprint 2","workspace_id":"1","startdate":"2026-04-16","enddate":"2026-04-30","status":"open","creator":"testuser","created":"2026-04-16 10:00:00"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	iteration, err := c.CreateIteration(context.Background(), &model.CreateIterationRequest{
		WorkspaceID: "1",
		Name:        "Sprint 2",
		StartDate:   "2026-04-16",
		EndDate:     "2026-04-30",
		Creator:     "testuser",
	})
	if err != nil {
		t.Fatalf("CreateIteration() unexpected error: %v", err)
	}
	if iteration.ID != "3002" {
		t.Errorf("ID = %q, want %q", iteration.ID, "3002")
	}
	if iteration.Name != "Sprint 2" {
		t.Errorf("Name = %q, want %q", iteration.Name, "Sprint 2")
	}
	if iteration.WorkspaceID != "1" {
		t.Errorf("WorkspaceID = %q, want %q", iteration.WorkspaceID, "1")
	}
}

func TestUpdateIteration(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/iterations" {
			t.Errorf("unexpected path: %s, want /iterations", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Iteration":{"id":"3001","name":"Sprint 1 Updated","workspace_id":"1","startdate":"2026-04-01","enddate":"2026-04-15","status":"done","description":"updated desc","creator":"testuser","modified":"2026-04-16 12:00:00"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.UpdateIteration(context.Background(), &model.UpdateIterationRequest{
		WorkspaceID: "1",
		ID:          "3001",
		CurrentUser: "testuser",
		Name:        "Sprint 1 Updated",
		Status:      "done",
		Description: "updated desc",
	})
	if err != nil {
		t.Fatalf("UpdateIteration() unexpected error: %v", err)
	}
	if result.ID != "3001" {
		t.Errorf("ID = %q, want %q", result.ID, "3001")
	}
	if result.Name != "Sprint 1 Updated" {
		t.Errorf("Name = %q, want %q", result.Name, "Sprint 1 Updated")
	}
	if result.Status != "done" {
		t.Errorf("Status = %q, want %q", result.Status, "done")
	}
	if result.Description != "updated desc" {
		t.Errorf("Description = %q, want %q", result.Description, "updated desc")
	}
}

func TestCountIterations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/count" {
			t.Errorf("unexpected path: %s, want /iterations/count", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"count":8},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountIterations(context.Background(), &model.CountIterationsRequest{
		WorkspaceID: "1",
	})
	if err != nil {
		t.Fatalf("CountIterations() unexpected error: %v", err)
	}
	if count != 8 {
		t.Errorf("count = %d, want 8", count)
	}
}

func TestCreateIteration_WithAllFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/iterations" {
			t.Errorf("unexpected path: %s, want /iterations", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Iteration":{"id":"iter_200","name":"Sprint 5","workspace_id":"1","startdate":"2026-04-01","enddate":"2026-04-15","creator":"admin","status":"open"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	iteration, err := c.CreateIteration(context.Background(), &model.CreateIterationRequest{
		WorkspaceID: "1",
		Name:        "Sprint 5",
		StartDate:   "2026-04-01",
		EndDate:     "2026-04-15",
		Creator:     "admin",
	})
	if err != nil {
		t.Fatalf("CreateIteration() unexpected error: %v", err)
	}
	if iteration.ID != "iter_200" {
		t.Errorf("ID = %q, want %q", iteration.ID, "iter_200")
	}
	if iteration.Name != "Sprint 5" {
		t.Errorf("Name = %q, want %q", iteration.Name, "Sprint 5")
	}
	if iteration.StartDate != "2026-04-01" {
		t.Errorf("StartDate = %q, want %q", iteration.StartDate, "2026-04-01")
	}
	if iteration.EndDate != "2026-04-15" {
		t.Errorf("EndDate = %q, want %q", iteration.EndDate, "2026-04-15")
	}
	if iteration.Creator != "admin" {
		t.Errorf("Creator = %q, want %q", iteration.Creator, "admin")
	}
	if iteration.Status != "open" {
		t.Errorf("Status = %q, want %q", iteration.Status, "open")
	}
}

func TestLockIteration(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/iterations/lock_iteration" {
			t.Errorf("unexpected path: %s, want /iterations/lock_iteration", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":"success","info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.LockIteration(context.Background(), &model.LockIterationRequest{
		WorkspaceID: "1",
		IterationID: "3001",
	})
	if err != nil {
		t.Fatalf("LockIteration() unexpected error: %v", err)
	}
}

func TestUnlockIteration(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/iterations/unlock_iteration" {
			t.Errorf("unexpected path: %s, want /iterations/unlock_iteration", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":"success","info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.UnlockIteration(context.Background(), &model.UnlockIterationRequest{
		WorkspaceID: "1",
		IterationID: "3001",
	})
	if err != nil {
		t.Fatalf("UnlockIteration() unexpected error: %v", err)
	}
}

func TestUpdateIteration_StatusChange(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/iterations" {
			t.Errorf("unexpected path: %s, want /iterations", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Iteration":{"id":"iter_100","name":"Sprint 4 Updated","status":"done"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.UpdateIteration(context.Background(), &model.UpdateIterationRequest{
		WorkspaceID: "1",
		ID:          "iter_100",
		CurrentUser: "admin",
		Name:        "Sprint 4 Updated",
		Status:      "done",
	})
	if err != nil {
		t.Fatalf("UpdateIteration() unexpected error: %v", err)
	}
	if result.ID != "iter_100" {
		t.Errorf("ID = %q, want %q", result.ID, "iter_100")
	}
	if result.Name != "Sprint 4 Updated" {
		t.Errorf("Name = %q, want %q", result.Name, "Sprint 4 Updated")
	}
	if result.Status != "done" {
		t.Errorf("Status = %q, want %q", result.Status, "done")
	}
}

func TestGetCustomDashBoardContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/get_custom_dash_board_content" {
			t.Errorf("unexpected path: %s, want /iterations/get_custom_dash_board_content", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"id":"1010104801000003949","template":"Custom","title":"自定义aaa","component_data":"[]","width":"6","height":"3","card_type":"RichContent","data":{"content":"<p>自定义卡片内容。支持 <strong>HTML</strong>。</p>","description_type":"1","value":"<p>自定义卡片内容。支持 <strong>HTML</strong>。</p>"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	cards, err := c.GetCustomDashBoardContent(context.Background(), &model.GetCustomDashBoardContentRequest{
		WorkspaceID: "10104801",
		IterationID: "1010104801000723579",
	})
	if err != nil {
		t.Fatalf("GetCustomDashBoardContent() unexpected error: %v", err)
	}
	if len(cards) != 1 {
		t.Fatalf("expected 1 card, got %d", len(cards))
	}
	if cards[0].ID != "1010104801000003949" {
		t.Errorf("card ID = %q, want %q", cards[0].ID, "1010104801000003949")
	}
	if cards[0].Title != "自定义aaa" {
		t.Errorf("card Title = %q, want %q", cards[0].Title, "自定义aaa")
	}
	if cards[0].CardType != "RichContent" {
		t.Errorf("card CardType = %q, want %q", cards[0].CardType, "RichContent")
	}
}

func TestUpdateCustomDashBoardContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/update_custom_dash_board_content" {
			t.Errorf("unexpected path: %s, want /iterations/update_custom_dash_board_content", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("ParseForm() error: %v", err)
		}
		if v := r.FormValue("workspace_id"); v != "10104801" {
			t.Errorf("workspace_id = %q, want %q", v, "10104801")
		}
		if v := r.FormValue("iteration_id"); v != "1010104801000723579" {
			t.Errorf("iteration_id = %q, want %q", v, "1010104801000723579")
		}
		if v := r.FormValue("card_id"); v != "1010104801000003949" {
			t.Errorf("card_id = %q, want %q", v, "1010104801000003949")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"id":"1010104801000003949"},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	id, err := c.UpdateCustomDashBoardContent(context.Background(), &model.UpdateCustomDashBoardContentRequest{
		WorkspaceID: "10104801",
		IterationID: "1010104801000723579",
		CardID:      "1010104801000003949",
		Content:     "<p>updated</p>",
	})
	if err != nil {
		t.Fatalf("UpdateCustomDashBoardContent() unexpected error: %v", err)
	}
	if id != "1010104801000003949" {
		t.Errorf("id = %q, want %q", id, "1010104801000003949")
	}
}

func TestGetIterationTemplateList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/template_list" {
			t.Errorf("unexpected path: %s, want /iterations/template_list", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"WorkitemTemplate":{"id":"1020375553000077579","workspace_id":"20375553","type":"iteration","name":"迭代模板1","creator":"SYSTEM","created":"2021-07-06 10:55:29","modified":"2021-07-06 10:55:29"}},{"WorkitemTemplate":{"id":"1020375553000091187","workspace_id":"20375553","type":"release","name":"发布计划模板1","creator":"TAPD system","created":"2023-03-23 15:23:35","modified":"2023-12-13 15:46:41"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	templates, err := c.GetIterationTemplateList(context.Background(), &model.WorkspaceIDRequest{
		WorkspaceID: "20375553",
	})
	if err != nil {
		t.Fatalf("GetIterationTemplateList() unexpected error: %v", err)
	}
	if len(templates) != 2 {
		t.Fatalf("expected 2 templates, got %d", len(templates))
	}
	if templates[0].ID != "1020375553000077579" {
		t.Errorf("template ID = %q, want %q", templates[0].ID, "1020375553000077579")
	}
	if templates[0].Name != "迭代模板1" {
		t.Errorf("template Name = %q, want %q", templates[0].Name, "迭代模板1")
	}
}

func TestGetIterationTemplateFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/template_fields" {
			t.Errorf("unexpected path: %s, want /iterations/template_fields", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"WorkitemTemplateField":{"id":"1020375553001067379","workspace_id":"20375553","type":"iteration","template_id":"1020375553000077579","field":"description","value":"","required":"1","sort":"0"}},{"WorkitemTemplateField":{"id":"1020375553001067381","workspace_id":"20375553","type":"iteration","template_id":"1020375553000077579","field":"name","value":"","required":"1","sort":"0"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	fields, err := c.GetIterationTemplateFields(context.Background(), &model.GetIterationTemplateFieldsRequest{
		WorkspaceID: "20375553",
		TemplateID:  "1020375553000077579",
	})
	if err != nil {
		t.Fatalf("GetIterationTemplateFields() unexpected error: %v", err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[0].Field != "description" {
		t.Errorf("field Field = %q, want %q", fields[0].Field, "description")
	}
	if fields[0].Required != "1" {
		t.Errorf("field Required = %q, want %q", fields[0].Required, "1")
	}
}

func TestGetIterationCustomFieldsSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/custom_fields_settings" {
			t.Errorf("unexpected path: %s, want /iterations/custom_fields_settings", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"CustomFieldConfig":{"id":"1010158231214902319","workspace_id":"10158231","entry_type":"iteration","custom_field":"custom_field_50","type":"text","name":"倒计时","options":null,"enabled":"1","sort":null}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	configs, err := c.GetIterationCustomFieldsSettings(context.Background(), &model.WorkspaceIDRequest{
		WorkspaceID: "10158231",
	})
	if err != nil {
		t.Fatalf("GetIterationCustomFieldsSettings() unexpected error: %v", err)
	}
	if len(configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(configs))
	}
	if configs[0].ID != "1010158231214902319" {
		t.Errorf("config ID = %q, want %q", configs[0].ID, "1010158231214902319")
	}
	if configs[0].Name != "倒计时" {
		t.Errorf("config Name = %q, want %q", configs[0].Name, "倒计时")
	}
}

func TestGetIterationWorkitemTypes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/workitem_types" {
			t.Errorf("unexpected path: %s, want /iterations/workitem_types", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"WorkitemType":{"id":"1020375553000072217","workspace_id":"20375553","entity_type":"release","name":"发布计划类别1","creator":"TAPD system","created":"2023-03-23 15:23:35","modified":"2023-12-13 15:46:41"}},{"WorkitemType":{"id":"1020375553000070695","workspace_id":"20375553","entity_type":"iteration","name":"迭代类别1","creator":"TAPD system","created":"2022-12-13 15:06:20","modified":"2022-12-13 15:06:20"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	types, err := c.GetIterationWorkitemTypes(context.Background(), &model.WorkspaceIDRequest{
		WorkspaceID: "20375553",
	})
	if err != nil {
		t.Fatalf("GetIterationWorkitemTypes() unexpected error: %v", err)
	}
	if len(types) != 2 {
		t.Fatalf("expected 2 types, got %d", len(types))
	}
	if types[0].ID != "1020375553000072217" {
		t.Errorf("type ID = %q, want %q", types[0].ID, "1020375553000072217")
	}
	if types[0].Name != "发布计划类别1" {
		t.Errorf("type Name = %q, want %q", types[0].Name, "发布计划类别1")
	}
}

func TestGetDefaultTemplateFieldsByWorkitemTypeID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/iterations/default_template_fields_by_workitem_type_id" {
			t.Errorf("unexpected path: %s, want /iterations/default_template_fields_by_workitem_type_id", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"WorkitemTemplateField":{"id":"1020375553001067379","workspace_id":"20375553","type":"iteration","template_id":"1020375553000077579","field":"description","value":"","required":"1","sort":"0"}},{"WorkitemTemplateField":{"id":"1020375553001067381","workspace_id":"20375553","type":"iteration","template_id":"1020375553000077579","field":"name","value":"","required":"1","sort":"0"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	fields, err := c.GetDefaultTemplateFieldsByWorkitemTypeID(context.Background(), &model.GetDefaultTemplateFieldsByWorkitemTypeIDRequest{
		WorkspaceID:    "20375553",
		WorkitemTypeID: "1020375553000070695",
	})
	if err != nil {
		t.Fatalf("GetDefaultTemplateFieldsByWorkitemTypeID() unexpected error: %v", err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[1].Field != "name" {
		t.Errorf("field Field = %q, want %q", fields[1].Field, "name")
	}
}

func TestGetPlanApps(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/plan_apps" {
			t.Errorf("unexpected path: %s, want /plan_apps", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"PlanApp":{"id":"1000000755000003485","name":"月度计划","workspace_id":"755","plan_id_field":"custom_plan_field_4","creator":"robertyang","created":"2023-07-25 12:15:13","modifier":"robertyang","modified":"2023-07-25 12:15:13"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	apps, err := c.GetPlanApps(context.Background(), &model.GetPlanAppsRequest{
		WorkspaceID: "755",
	})
	if err != nil {
		t.Fatalf("GetPlanApps() unexpected error: %v", err)
	}
	if len(apps) != 1 {
		t.Fatalf("expected 1 app, got %d", len(apps))
	}
	if apps[0].ID != "1000000755000003485" {
		t.Errorf("app ID = %q, want %q", apps[0].ID, "1000000755000003485")
	}
	if apps[0].Name != "月度计划" {
		t.Errorf("app Name = %q, want %q", apps[0].Name, "月度计划")
	}
	if apps[0].PlanIDField != "custom_plan_field_4" {
		t.Errorf("app PlanIDField = %q, want %q", apps[0].PlanIDField, "custom_plan_field_4")
	}
}

func TestCountPlanApps(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/plan_apps/count" {
			t.Errorf("unexpected path: %s, want /plan_apps/count", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"count":2},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountPlanApps(context.Background(), &model.CountPlanAppsRequest{
		WorkspaceID: "755",
	})
	if err != nil {
		t.Fatalf("CountPlanApps() unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}
