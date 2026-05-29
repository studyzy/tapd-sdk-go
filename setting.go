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

// AddCustomFieldConfig 创建自定义字段（需求/缺陷/任务/测试用例）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/add_custom_field_config.html
func (c *Client) AddCustomFieldConfig(ctx context.Context, req *model.AddCustomFieldConfigRequest) (*model.CustomFieldConfig, error) {
	data, err := c.doPost(ctx, "/custom_field_configs", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.CustomFieldConfig](data, "CustomFieldConfig")
}

// UpdateBugSelectFieldOptions 更新缺陷下拉类型自定义字段候选值
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_bug_select_field_options.html
func (c *Client) UpdateBugSelectFieldOptions(ctx context.Context, req *model.UpdateSelectFieldOptionsRequest) error {
	_, err := c.doPost(ctx, "/custom_field_configs/update_bug_select_field_options", req.ToParams())
	return err
}

// UpdateStorySelectFieldOptions 更新需求下拉类型自定义字段候选值
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_story_select_field_options.html
func (c *Client) UpdateStorySelectFieldOptions(ctx context.Context, req *model.UpdateSelectFieldOptionsRequest) error {
	_, err := c.doPost(ctx, "/custom_field_configs/update_story_select_field_options", req.ToParams())
	return err
}

// UpdateCascadeFieldOptions 更新级联自定义字段候选值
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_cascade_field_options.html
func (c *Client) UpdateCascadeFieldOptions(ctx context.Context, req *model.UpdateCascadeFieldOptionsRequest) error {
	_, err := c.doPost(ctx, "/custom_field_configs/update_cascade_field_options", req.ToParams())
	return err
}

// GetWorkspaceSetting 获取项目配置开关，返回原始 key/value（如 {"workspace_metrology":"day"}）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/get_workspace_setting.html
func (c *Client) GetWorkspaceSetting(ctx context.Context, req *model.GetWorkspaceSettingRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/settings/get_workspace_setting", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse workspace setting: %w", err)
	}
	return result, nil
}

// UpdateSelectFieldOptions 更新下拉类型自定义字段候选值（统一接口，支持需求/缺陷/任务/迭代/测试用例）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_select_field_options.html
func (c *Client) UpdateSelectFieldOptions(ctx context.Context, req *model.UpdateSelectFieldOptionsUnifiedRequest) error {
	_, err := c.doPost(ctx, "/custom_field_configs/update_select_field_options", req.ToParams())
	return err
}

// CopyWorkitemTypeSetting 复制需求类别配置到目标项目
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/copy_workitem_type_setting.html
func (c *Client) CopyWorkitemTypeSetting(ctx context.Context, req *model.CopyWorkitemTypeSettingRequest) (*model.WorkitemType, error) {
	data, err := c.doPost(ctx, "/stories/copy_workitem_type_setting", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.WorkitemType](data, "WorkitemType")
}

// CopyBugSetting 复制缺陷配置到目标项目
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/copy_bug_setting.html
func (c *Client) CopyBugSetting(ctx context.Context, req *model.CopyBugSettingRequest) error {
	_, err := c.doPost(ctx, "/bugs/copy_settings", req.ToParams())
	return err
}
