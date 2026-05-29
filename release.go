package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateRelease 创建发布计划
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/add_release.html
func (c *Client) CreateRelease(ctx context.Context, req *model.CreateReleaseRequest) (*model.Release, error) {
	data, err := c.doPost(ctx, "/releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Release](data, "Release")
}

// UpdateRelease 更新发布计划
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/update_release.html
func (c *Client) UpdateRelease(ctx context.Context, req *model.UpdateReleaseRequest) (*model.Release, error) {
	data, err := c.doPost(ctx, "/releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Release](data, "Release")
}

// CountReleases 获取发布计划数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_releases_count.html
func (c *Client) CountReleases(ctx context.Context, req *model.WorkspaceIDRequest) (int, error) {
	data, err := c.doGet(ctx, "/releases/count", req.ToParams())
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
