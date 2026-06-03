package tapd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

// --- Module 模块测试 ---

func TestCreateModule(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/modules" {
			t.Errorf("unexpected path: %s, want /modules", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Module":{"id":"1001","workspace_id":"11111111","name":"核心模块","description":"核心功能模块"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	module, err := c.CreateModule(context.Background(), &model.CreateModuleRequest{
		WorkspaceID: "11111111",
		Name:        "核心模块",
		Description: "核心功能模块",
	})
	if err != nil {
		t.Fatalf("CreateModule() unexpected error: %v", err)
	}
	if module.ID != "1001" {
		t.Errorf("ID = %q, want %q", module.ID, "1001")
	}
	if module.Name != "核心模块" {
		t.Errorf("Name = %q, want %q", module.Name, "核心模块")
	}
	if module.Description != "核心功能模块" {
		t.Errorf("Description = %q, want %q", module.Description, "核心功能模块")
	}
}

func TestUpdateModule(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/modules" {
			t.Errorf("unexpected path: %s, want /modules", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Module":{"id":"1001","workspace_id":"11111111","name":"核心模块更新","description":"更新描述"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	module, err := c.UpdateModule(context.Background(), &model.UpdateModuleRequest{
		WorkspaceID: "11111111",
		ID:          "1001",
		Name:        "核心模块更新",
		Description: "更新描述",
	})
	if err != nil {
		t.Fatalf("UpdateModule() unexpected error: %v", err)
	}
	if module.ID != "1001" {
		t.Errorf("ID = %q, want %q", module.ID, "1001")
	}
	if module.Name != "核心模块更新" {
		t.Errorf("Name = %q, want %q", module.Name, "核心模块更新")
	}
}

func TestGetModules(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/modules" {
			t.Errorf("unexpected path: %s, want /modules", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"Module":{"id":"1001","workspace_id":"11111111","name":"核心模块","description":"核心功能模块"}},{"Module":{"id":"1002","workspace_id":"11111111","name":"辅助模块","description":"辅助功能模块"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	modules, err := c.GetModules(context.Background(), &model.GetModulesRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("GetModules() unexpected error: %v", err)
	}
	if len(modules) != 2 {
		t.Fatalf("expected 2 modules, got %d", len(modules))
	}
	if modules[0].ID != "1001" {
		t.Errorf("modules[0].ID = %q, want %q", modules[0].ID, "1001")
	}
	if modules[1].Name != "辅助模块" {
		t.Errorf("modules[1].Name = %q, want %q", modules[1].Name, "辅助模块")
	}
}

func TestCountModules(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/modules/count" {
			t.Errorf("unexpected path: %s, want /modules/count", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"count":3},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountModules(context.Background(), &model.CountModulesRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("CountModules() unexpected error: %v", err)
	}
	if count != 3 {
		t.Errorf("count = %d, want 3", count)
	}
}

// --- Version 版本测试 ---

func TestCreateVersion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/versions" {
			t.Errorf("unexpected path: %s, want /versions", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Version":{"id":"1001","workspace_id":"11111111","name":"v1.0","description":"第一版"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	version, err := c.CreateVersion(context.Background(), &model.CreateVersionRequest{
		WorkspaceID: "11111111",
		Name:        "v1.0",
		Description: "第一版",
	})
	if err != nil {
		t.Fatalf("CreateVersion() unexpected error: %v", err)
	}
	if version.ID != "1001" {
		t.Errorf("ID = %q, want %q", version.ID, "1001")
	}
	if version.Name != "v1.0" {
		t.Errorf("Name = %q, want %q", version.Name, "v1.0")
	}
}

func TestUpdateVersion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/versions" {
			t.Errorf("unexpected path: %s, want /versions", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Version":{"id":"1001","workspace_id":"11111111","name":"v1.1","description":"更新版本"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	version, err := c.UpdateVersion(context.Background(), &model.UpdateVersionRequest{
		WorkspaceID: "11111111",
		ID:          "1001",
		Name:        "v1.1",
		Description: "更新版本",
	})
	if err != nil {
		t.Fatalf("UpdateVersion() unexpected error: %v", err)
	}
	if version.ID != "1001" {
		t.Errorf("ID = %q, want %q", version.ID, "1001")
	}
	if version.Name != "v1.1" {
		t.Errorf("Name = %q, want %q", version.Name, "v1.1")
	}
}

