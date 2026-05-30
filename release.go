package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateRelease 创建发布计划
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/add_new_release.html
func (c *Client) CreateRelease(ctx context.Context, req *model.CreateReleaseRequest) (*model.Release, error) {
	data, err := c.doPost(ctx, "/new_releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Release](data, "Release")
}

// UpdateRelease 更新发布计划
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/update_new_release.html
func (c *Client) UpdateRelease(ctx context.Context, req *model.UpdateReleaseRequest) (*model.Release, error) {
	data, err := c.doPost(ctx, "/new_releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Release](data, "Release")
}

// CountReleases 获取发布计划数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_new_releases_count.html
func (c *Client) CountReleases(ctx context.Context, req *model.CountReleasesRequest) (int, error) {
	data, err := c.doGet(ctx, "/new_releases/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// GetLaunchForms 获取发布评审列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_launch_forms.html
func (c *Client) GetLaunchForms(ctx context.Context, req *model.GetLaunchFormsRequest) ([]model.LaunchForm, error) {
	data, err := c.doGet(ctx, "/launch_forms", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.LaunchForm](data, "LaunchForm")
}

// UpdateLaunchForm 更新发布评审单
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/update_launch_form.html
func (c *Client) UpdateLaunchForm(ctx context.Context, req *model.UpdateLaunchFormRequest) (*model.LaunchForm, error) {
	data, err := c.doPost(ctx, "/launch_forms", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.LaunchForm](data, "LaunchForm")
}

// CountLaunchForms 获取发布评审数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_launch_forms_count.html
func (c *Client) CountLaunchForms(ctx context.Context, req *model.CountLaunchFormsRequest) (int, error) {
	data, err := c.doGet(ctx, "/launch_forms/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}

// CreateLaunchForm 创建发布评审单
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/add_launch_form.html
func (c *Client) CreateLaunchForm(ctx context.Context, req *model.CreateLaunchFormRequest) (*model.LaunchForm, error) {
	data, err := c.doPost(ctx, "/launch_forms", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.LaunchForm](data, "LaunchForm")
}

// GetLaunchFormsTemplates 获取发布评审模板列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_launch_forms_templates.html
func (c *Client) GetLaunchFormsTemplates(ctx context.Context, req *model.GetLaunchFormsTemplatesRequest) ([]model.LaunchFormTemplate, error) {
	data, err := c.doGet(ctx, "/launch_forms/templates", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.LaunchFormTemplate](data, "template")
}

// GetLaunchFormsActivityLogs 获取发布评审活动日志
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_launch_forms_activity_logs.html
func (c *Client) GetLaunchFormsActivityLogs(ctx context.Context, req *model.GetLaunchFormsActivityLogsRequest) ([]model.LaunchFormActivityLog, error) {
	data, err := c.doGet(ctx, "/launch_forms/get_activity_logs", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.LaunchFormActivityLog](data, "LaunchChange")
}

// CreateLaunchAccessory 创建发布评审附件
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/add_launch_accessory.html
func (c *Client) CreateLaunchAccessory(ctx context.Context, req *model.CreateLaunchAccessoryRequest) (*model.LaunchAccessory, error) {
	data, err := c.doPost(ctx, "/launch_accessories", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.LaunchAccessory](data, "LaunchAccessory")
}

// GetLaunchAccessories 获取发布评审附件列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_launch_accessories.html
func (c *Client) GetLaunchAccessories(ctx context.Context, req *model.GetLaunchAccessoriesRequest) ([]model.LaunchAccessory, error) {
	data, err := c.doGet(ctx, "/launch_accessories", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.LaunchAccessory](data, "LaunchAccessory")
}

// GetLaunchFormsCustomFieldsSettings 获取发布评审自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_launch_forms_custom_fields_settings.html
func (c *Client) GetLaunchFormsCustomFieldsSettings(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.CustomFieldConfig, error) {
	data, err := c.doGet(ctx, "/launch_forms/custom_fields_settings", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.CustomFieldConfig](data, "CustomFieldConfig")
}
