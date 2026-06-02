package tapd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestListTCases(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcases" {
			t.Errorf("unexpected path: %s, want /tcases", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Tcase":{"id":"1001","name":"TC1","workspace_id":"1","status":"normal","priority":"high","creator":"tester1","created":"2026-03-10"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListTCasesRequest{
		WorkspaceID: "1",
	}
	tcases, err := c.ListTCases(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTCases() unexpected error: %v", err)
	}
	if len(tcases) != 1 {
		t.Fatalf("expected 1 tcase, got %d", len(tcases))
	}
	if tcases[0].ID != "1001" {
		t.Errorf("tcase id = %q, want %q", tcases[0].ID, "1001")
	}
	if tcases[0].Name != "TC1" {
		t.Errorf("tcase name = %q, want %q", tcases[0].Name, "TC1")
	}
	if tcases[0].Status != "normal" {
		t.Errorf("tcase status = %q, want %q", tcases[0].Status, "normal")
	}
	if tcases[0].Priority != "high" {
		t.Errorf("tcase priority = %q, want %q", tcases[0].Priority, "high")
	}
	if tcases[0].Creator != "tester1" {
		t.Errorf("tcase creator = %q, want %q", tcases[0].Creator, "tester1")
	}
}

func TestCountTCases(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcases/count" {
			t.Errorf("unexpected path: %s, want /tcases/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":25},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountTCasesRequest{
		WorkspaceID: "1",
	}
	count, err := c.CountTCases(context.Background(), req)
	if err != nil {
		t.Fatalf("CountTCases() unexpected error: %v", err)
	}
	if count != 25 {
		t.Errorf("count = %d, want 25", count)
	}
}

func TestCreateTCase(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcases" {
			t.Errorf("unexpected path: %s, want /tcases", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Tcase":{"id":"1002","name":"New TC","workspace_id":"1","status":"normal"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateTCaseRequest{
		WorkspaceID: "1",
		Name:        "New TC",
	}
	tc, err := c.CreateTCase(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateTCase() unexpected error: %v", err)
	}
	if tc.ID != "1002" {
		t.Errorf("tcase id = %q, want %q", tc.ID, "1002")
	}
	if tc.Name != "New TC" {
		t.Errorf("tcase name = %q, want %q", tc.Name, "New TC")
	}
}

func TestBatchCreateTCases(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcases/batch_save" {
			t.Errorf("unexpected path: %s, want /tcases/batch_save", r.URL.Path)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Tcase":{"id":"1001","name":"TC1"}},{"Tcase":{"id":"1002","name":"TC2"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.BatchCreateTCasesRequest{
		Items: []model.BatchCreateTCaseItem{
			{WorkspaceID: "1", Name: "TC1"},
			{WorkspaceID: "1", Name: "TC2"},
		},
	}
	result, err := c.BatchCreateTCases(context.Background(), req)
	if err != nil {
		t.Fatalf("BatchCreateTCases() unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 tcases, got %d", len(result))
	}
}

func TestGetTCaseFieldsInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcases/get_fields_info" {
			t.Errorf("unexpected path: %s, want /tcases/get_fields_info", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "1" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "1")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"id":{"html_type":"input","label":"ID","options":[],"name":"id"},"name":{"html_type":"input","label":"用例名称","options":[],"name":"name"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetTCaseFieldsInfo(context.Background(), &model.WorkspaceIDRequest{WorkspaceID: "1"})
	if err != nil {
		t.Fatalf("GetTCaseFieldsInfo() unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(result))
	}
	if result["id"].Label != "ID" {
		t.Errorf("id label = %q, want %q", result["id"].Label, "ID")
	}
	if result["name"].Label != "用例名称" {
		t.Errorf("name label = %q, want %q", result["name"].Label, "用例名称")
	}
}

func TestUpdateTCase(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcases" {
			t.Errorf("unexpected path: %s, want /tcases", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"Tcase":{"id":"1010158231077224799","name":"测试浏览器兼容性","workspace_id":"10158231","status":"abandon","priority":"高"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.UpdateTCaseRequest{
		WorkspaceID: "10158231",
		ID:          "1010158231077224799",
		Name:        "测试浏览器兼容性",
	}
	tc, err := c.UpdateTCase(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateTCase() unexpected error: %v", err)
	}
	if tc.ID != "1010158231077224799" {
		t.Errorf("tcase id = %q, want %q", tc.ID, "1010158231077224799")
	}
	if tc.Name != "测试浏览器兼容性" {
		t.Errorf("tcase name = %q, want %q", tc.Name, "测试浏览器兼容性")
	}
	if tc.Status != "abandon" {
		t.Errorf("tcase status = %q, want %q", tc.Status, "abandon")
	}
}

func TestListTCaseCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcase_categories" {
			t.Errorf("unexpected path: %s, want /tcase_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"TcaseCategory":{"id":"1010158231075917759","workspace_id":"10158231","name":"None Category","description":"未规划目录","parent_id":"0","created":"2019-06-26 16:42:50"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListTCaseCategoriesRequest{WorkspaceID: "10158231"}
	categories, err := c.ListTCaseCategories(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTCaseCategories() unexpected error: %v", err)
	}
	if len(categories) != 1 {
		t.Fatalf("expected 1 category, got %d", len(categories))
	}
	if categories[0].ID != "1010158231075917759" {
		t.Errorf("category id = %q, want %q", categories[0].ID, "1010158231075917759")
	}
	if categories[0].Name != "None Category" {
		t.Errorf("category name = %q, want %q", categories[0].Name, "None Category")
	}
}

func TestCountTCaseCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcase_categories/count" {
			t.Errorf("unexpected path: %s, want /tcase_categories/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"count":4},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListTCaseCategoriesRequest{WorkspaceID: "10158231"}
	count, err := c.CountTCaseCategories(context.Background(), req)
	if err != nil {
		t.Fatalf("CountTCaseCategories() unexpected error: %v", err)
	}
	if count != 4 {
		t.Errorf("count = %d, want 4", count)
	}
}

func TestCreateTCaseCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcase_categories" {
			t.Errorf("unexpected path: %s, want /tcase_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"TcaseCategory":{"id":"1020355782075922101","workspace_id":"20355782","name":"test","parent_id":"0","created":"2020-05-26 15:04:19","creator":"v_xuanfang"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateTCaseCategoryRequest{
		WorkspaceID: "20355782",
		Name:        "test",
	}
	cat, err := c.CreateTCaseCategory(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateTCaseCategory() unexpected error: %v", err)
	}
	if cat.ID != "1020355782075922101" {
		t.Errorf("category id = %q, want %q", cat.ID, "1020355782075922101")
	}
	if cat.Name != "test" {
		t.Errorf("category name = %q, want %q", cat.Name, "test")
	}
	if cat.Creator != "v_xuanfang" {
		t.Errorf("category creator = %q, want %q", cat.Creator, "v_xuanfang")
	}
}

func TestGetStoryByTCaseID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcases/get_story_by_tcase_id" {
			t.Errorf("unexpected path: %s, want /tcases/get_story_by_tcase_id", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"workspace_id":"20358306","tcase_id":"1020358306077237053","story_id":"1020358306854812395","test_plan_id":"0"}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetStoryByTCaseIDRequest{
		WorkspaceID: "20358306",
		TCaseIDs:    "1020358306077237053",
	}
	relations, err := c.GetStoryByTCaseID(context.Background(), req)
	if err != nil {
		t.Fatalf("GetStoryByTCaseID() unexpected error: %v", err)
	}
	if len(relations) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(relations))
	}
	if relations[0].TCaseID != "1020358306077237053" {
		t.Errorf("tcase_id = %q, want %q", relations[0].TCaseID, "1020358306077237053")
	}
	if relations[0].StoryID != "1020358306854812395" {
		t.Errorf("story_id = %q, want %q", relations[0].StoryID, "1020358306854812395")
	}
}

func TestGetTCaseCustomFieldsSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcases/custom_fields_settings" {
			t.Errorf("unexpected path: %s, want /tcases/custom_fields_settings", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"CustomFieldConfig":{"id":"1000000755214854654","workspace_id":"755","entry_type":"tcase","custom_field":"custom_field_30","type":"select","name":"AT已实现？","enabled":"1"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	configs, err := c.GetTCaseCustomFieldsSettings(context.Background(), &model.WorkspaceIDRequest{WorkspaceID: "755"})
	if err != nil {
		t.Fatalf("GetTCaseCustomFieldsSettings() unexpected error: %v", err)
	}
	if len(configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(configs))
	}
	if configs[0].ID != "1000000755214854654" {
		t.Errorf("config id = %q, want %q", configs[0].ID, "1000000755214854654")
	}
	if configs[0].Name != "AT已实现？" {
		t.Errorf("config name = %q, want %q", configs[0].Name, "AT已实现？")
	}
	if configs[0].CustomField != "custom_field_30" {
		t.Errorf("config custom_field = %q, want %q", configs[0].CustomField, "custom_field_30")
	}
}

func TestAssignTCaseInstance(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcase_instance/assign" {
			t.Errorf("unexpected path: %s, want /tcase_instance/assign", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.AssignTCaseInstanceRequest{
		WorkspaceID: "10158231",
		TestPlanID:  "1010158231077233617",
		TCaseID:     "1010158231077224799",
		Executor:    "tester1",
	}
	err := c.AssignTCaseInstance(context.Background(), req)
	if err != nil {
		t.Fatalf("AssignTCaseInstance() unexpected error: %v", err)
	}
}

func TestExecuteTCaseInstance(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcase_instance/execute" {
			t.Errorf("unexpected path: %s, want /tcase_instance/execute", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ExecuteTCaseInstanceRequest{
		WorkspaceID:  "10158231",
		TestPlanID:   "1010158231077233617",
		TCaseID:      "1010158231077224799",
		ResultStatus: "pass",
		LastExecutor: "tester1",
	}
	err := c.ExecuteTCaseInstance(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecuteTCaseInstance() unexpected error: %v", err)
	}
}

func TestRemoveTCaseInstance(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcase_instance/remove_tcase" {
			t.Errorf("unexpected path: %s, want /tcase_instance/remove_tcase", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":null,"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.RemoveTCaseInstanceRequest{
		WorkspaceID: "10158231",
		TestPlanID:  "1010158231077233617",
		TCaseID:     "1010158231077224799",
	}
	err := c.RemoveTCaseInstance(context.Background(), req)
	if err != nil {
		t.Fatalf("RemoveTCaseInstance() unexpected error: %v", err)
	}
}

func TestGetTCaseResult(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tcase_instance/result" {
			t.Errorf("unexpected path: %s, want /tcase_instance/result", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"1020357849000703565":{"id":"1020357849000703565","executed_at":"2020-03-06 17:46:13","executor":"jeffjffang","result_status":"pass","result_remark":null,"bug_id":[],"Bug":[]}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TCaseResultRequest{
		WorkspaceID: "20357849",
		TestPlanID:  "1020357849077233617",
		TCaseID:     "1020357849000703565",
	}
	result, err := c.GetTCaseResult(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTCaseResult() unexpected error: %v", err)
	}
	item, ok := result["1020357849000703565"]
	if !ok {
		t.Fatal("expected key 1020357849000703565 in result")
	}
	if item.ResultStatus != "pass" {
		t.Errorf("result_status = %q, want %q", item.ResultStatus, "pass")
	}
	if item.Executor != "jeffjffang" {
		t.Errorf("executor = %q, want %q", item.Executor, "jeffjffang")
	}
}

func TestDeleteTCaseStoryRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/tcase_instance/delete_tcase_story_relation" {
			t.Errorf("unexpected path: %s, want /tcase_instance/delete_tcase_story_relation", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":true,"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.DeleteTCaseStoryRelationRequest{
		WorkspaceID: "10158231",
		StoryID:     "1010158231854812395",
		TCaseID:     "1010158231077224799",
		TestPlanID:  "1010158231077233617",
	}
	err := c.DeleteTCaseStoryRelation(context.Background(), req)
	if err != nil {
		t.Fatalf("DeleteTCaseStoryRelation() unexpected error: %v", err)
	}
}

func TestCreateTestPlan(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/test_plans" {
			t.Errorf("unexpected path: %s, want /test_plans", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"TestPlan":{"id":"1000000755000016443","workspace_id":"755","name":"test_plan_12","description":"这不是一个测试","status":"open","creator":"dev","created":"2020-01-09 21:11:52"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateTestPlanRequest{
		WorkspaceID: "755",
		Name:        "test_plan_12",
		Description: "这不是一个测试",
	}
	plan, err := c.CreateTestPlan(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateTestPlan() unexpected error: %v", err)
	}
	if plan.ID != "1000000755000016443" {
		t.Errorf("plan id = %q, want %q", plan.ID, "1000000755000016443")
	}
	if plan.Name != "test_plan_12" {
		t.Errorf("plan name = %q, want %q", plan.Name, "test_plan_12")
	}
	if plan.Status != "open" {
		t.Errorf("plan status = %q, want %q", plan.Status, "open")
	}
}

func TestUpdateTestPlan(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/test_plans" {
			t.Errorf("unexpected path: %s, want /test_plans", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"TestPlan":{"id":"1000000755000016443","workspace_id":"755","name":"test_plan_12","status":"open","modified":"2020-01-09 21:11:52"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.UpdateTestPlanRequest{
		WorkspaceID: "755",
		ID:          "1000000755000016443",
		Name:        "test_plan_12",
	}
	plan, err := c.UpdateTestPlan(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateTestPlan() unexpected error: %v", err)
	}
	if plan.ID != "1000000755000016443" {
		t.Errorf("plan id = %q, want %q", plan.ID, "1000000755000016443")
	}
	if plan.Name != "test_plan_12" {
		t.Errorf("plan name = %q, want %q", plan.Name, "test_plan_12")
	}
}

func TestListTestPlans(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans" {
			t.Errorf("unexpected path: %s, want /test_plans", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"TestPlan":{"id":"1010104801000045351","workspace_id":"10104801","name":"aada","status":"open","creator":"anyechen","created":"2020-01-09 12:12:37"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListTestPlansRequest{WorkspaceID: "10104801"}
	plans, err := c.ListTestPlans(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTestPlans() unexpected error: %v", err)
	}
	if len(plans) != 1 {
		t.Fatalf("expected 1 plan, got %d", len(plans))
	}
	if plans[0].ID != "1010104801000045351" {
		t.Errorf("plan id = %q, want %q", plans[0].ID, "1010104801000045351")
	}
	if plans[0].Name != "aada" {
		t.Errorf("plan name = %q, want %q", plans[0].Name, "aada")
	}
}

func TestCountTestPlans(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/count" {
			t.Errorf("unexpected path: %s, want /test_plans/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"count":4},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListTestPlansRequest{WorkspaceID: "10104801"}
	count, err := c.CountTestPlans(context.Background(), req)
	if err != nil {
		t.Fatalf("CountTestPlans() unexpected error: %v", err)
	}
	if count != 4 {
		t.Errorf("count = %d, want 4", count)
	}
}

func TestGetTestPlanDetails(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/details" {
			t.Errorf("unexpected path: %s, want /test_plans/details", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"Tcase":{"id":"1010158231075919347","workspace_id":"10158231","name":"Firefox浏览器兼容性测试","status":"normal"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanIDRequest{
		WorkspaceID: "10158231",
		ID:          "1010158231077233617",
	}
	details, err := c.GetTestPlanDetails(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTestPlanDetails() unexpected error: %v", err)
	}
	if len(details) != 1 {
		t.Fatalf("expected 1 tcase, got %d", len(details))
	}
	if details[0].ID != "1010158231075919347" {
		t.Errorf("tcase id = %q, want %q", details[0].ID, "1010158231075919347")
	}
	if details[0].Name != "Firefox浏览器兼容性测试" {
		t.Errorf("tcase name = %q, want %q", details[0].Name, "Firefox浏览器兼容性测试")
	}
}

func TestGetTestPlanFieldsInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/get_fields_info" {
			t.Errorf("unexpected path: %s, want /test_plans/get_fields_info", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"id":{"html_type":"input","options":[],"label":"ID"},"name":{"html_type":"input","options":[],"label":"计划名称"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetTestPlanFieldsInfo(context.Background(), &model.WorkspaceIDRequest{WorkspaceID: "10158231"})
	if err != nil {
		t.Fatalf("GetTestPlanFieldsInfo() unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(result))
	}
	if result["id"].Label != "ID" {
		t.Errorf("id label = %q, want %q", result["id"].Label, "ID")
	}
	if result["name"].Label != "计划名称" {
		t.Errorf("name label = %q, want %q", result["name"].Label, "计划名称")
	}
}

func TestGetTestPlanProgress(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/progress" {
			t.Errorf("unexpected path: %s, want /test_plans/progress", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"story_count":1,"tcase_count":10,"status_counter":{"pass":"5","no_pass":"0","block":"0","unexecuted":5},"executed_rate":"50%"},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanIDRequest{
		WorkspaceID: "10158231",
		ID:          "1010158231077233617",
	}
	progress, err := c.GetTestPlanProgress(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTestPlanProgress() unexpected error: %v", err)
	}
	if progress.StoryCount != 1 {
		t.Errorf("story_count = %d, want 1", progress.StoryCount)
	}
	if progress.TCaseCount != 10 {
		t.Errorf("tcase_count = %d, want 10", progress.TCaseCount)
	}
	if progress.ExecutedRate != "50%" {
		t.Errorf("executed_rate = %q, want %q", progress.ExecutedRate, "50%")
	}
}

func TestGetTestPlanRelativeStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/get_relative_stories" {
			t.Errorf("unexpected path: %s, want /test_plans/get_relative_stories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"story_ids":["1010104801500706241","1010104801854890913"]},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanRelativeStoriesRequest{
		WorkspaceID: "10104801",
		TestPlanID:  "1010104801000045351",
	}
	storyIDs, err := c.GetTestPlanRelativeStories(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTestPlanRelativeStories() unexpected error: %v", err)
	}
	if len(storyIDs) != 2 {
		t.Fatalf("expected 2 story ids, got %d", len(storyIDs))
	}
	if storyIDs[0] != "1010104801500706241" {
		t.Errorf("story_ids[0] = %q, want %q", storyIDs[0], "1010104801500706241")
	}
	if storyIDs[1] != "1010104801854890913" {
		t.Errorf("story_ids[1] = %q, want %q", storyIDs[1], "1010104801854890913")
	}
}

func TestListTestPlanTCaseRelations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/get_test_plan_tcase" {
			t.Errorf("unexpected path: %s, want /test_plans/get_test_plan_tcase", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"TestPlanStoryTcaseRelation":{"id":"1000000755002248699","workspace_id":"755","test_plan_id":"1000000755077233617","story_id":"0","tcase_id":"1000000755000026804","sort":"0","creator":"v_xuanfang"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanTCasesRequest{
		WorkspaceID: "755",
		TestPlanID:  "1000000755077233617",
	}
	relations, err := c.ListTestPlanTCaseRelations(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTestPlanTCaseRelations() unexpected error: %v", err)
	}
	if len(relations) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(relations))
	}
	if relations[0].ID != "1000000755002248699" {
		t.Errorf("relation id = %q, want %q", relations[0].ID, "1000000755002248699")
	}
	if relations[0].TcaseID != "1000000755000026804" {
		t.Errorf("tcase_id = %q, want %q", relations[0].TcaseID, "1000000755000026804")
	}
	if relations[0].Creator != "v_xuanfang" {
		t.Errorf("creator = %q, want %q", relations[0].Creator, "v_xuanfang")
	}
}

func TestGetTestPlanBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/result_relation_bugs" {
			t.Errorf("unexpected path: %s, want /test_plans/result_relation_bugs", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"id":1020357849077231365,"name":"用例2","tcase_result_relate_bugs":{}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanIDRequest{
		WorkspaceID: "20357849",
		ID:          "1020357849077233617",
	}
	bugs, err := c.GetTestPlanBugs(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTestPlanBugs() unexpected error: %v", err)
	}
	if len(bugs) != 1 {
		t.Fatalf("expected 1 bug item, got %d", len(bugs))
	}
	if bugs[0].Name != "用例2" {
		t.Errorf("bug name = %q, want %q", bugs[0].Name, "用例2")
	}
	if bugs[0].ID.String() != "1020357849077231365" {
		t.Errorf("bug id = %q, want %q", bugs[0].ID.String(), "1020357849077231365")
	}
}

func TestListTestPlansByIterationID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test_plans/get_by_iteration_id" {
			t.Errorf("unexpected path: %s, want /test_plans/get_by_iteration_id", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"workspace_id":"51650666","iteration_id":"1151650666001000111","test_plan_id":"1151650666001000019"}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlansByIterationIDRequest{
		WorkspaceID: "51650666",
		IterationID: "1151650666001000111",
	}
	plans, err := c.ListTestPlansByIterationID(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTestPlansByIterationID() unexpected error: %v", err)
	}
	if len(plans) != 1 {
		t.Fatalf("expected 1 plan, got %d", len(plans))
	}
	if plans[0].WorkspaceID != "51650666" {
		t.Errorf("workspace_id = %q, want %q", plans[0].WorkspaceID, "51650666")
	}
	if plans[0].IterationID != "1151650666001000111" {
		t.Errorf("iteration_id = %q, want %q", plans[0].IterationID, "1151650666001000111")
	}
	if plans[0].TestPlanID != "1151650666001000019" {
		t.Errorf("test_plan_id = %q, want %q", plans[0].TestPlanID, "1151650666001000019")
	}
}

func TestCreateTestPlanStoryRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/test_plans/create_story_relation" {
			t.Errorf("unexpected path: %s, want /test_plans/create_story_relation", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"create plan story relation success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanStoryRelationRequest{
		WorkspaceID: "755",
		PlanID:      "1000000755000016443",
		StoryIDs:    "1000000755854812395",
		Creator:     "dev",
	}
	err := c.CreateTestPlanStoryRelation(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateTestPlanStoryRelation() unexpected error: %v", err)
	}
}

func TestDeleteTestPlanStoryRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/test_plans/delete_story_relation" {
			t.Errorf("unexpected path: %s, want /test_plans/delete_story_relation", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"delete plan story relation success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanStoryRelationRequest{
		WorkspaceID: "755",
		PlanID:      "1000000755000016443",
		StoryIDs:    "1000000755854812395",
		Creator:     "dev",
	}
	err := c.DeleteTestPlanStoryRelation(context.Background(), req)
	if err != nil {
		t.Fatalf("DeleteTestPlanStoryRelation() unexpected error: %v", err)
	}
}

func TestCreateTestPlanTCaseRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/test_plans/create_tcase_relation" {
			t.Errorf("unexpected path: %s, want /test_plans/create_tcase_relation", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"create tcase relation success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.TestPlanTCaseRelationRequest{
		WorkspaceID: "755",
		TestPlanID:  "1000000755077233617",
		TCaseIDs:    "1000000755000026804",
		Creator:     "dev",
	}
	err := c.CreateTestPlanTCaseRelation(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateTestPlanTCaseRelation() unexpected error: %v", err)
	}
}