func TestGetVersions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/versions" {
			t.Errorf("unexpected path: %s, want /versions", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"Version":{"id":"1001","workspace_id":"11111111","name":"v1.0","description":"第一版"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	versions, err := c.GetVersions(context.Background(), &model.GetVersionsRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("GetVersions() unexpected error: %v", err)
	}
	if len(versions) != 1 {
		t.Fatalf("expected 1 version, got %d", len(versions))
	}
	if versions[0].ID != "1001" {
		t.Errorf("versions[0].ID = %q, want %q", versions[0].ID, "1001")
	}
}

func TestCountVersions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/versions/count" {
			t.Errorf("unexpected path: %s, want /versions/count", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"count":5},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountVersions(context.Background(), &model.CountVersionsRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("CountVersions() unexpected error: %v", err)
	}
	if count != 5 {
		t.Errorf("count = %d, want 5", count)
	}
}

// --- Baseline 基线测试 ---

func TestCreateBaseline(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/baselines" {
			t.Errorf("unexpected path: %s, want /baselines", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Baseline":{"id":"2001","workspace_id":"11111111","name":"基线1","description":"第一条基线"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	baseline, err := c.CreateBaseline(context.Background(), &model.CreateBaselineRequest{
		WorkspaceID: "11111111",
		Name:        "基线1",
		Description: "第一条基线",
	})
	if err != nil {
		t.Fatalf("CreateBaseline() unexpected error: %v", err)
	}
	if baseline.ID != "2001" {
		t.Errorf("ID = %q, want %q", baseline.ID, "2001")
	}
	if baseline.Name != "基线1" {
		t.Errorf("Name = %q, want %q", baseline.Name, "基线1")
	}
}

func TestUpdateBaseline(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/baselines" {
			t.Errorf("unexpected path: %s, want /baselines", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Baseline":{"id":"2001","workspace_id":"11111111","name":"基线1更新","description":"更新描述"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	baseline, err := c.UpdateBaseline(context.Background(), &model.UpdateBaselineRequest{
		WorkspaceID: "11111111",
		ID:          "2001",
		Name:        "基线1更新",
		Description: "更新描述",
	})
	if err != nil {
		t.Fatalf("UpdateBaseline() unexpected error: %v", err)
	}
	if baseline.ID != "2001" {
		t.Errorf("ID = %q, want %q", baseline.ID, "2001")
	}
	if baseline.Name != "基线1更新" {
		t.Errorf("Name = %q, want %q", baseline.Name, "基线1更新")
	}
}

