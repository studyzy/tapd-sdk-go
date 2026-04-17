package tapd

import (
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateModule 创建模块
func (c *Client) CreateModule(req *model.CreateModuleRequest) (*model.Module, error) {
	data, err := c.doPost("/modules", req.ToParams())
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
func (c *Client) UpdateModule(req *model.UpdateModuleRequest) (*model.Module, error) {
	data, err := c.doPost("/modules", req.ToParams())
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
func (c *Client) GetModules(req *model.GetModulesRequest) ([]model.Module, error) {
	data, err := c.doGet("/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse module list: %w", err)
	}

	var modules []model.Module
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
func (c *Client) CountModules(req *model.CountModulesRequest) (int, error) {
	data, err := c.doGet("/modules/count", req.ToParams())
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
func (c *Client) CreateVersion(req *model.CreateVersionRequest) (*model.Version, error) {
	data, err := c.doPost("/versions", req.ToParams())
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
func (c *Client) UpdateVersion(req *model.UpdateVersionRequest) (*model.Version, error) {
	data, err := c.doPost("/versions", req.ToParams())
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
func (c *Client) GetVersions(req *model.GetVersionsRequest) ([]model.Version, error) {
	data, err := c.doGet("/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse version list: %w", err)
	}

	var versions []model.Version
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
func (c *Client) CountVersions(req *model.CountVersionsRequest) (int, error) {
	data, err := c.doGet("/versions/count", req.ToParams())
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
func (c *Client) CreateBaseline(req *model.CreateBaselineRequest) (*model.Baseline, error) {
	data, err := c.doPost("/baselines", req.ToParams())
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
func (c *Client) UpdateBaseline(req *model.UpdateBaselineRequest) (*model.Baseline, error) {
	data, err := c.doPost("/baselines", req.ToParams())
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
func (c *Client) GetBaselines(req *model.GetBaselinesRequest) ([]model.Baseline, error) {
	data, err := c.doGet("/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse baseline list: %w", err)
	}

	var baselines []model.Baseline
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
func (c *Client) CountBaselines(req *model.CountBaselinesRequest) (int, error) {
	data, err := c.doGet("/baselines/count", req.ToParams())
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
func (c *Client) CreateFeature(req *model.CreateFeatureRequest) (*model.Feature, error) {
	data, err := c.doPost("/features", req.ToParams())
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
func (c *Client) UpdateFeature(req *model.UpdateFeatureRequest) (*model.Feature, error) {
	data, err := c.doPost("/features", req.ToParams())
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
func (c *Client) GetFeatures(req *model.GetFeaturesRequest) ([]model.Feature, error) {
	data, err := c.doGet("/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse feature list: %w", err)
	}

	var features []model.Feature
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
func (c *Client) CountFeatures(req *model.CountFeaturesRequest) (int, error) {
	data, err := c.doGet("/features/count", req.ToParams())
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
