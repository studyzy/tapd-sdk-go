package tapd

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestCreateRelease(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/new_releases" {
			t.Errorf("unexpected path: %s, want /new_releases", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		bodyStr := string(body)
		if bodyStr == "" {
			t.Fatal("expected non-empty body")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Release":{"id":"1111111111001000001","workspace_id":"11111111","name":"v1.0","description":"first release","startdate":"2026-01-01","enddate":"2026-02-01","status":"open","created":"2026-01-01 10:00:00","modified":"2026-01-01 10:00:00"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	release, err := c.CreateRelease(context.Background(), &model.CreateReleaseRequest{
		WorkspaceID: "11111111",
		Name:        "v1.0",
		Description: "first release",
		StartDate:   "2026-01-01",
		EndDate:     "2026-02-01",
	})
	if err != nil {
		t.Fatalf("CreateRelease() unexpected error: %v", err)
	}
	if release.ID != "1111111111001000001" {
		t.Errorf("release id = %q, want %q", release.ID, "1111111111001000001")
	}
	if release.Name != "v1.0" {
		t.Errorf("release name = %q, want %q", release.Name, "v1.0")
	}
	if release.Description != "first release" {
		t.Errorf("release description = %q, want %q", release.Description, "first release")
	}
	if release.Status != "open" {
		t.Errorf("release status = %q, want %q", release.Status, "open")
	}
}

func TestUpdateRelease(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/new_releases" {
			t.Errorf("unexpected path: %s, want /new_releases", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Release":{"id":"1111111111001000001","workspace_id":"11111111","name":"v1.1","status":"done"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	release, err := c.UpdateRelease(context.Background(), &model.UpdateReleaseRequest{
		WorkspaceID: "11111111",
		ID:          "1111111111001000001",
		Name:        "v1.1",
		Status:      "done",
	})
	if err != nil {
		t.Fatalf("UpdateRelease() unexpected error: %v", err)
	}
	if release.ID != "1111111111001000001" {
		t.Errorf("release id = %q, want %q", release.ID, "1111111111001000001")
	}
	if release.Name != "v1.1" {
		t.Errorf("release name = %q, want %q", release.Name, "v1.1")
	}
	if release.Status != "done" {
		t.Errorf("release status = %q, want %q", release.Status, "done")
	}
}

func TestCountReleases(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/new_releases/count" {
			t.Errorf("unexpected path: %s, want /new_releases/count", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "11111111" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "11111111")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":5},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountReleases(context.Background(), &model.CountReleasesRequest{WorkspaceID: "11111111"})
	if err != nil {
		t.Fatalf("CountReleases() unexpected error: %v", err)
	}
	if count != 5 {
		t.Errorf("count = %d, want %d", count, 5)
	}
}

