package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateModule 创建模块
func (c *Client) CreateModule(ctx context.Context, req *model.CreateModuleRequest) (*model.Module, error) {
	data, err := c.doPost(ctx, "/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Module](data, "Module")
}

// UpdateModule 更新模块
func (c *Client) UpdateModule(ctx context.Context, req *model.UpdateModuleRequest) (*model.Module, error) {
	data, err := c.doPost(ctx, "/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Module](data, "Module")
}

// GetModules 获取模块列表
func (c *Client) GetModules(ctx context.Context, req *model.GetModulesRequest) ([]model.Module, error) {
	data, err := c.doGet(ctx, "/modules", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Module](data, "Module")
}

// CountModules 获取模块数量
func (c *Client) CountModules(ctx context.Context, req *model.CountModulesRequest) (int, error) {
	data, err := c.doGet(ctx, "/modules/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// CreateVersion 创建版本
func (c *Client) CreateVersion(ctx context.Context, req *model.CreateVersionRequest) (*model.Version, error) {
	data, err := c.doPost(ctx, "/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Version](data, "Version")
}

// UpdateVersion 更新版本
func (c *Client) UpdateVersion(ctx context.Context, req *model.UpdateVersionRequest) (*model.Version, error) {
	data, err := c.doPost(ctx, "/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Version](data, "Version")
}

// GetVersions 获取版本列表
func (c *Client) GetVersions(ctx context.Context, req *model.GetVersionsRequest) ([]model.Version, error) {
	data, err := c.doGet(ctx, "/versions", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Version](data, "Version")
}

// CountVersions 获取版本数量
func (c *Client) CountVersions(ctx context.Context, req *model.CountVersionsRequest) (int, error) {
	data, err := c.doGet(ctx, "/versions/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// CreateBaseline 创建基线
func (c *Client) CreateBaseline(ctx context.Context, req *model.CreateBaselineRequest) (*model.Baseline, error) {
	data, err := c.doPost(ctx, "/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Baseline](data, "Baseline")
}

// UpdateBaseline 更新基线
func (c *Client) UpdateBaseline(ctx context.Context, req *model.UpdateBaselineRequest) (*model.Baseline, error) {
	data, err := c.doPost(ctx, "/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Baseline](data, "Baseline")
}

// GetBaselines 获取基线列表
func (c *Client) GetBaselines(ctx context.Context, req *model.GetBaselinesRequest) ([]model.Baseline, error) {
	data, err := c.doGet(ctx, "/baselines", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Baseline](data, "Baseline")
}

// CountBaselines 获取基线数量
func (c *Client) CountBaselines(ctx context.Context, req *model.CountBaselinesRequest) (int, error) {
	data, err := c.doGet(ctx, "/baselines/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// CreateFeature 创建特性
func (c *Client) CreateFeature(ctx context.Context, req *model.CreateFeatureRequest) (*model.Feature, error) {
	data, err := c.doPost(ctx, "/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Feature](data, "Feature")
}

// UpdateFeature 更新特性
func (c *Client) UpdateFeature(ctx context.Context, req *model.UpdateFeatureRequest) (*model.Feature, error) {
	data, err := c.doPost(ctx, "/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Feature](data, "Feature")
}

// GetFeatures 获取特性列表
func (c *Client) GetFeatures(ctx context.Context, req *model.GetFeaturesRequest) ([]model.Feature, error) {
	data, err := c.doGet(ctx, "/features", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Feature](data, "Feature")
}

// CountFeatures 获取特性数量
func (c *Client) CountFeatures(ctx context.Context, req *model.CountFeaturesRequest) (int, error) {
	data, err := c.doGet(ctx, "/features/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}
