package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateModule 创建模块
func (c *Client) CreateModule(ctx context.Context, req *model.CreateModuleRequest) (*model.Module, error) {
	data, err := c.doPost(ctx, "/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create module response: %w", err)
	}

	raw, ok := wrapper["Module"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var module model.Module
	if err := json.Unmarshal(raw, &module); err != nil {
		return nil, fmt.Errorf("failed to parse created module: %w", err)
	}

	return &module, nil
}

// UpdateModule 更新模块
func (c *Client) UpdateModule(ctx context.Context, req *model.UpdateModuleRequest) (*model.Module, error) {
	data, err := c.doPost(ctx, "/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update module response: %w", err)
	}

	raw, ok := wrapper["Module"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var module model.Module
	if err := json.Unmarshal(raw, &module); err != nil {
		return nil, fmt.Errorf("failed to parse updated module: %w", err)
	}

	return &module, nil
}

// GetModules 获取模块列表
func (c *Client) GetModules(ctx context.Context, req *model.GetModulesRequest) ([]model.Module, error) {
	data, err := c.doGet(ctx, "/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse module list: %w", err)
	}

	modules := make([]model.Module, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["Module"]; ok {
			var m model.Module
			if err := json.Unmarshal(raw, &m); err == nil {
				modules = append(modules, m)
			}
		}
	}
	return modules, nil
}

// CountModules 获取模块数量
func (c *Client) CountModules(ctx context.Context, req *model.CountModulesRequest) (int, error) {
	data, err := c.doGet(ctx, "/modules/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}

// CreateVersion 创建版本
func (c *Client) CreateVersion(ctx context.Context, req *model.CreateVersionRequest) (*model.Version, error) {
	data, err := c.doPost(ctx, "/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create version response: %w", err)
	}

	raw, ok := wrapper["Version"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var version model.Version
	if err := json.Unmarshal(raw, &version); err != nil {
		return nil, fmt.Errorf("failed to parse created version: %w", err)
	}

	return &version, nil
}

// UpdateVersion 更新版本
func (c *Client) UpdateVersion(ctx context.Context, req *model.UpdateVersionRequest) (*model.Version, error) {
	data, err := c.doPost(ctx, "/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update version response: %w", err)
	}

	raw, ok := wrapper["Version"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var version model.Version
	if err := json.Unmarshal(raw, &version); err != nil {
		return nil, fmt.Errorf("failed to parse updated version: %w", err)
	}

	return &version, nil
}

// GetVersions 获取版本列表
func (c *Client) GetVersions(ctx context.Context, req *model.GetVersionsRequest) ([]model.Version, error) {
	data, err := c.doGet(ctx, "/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse version list: %w", err)
	}

	versions := make([]model.Version, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["Version"]; ok {
			var v model.Version
			if err := json.Unmarshal(raw, &v); err == nil {
				versions = append(versions, v)
			}
		}
	}
	return versions, nil
}

// CountVersions 获取版本数量
func (c *Client) CountVersions(ctx context.Context, req *model.CountVersionsRequest) (int, error) {
	data, err := c.doGet(ctx, "/versions/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}

// CreateBaseline 创建基线
func (c *Client) CreateBaseline(ctx context.Context, req *model.CreateBaselineRequest) (*model.Baseline, error) {
	data, err := c.doPost(ctx, "/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create baseline response: %w", err)
	}

	raw, ok := wrapper["Baseline"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var baseline model.Baseline
	if err := json.Unmarshal(raw, &baseline); err != nil {
		return nil, fmt.Errorf("failed to parse created baseline: %w", err)
	}

	return &baseline, nil
}

// UpdateBaseline 更新基线
func (c *Client) UpdateBaseline(ctx context.Context, req *model.UpdateBaselineRequest) (*model.Baseline, error) {
	data, err := c.doPost(ctx, "/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update baseline response: %w", err)
	}

	raw, ok := wrapper["Baseline"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var baseline model.Baseline
	if err := json.Unmarshal(raw, &baseline); err != nil {
		return nil, fmt.Errorf("failed to parse updated baseline: %w", err)
	}

	return &baseline, nil
}

// GetBaselines 获取基线列表
func (c *Client) GetBaselines(ctx context.Context, req *model.GetBaselinesRequest) ([]model.Baseline, error) {
	data, err := c.doGet(ctx, "/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse baseline list: %w", err)
	}

	baselines := make([]model.Baseline, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["Baseline"]; ok {
			var b model.Baseline
			if err := json.Unmarshal(raw, &b); err == nil {
				baselines = append(baselines, b)
			}
		}
	}
	return baselines, nil
}

// CountBaselines 获取基线数量
func (c *Client) CountBaselines(ctx context.Context, req *model.CountBaselinesRequest) (int, error) {
	data, err := c.doGet(ctx, "/baselines/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}

// CreateFeature 创建特性
func (c *Client) CreateFeature(ctx context.Context, req *model.CreateFeatureRequest) (*model.Feature, error) {
	data, err := c.doPost(ctx, "/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create feature response: %w", err)
	}

	raw, ok := wrapper["Feature"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var feature model.Feature
	if err := json.Unmarshal(raw, &feature); err != nil {
		return nil, fmt.Errorf("failed to parse created feature: %w", err)
	}

	return &feature, nil
}

// UpdateFeature 更新特性
func (c *Client) UpdateFeature(ctx context.Context, req *model.UpdateFeatureRequest) (*model.Feature, error) {
	data, err := c.doPost(ctx, "/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update feature response: %w", err)
	}

	raw, ok := wrapper["Feature"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var feature model.Feature
	if err := json.Unmarshal(raw, &feature); err != nil {
		return nil, fmt.Errorf("failed to parse updated feature: %w", err)
	}

	return &feature, nil
}

// GetFeatures 获取特性列表
func (c *Client) GetFeatures(ctx context.Context, req *model.GetFeaturesRequest) ([]model.Feature, error) {
	data, err := c.doGet(ctx, "/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse feature list: %w", err)
	}

	features := make([]model.Feature, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["Feature"]; ok {
			var f model.Feature
			if err := json.Unmarshal(raw, &f); err == nil {
				features = append(features, f)
			}
		}
	}
	return features, nil
}

// CountFeatures 获取特性数量
func (c *Client) CountFeatures(ctx context.Context, req *model.CountFeaturesRequest) (int, error) {
	data, err := c.doGet(ctx, "/features/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}