func TestGetLaunchForms(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms" {
			t.Errorf("unexpected path: %s, want /launch_forms", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"LaunchForm":{"id":"1010104801079697767","title":"","name":"202101150008","creator":"v_xuanfang","created":"2021-01-15 19:47:37","workspace_id":"10104801","status":"initial","release_type":"正常发布","participator":";v_xuanfang;","template_id":"1010104801065798331","flows":"initial","custom_field_one":"test_value"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	forms, err := c.GetLaunchForms(context.Background(), &model.GetLaunchFormsRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("GetLaunchForms() unexpected error: %v", err)
	}
	if len(forms) != 1 {
		t.Fatalf("expected 1 form, got %d", len(forms))
	}
	if forms[0].ID != "1010104801079697767" {
		t.Errorf("form id = %q, want %q", forms[0].ID, "1010104801079697767")
	}
	if forms[0].Name != "202101150008" {
		t.Errorf("form name = %q, want %q", forms[0].Name, "202101150008")
	}
	if forms[0].Status != "initial" {
		t.Errorf("form status = %q, want %q", forms[0].Status, "initial")
	}
	if forms[0].CustomFields["custom_field_one"] != "test_value" {
		t.Errorf("custom_field_one = %q, want %q", forms[0].CustomFields["custom_field_one"], "test_value")
	}
}

func TestUpdateLaunchForm(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms" {
			t.Errorf("unexpected path: %s, want /launch_forms", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) == "" {
			t.Fatal("expected non-empty body")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"LaunchForm":{"id":"1010104801079724013","name":"202107210009","creator":"v_xuanfang","created":"2021-07-21 14:27:36","workspace_id":"10104801","status":"initial","release_type":"正常发布","template_id":"1010104801065798331","flows":"initial"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	form, err := c.UpdateLaunchForm(context.Background(), &model.UpdateLaunchFormRequest{
		WorkspaceID: "10104801",
		ID:          "1010104801079724013",
		Status:      "initial",
	})
	if err != nil {
		t.Fatalf("UpdateLaunchForm() unexpected error: %v", err)
	}
	if form.ID != "1010104801079724013" {
		t.Errorf("form id = %q, want %q", form.ID, "1010104801079724013")
	}
	if form.Status != "initial" {
		t.Errorf("form status = %q, want %q", form.Status, "initial")
	}
	if form.WorkspaceID != "10104801" {
		t.Errorf("form workspace_id = %q, want %q", form.WorkspaceID, "10104801")
	}
}

func TestCountLaunchForms(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms/count" {
			t.Errorf("unexpected path: %s, want /launch_forms/count", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "20003271" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "20003271")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":1},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountLaunchForms(context.Background(), &model.CountLaunchFormsRequest{WorkspaceID: "20003271"})
	if err != nil {
		t.Fatalf("CountLaunchForms() unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("count = %d, want %d", count, 1)
	}
}

func TestCreateLaunchForm(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms" {
			t.Errorf("unexpected path: %s, want /launch_forms", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) == "" {
			t.Fatal("expected non-empty body")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"LaunchForm":{"id":"1010104801079724013","name":"202107210009","creator":"v_xuanfang","created":"2021-07-21 14:27:36","workspace_id":"10104801","status":"initial","release_type":"正常发布","template_id":"1010104801065798331","flows":"initial"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	form, err := c.CreateLaunchForm(context.Background(), &model.CreateLaunchFormRequest{
		WorkspaceID: "10104801",
		Creator:     "v_xuanfang",
		TemplateID:  "1010104801065798331",
	})
	if err != nil {
		t.Fatalf("CreateLaunchForm() unexpected error: %v", err)
	}
	if form.ID != "1010104801079724013" {
		t.Errorf("form id = %q, want %q", form.ID, "1010104801079724013")
	}
	if form.Creator != "v_xuanfang" {
		t.Errorf("form creator = %q, want %q", form.Creator, "v_xuanfang")
	}
	if form.Status != "initial" {
		t.Errorf("form status = %q, want %q", form.Status, "initial")
	}
}

func TestGetLaunchFormsTemplates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms/templates" {
			t.Errorf("unexpected path: %s, want /launch_forms/templates", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "20042301" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "20042301")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"template":{"id":"1120042301001000009","name":"系统默认流程"}},{"template":{"id":"1120042301001000076","name":"广告歌"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	templates, err := c.GetLaunchFormsTemplates(context.Background(), &model.GetLaunchFormsTemplatesRequest{WorkspaceID: "20042301"})
	if err != nil {
		t.Fatalf("GetLaunchFormsTemplates() unexpected error: %v", err)
	}
	if len(templates) != 2 {
		t.Fatalf("expected 2 templates, got %d", len(templates))
	}
	if templates[0].ID != "1120042301001000009" {
		t.Errorf("template[0] id = %q, want %q", templates[0].ID, "1120042301001000009")
	}
	if templates[0].Name != "系统默认流程" {
		t.Errorf("template[0] name = %q, want %q", templates[0].Name, "系统默认流程")
	}
	if templates[1].ID != "1120042301001000076" {
		t.Errorf("template[1] id = %q, want %q", templates[1].ID, "1120042301001000076")
	}
}

func TestGetLaunchFormsActivityLogs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms/get_activity_logs" {
			t.Errorf("unexpected path: %s, want /launch_forms/get_activity_logs", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		if r.URL.Query().Get("form_id") != "1010104801079777231" {
			t.Errorf("form_id = %q, want %q", r.URL.Query().Get("form_id"), "1010104801079777231")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"LaunchChange":{"id":"1010104801083909701","workspace_id":"10104801","type":"signing","form_id":"1010104801079777231","activity_form_id":"","field":"sign","old_value":"signing","new_value":"sign_agree","created_by":"v_xuanfang","created":"2023-07-13 15:13:34","operation":"sign_agree"}},{"LaunchChange":{"id":"1010104801083448609","workspace_id":"10104801","type":"initial","form_id":"1010104801079777231","activity_form_id":"","field":"initialization","old_value":"","new_value":"","created_by":"v_xuanfang","created":"2022-09-08 17:27:34","operation":"initialization"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	logs, err := c.GetLaunchFormsActivityLogs(context.Background(), &model.GetLaunchFormsActivityLogsRequest{
		WorkspaceID: "10104801",
		FormID:      "1010104801079777231",
	})
	if err != nil {
		t.Fatalf("GetLaunchFormsActivityLogs() unexpected error: %v", err)
	}
	if len(logs) != 2 {
		t.Fatalf("expected 2 logs, got %d", len(logs))
	}
	if logs[0].ID != "1010104801083909701" {
		t.Errorf("log[0] id = %q, want %q", logs[0].ID, "1010104801083909701")
	}
	if logs[0].Type != "signing" {
		t.Errorf("log[0] type = %q, want %q", logs[0].Type, "signing")
	}
	if logs[0].Operation != "sign_agree" {
		t.Errorf("log[0] operation = %q, want %q", logs[0].Operation, "sign_agree")
	}
	if logs[1].Type != "initial" {
		t.Errorf("log[1] type = %q, want %q", logs[1].Type, "initial")
	}
}

func TestCreateLaunchAccessory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_accessories" {
			t.Errorf("unexpected path: %s, want /launch_accessories", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) == "" {
			t.Fatal("expected non-empty body")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"LaunchAccessory":{"id":"1010104801000254035","form_id":"1010104801079533889","workspace_id":"10104801","type":"launch_url","tag":"","title":"URL","content":"https://www.tapd.cn/","description":"","content_type":"","created_by":"tapd","created":"2022-09-08 16:45:30","group_id":"","source":""}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	accessory, err := c.CreateLaunchAccessory(context.Background(), &model.CreateLaunchAccessoryRequest{
		WorkspaceID: "10104801",
		FormID:      "1010104801079533889",
		Type:        "launch_url",
		Content:     "https://www.tapd.cn/",
	})
	if err != nil {
		t.Fatalf("CreateLaunchAccessory() unexpected error: %v", err)
	}
	if accessory.ID != "1010104801000254035" {
		t.Errorf("accessory id = %q, want %q", accessory.ID, "1010104801000254035")
	}
	if accessory.Type != "launch_url" {
		t.Errorf("accessory type = %q, want %q", accessory.Type, "launch_url")
	}
	if accessory.Content != "https://www.tapd.cn/" {
		t.Errorf("accessory content = %q, want %q", accessory.Content, "https://www.tapd.cn/")
	}
	if accessory.CreatedBy != "tapd" {
		t.Errorf("accessory created_by = %q, want %q", accessory.CreatedBy, "tapd")
	}
}

func TestGetLaunchAccessories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_accessories" {
			t.Errorf("unexpected path: %s, want /launch_accessories", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		if r.URL.Query().Get("form_id") != "1010104801000402051" {
			t.Errorf("form_id = %q, want %q", r.URL.Query().Get("form_id"), "1010104801000402051")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"LaunchAccessory":{"id":"1010104801000253485","form_id":"1010104801000402051","workspace_id":"10104801","type":"launch_tasks_list","tag":"","title":"任务列表","content":"1010104801500601739","description":"","content_type":"task","created_by":"v_xuanfang","created":"2020-06-11 16:17:56","group_id":"","source":""}},{"LaunchAccessory":{"id":"1010104801000253477","form_id":"1010104801000402051","workspace_id":"10104801","type":"launch_url","tag":"","title":"URL","content":"baidu.com","description":"","content_type":"","created_by":"v_xuanfang","created":"2020-06-10 19:30:56","group_id":"","source":""}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	accessories, err := c.GetLaunchAccessories(context.Background(), &model.GetLaunchAccessoriesRequest{
		WorkspaceID: "10104801",
		FormID:      "1010104801000402051",
	})
	if err != nil {
		t.Fatalf("GetLaunchAccessories() unexpected error: %v", err)
	}
	if len(accessories) != 2 {
		t.Fatalf("expected 2 accessories, got %d", len(accessories))
	}
	if accessories[0].ID != "1010104801000253485" {
		t.Errorf("accessory[0] id = %q, want %q", accessories[0].ID, "1010104801000253485")
	}
	if accessories[0].Type != "launch_tasks_list" {
		t.Errorf("accessory[0] type = %q, want %q", accessories[0].Type, "launch_tasks_list")
	}
	if accessories[1].Content != "baidu.com" {
		t.Errorf("accessory[1] content = %q, want %q", accessories[1].Content, "baidu.com")
	}
}

func TestGetLaunchFormsCustomFieldsSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/launch_forms/custom_fields_settings" {
			t.Errorf("unexpected path: %s, want /launch_forms/custom_fields_settings", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "20003271" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "20003271")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"CustomFieldConfig":{"id":"1120003271001000004","workspace_id":"20003271","entry_type":"launchform","custom_field":"custom_field_one","type":"textarea","name":"DB变更","options":"","enabled":"1","sort":""}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	configs, err := c.GetLaunchFormsCustomFieldsSettings(context.Background(), &model.WorkspaceIDRequest{WorkspaceID: "20003271"})
	if err != nil {
		t.Fatalf("GetLaunchFormsCustomFieldsSettings() unexpected error: %v", err)
	}
	if len(configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(configs))
	}
	if configs[0].ID != "1120003271001000004" {
		t.Errorf("config id = %q, want %q", configs[0].ID, "1120003271001000004")
	}
	if configs[0].EntryType != "launchform" {
		t.Errorf("config entry_type = %q, want %q", configs[0].EntryType, "launchform")
	}
	if configs[0].CustomField != "custom_field_one" {
		t.Errorf("config custom_field = %q, want %q", configs[0].CustomField, "custom_field_one")
	}
	if configs[0].Name != "DB变更" {
		t.Errorf("config name = %q, want %q", configs[0].Name, "DB变更")
	}
	if configs[0].Enabled != "1" {
		t.Errorf("config enabled = %q, want %q", configs[0].Enabled, "1")
	}
}