func TestGetBaselines(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/baselines" {
			t.Errorf("unexpected path: %s, want /baselines", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"Baseline":{"id":"2001","workspace_id":"11111111","name":"基线1","description":"第一条基线"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	baselines, err := c.GetBaselines(context.Background(), &model.GetBaselinesRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("GetBaselines() unexpected error: %v", err)
	}
	if len(baselines) != 1 {
		t.Fatalf("expected 1 baseline, got %d", len(baselines))
	}
	if baselines[0].ID != "2001" {
		t.Errorf("baselines[0].ID = %q, want %q", baselines[0].ID, "2001")
	}
}

func TestCountBaselines(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/baselines/count" {
			t.Errorf("unexpected path: %s, want /baselines/count", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"count":2},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountBaselines(context.Background(), &model.CountBaselinesRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("CountBaselines() unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}

// --- Feature 特性测试 ---

func TestCreateFeature(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/features" {
			t.Errorf("unexpected path: %s, want /features", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"Feature":{"id":"3001","workspace_id":"11111111","name":"特性1","description":"第一个特性"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	feature, err := c.CreateFeature(context.Background(), &model.CreateFeatureRequest{
		WorkspaceID: "11111111",
		Name:        "特性1",
		Description: "第一个特性",
	})
	if err != nil {
		t.Fatalf("CreateFeature() unexpected error: %v", err)
	}
	if feature.ID != "3001" {
		t.Errorf("ID = %q, want %q", feature.ID, "3001")
	}
	if feature.Name != "特性1" {
		t.Errorf("Name = %q, want %q", feature.Name, "特性1")
	}
}

func TestUpdateFeature(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/features" {
			t.Errorf("unexpected path: %s, want /features", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"Feature":{"id":"3001","workspace_id":"11111111","name":"特性1更新","description":"更新描述"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	feature, err := c.UpdateFeature(context.Background(), &model.UpdateFeatureRequest{
		WorkspaceID: "11111111",
		ID:          "3001",
		Name:        "特性1更新",
		Description: "更新描述",
	})
	if err != nil {
		t.Fatalf("UpdateFeature() unexpected error: %v", err)
	}
	if feature.ID != "3001" {
		t.Errorf("ID = %q, want %q", feature.ID, "3001")
	}
	if feature.Name != "特性1更新" {
		t.Errorf("Name = %q, want %q", feature.Name, "特性1更新")
	}
}

func TestGetFeatures(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/features" {
			t.Errorf("unexpected path: %s, want /features", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":[{"Feature":{"id":"3001","workspace_id":"11111111","name":"特性1","description":"第一个特性"}},{"Feature":{"id":"3002","workspace_id":"11111111","name":"特性2","description":"第二个特性"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	features, err := c.GetFeatures(context.Background(), &model.GetFeaturesRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("GetFeatures() unexpected error: %v", err)
	}
	if len(features) != 2 {
		t.Fatalf("expected 2 features, got %d", len(features))
	}
	if features[0].ID != "3001" {
		t.Errorf("features[0].ID = %q, want %q", features[0].ID, "3001")
	}
	if features[1].Name != "特性2" {
		t.Errorf("features[1].Name = %q, want %q", features[1].Name, "特性2")
	}
}

func TestCountFeatures(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/features/count" {
			t.Errorf("unexpected path: %s, want /features/count", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"count":7},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountFeatures(context.Background(), &model.CountFeaturesRequest{
		WorkspaceID: "11111111",
	})
	if err != nil {
		t.Fatalf("CountFeatures() unexpected error: %v", err)
	}
	if count != 7 {
		t.Errorf("count = %d, want 7", count)
	}
}

// --- Setting 配置测试 ---

func TestAddCustomFieldConfig(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/custom_field_configs" {
			t.Errorf("unexpected path: %s, want /custom_field_configs", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"CustomFieldConfig":{"id":"1010158231215016293","workspace_id":"10158231","entry_type":"story","custom_field":"custom_field_three","type":"text","name":"add_field","options":"","enabled":"1","sort":"","memo":""}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	config, err := c.AddCustomFieldConfig(context.Background(), &model.AddCustomFieldConfigRequest{
		WorkspaceID: "10158231",
		EntryType:   "story",
		Name:        "add_field",
		Type:        "text",
	})
	if err != nil {
		t.Fatalf("AddCustomFieldConfig() unexpected error: %v", err)
	}
	if config.ID != "1010158231215016293" {
		t.Errorf("ID = %q, want %q", config.ID, "1010158231215016293")
	}
	if config.EntryType != "story" {
		t.Errorf("EntryType = %q, want %q", config.EntryType, "story")
	}
	if config.CustomField != "custom_field_three" {
		t.Errorf("CustomField = %q, want %q", config.CustomField, "custom_field_three")
	}
	if config.Name != "add_field" {
		t.Errorf("Name = %q, want %q", config.Name, "add_field")
	}
	if config.Enabled != "1" {
		t.Errorf("Enabled = %q, want %q", config.Enabled, "1")
	}
}

func TestUpdateBugSelectFieldOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/custom_field_configs/update_bug_select_field_options" {
			t.Errorf("unexpected path: %s, want /custom_field_configs/update_bug_select_field_options", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"status":1},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.UpdateBugSelectFieldOptions(context.Background(), &model.UpdateSelectFieldOptionsRequest{
		WorkspaceID: "10104801",
		ID:          "1010104801214991279",
		Options:     "开发|测试|产品|运营",
	})
	if err != nil {
		t.Fatalf("UpdateBugSelectFieldOptions() unexpected error: %v", err)
	}
}

func TestUpdateStorySelectFieldOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/custom_field_configs/update_story_select_field_options" {
			t.Errorf("unexpected path: %s, want /custom_field_configs/update_story_select_field_options", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"status":1},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.UpdateStorySelectFieldOptions(context.Background(), &model.UpdateSelectFieldOptionsRequest{
		WorkspaceID: "10104801",
		ID:          "1010104801214991280",
		Options:     "高|中|低",
	})
	if err != nil {
		t.Fatalf("UpdateStorySelectFieldOptions() unexpected error: %v", err)
	}
}

func TestUpdateCascadeFieldOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/custom_field_configs/update_cascade_field_options" {
			t.Errorf("unexpected path: %s, want /custom_field_configs/update_cascade_field_options", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"status":1},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.UpdateCascadeFieldOptions(context.Background(), &model.UpdateCascadeFieldOptionsRequest{
		WorkspaceID: "10104801",
		ID:          "1010104801214991281",
		Options:     `[{"name":"1","children":[{"name":"11"}]},{"name":"2","children":[{"name":"21"}]}]`,
	})
	if err != nil {
		t.Fatalf("UpdateCascadeFieldOptions() unexpected error: %v", err)
	}
}

func TestGetWorkspaceSetting(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/settings/get_workspace_setting" {
			t.Errorf("unexpected path: %s, want /settings/get_workspace_setting", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		if r.URL.Query().Get("type") != "workspace_metrology" {
			t.Errorf("type = %q, want %q", r.URL.Query().Get("type"), "workspace_metrology")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"workspace_metrology":"day"},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetWorkspaceSetting(context.Background(), &model.GetWorkspaceSettingRequest{
		WorkspaceID: "10104801",
		Type:        "workspace_metrology",
	})
	if err != nil {
		t.Fatalf("GetWorkspaceSetting() unexpected error: %v", err)
	}
	if result["workspace_metrology"] != "day" {
		t.Errorf("workspace_metrology = %q, want %q", result["workspace_metrology"], "day")
	}
}

func TestUpdateSelectFieldOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/custom_field_configs/update_select_field_options" {
			t.Errorf("unexpected path: %s, want /custom_field_configs/update_select_field_options", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"status":1},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.UpdateSelectFieldOptions(context.Background(), &model.UpdateSelectFieldOptionsUnifiedRequest{
		WorkspaceID: "10104801",
		ID:          "1010104801214991282",
		Options:     "选项A|选项B|选项C",
	})
	if err != nil {
		t.Fatalf("UpdateSelectFieldOptions() unexpected error: %v", err)
	}
}

func TestCopyWorkitemTypeSetting(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/stories/copy_workitem_type_setting" {
			t.Errorf("unexpected path: %s, want /stories/copy_workitem_type_setting", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"data":{"WorkitemType":{"id":"1010104801000061611","workspace_id":"10104801","entity_type":"story","name":"技术需求","english_name":"TSTORY","status":"3","color":"#5c88c5","workflow_id":"1010104801001087491","creator":"v_xuanfang","created":"2020-12-02 15:47:02","modified_by":"v_xuanfang","modified":"2021-01-21 10:49:57"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	wt, err := c.CopyWorkitemTypeSetting(context.Background(), &model.CopyWorkitemTypeSettingRequest{
		SrcWorkspaceID:    "755",
		SrcWorkitemTypeID: "1020358527000067037",
		WorkspaceID:       "10104801",
	})
	if err != nil {
		t.Fatalf("CopyWorkitemTypeSetting() unexpected error: %v", err)
	}
	if wt.ID != "1010104801000061611" {
		t.Errorf("ID = %q, want %q", wt.ID, "1010104801000061611")
	}
	if wt.Name != "技术需求" {
		t.Errorf("Name = %q, want %q", wt.Name, "技术需求")
	}
	if wt.EntityType != "story" {
		t.Errorf("EntityType = %q, want %q", wt.EntityType, "story")
	}
	if wt.WorkspaceID != "10104801" {
		t.Errorf("WorkspaceID = %q, want %q", wt.WorkspaceID, "10104801")
	}
	if wt.Color != "#5c88c5" {
		t.Errorf("Color = %q, want %q", wt.Color, "#5c88c5")
	}
}

func TestCopyBugSetting(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/copy_settings" {
			t.Errorf("unexpected path: %s, want /bugs/copy_settings", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":1,"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.CopyBugSetting(context.Background(), &model.CopyBugSettingRequest{
		SrcWorkspaceID:    "10104801",
		TargetWorkspaceID: "755",
	})
	if err != nil {
		t.Fatalf("CopyBugSetting() unexpected error: %v", err)
	}
}
