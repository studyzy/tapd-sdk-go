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
		fmt.Fprint(w, `{"status":1,"data":{"Feature":{"id":"3001","workspace_id":"11111111","name":"特性1","description":"第一个特性"}},"info":"success"}`)
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
