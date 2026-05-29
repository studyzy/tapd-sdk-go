package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListIterations 查询迭代列表，返回强类型 Iteration 切片
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations.html
func (c *Client) ListIterations(ctx context.Context, req *model.ListIterationsRequest) ([]model.Iteration, error) {
	data, err := c.doGet(ctx, "/iterations", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Iteration](data, "Iteration")
}

// CreateIteration 创建迭代，返回创建后的完整 Iteration 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/add_iteration.html
func (c *Client) CreateIteration(ctx context.Context, req *model.CreateIterationRequest) (*model.Iteration, error) {
	data, err := c.doPost(ctx, "/iterations", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Iteration](data, "Iteration")
}

// UpdateIteration 更新迭代，返回更新后的完整 Iteration 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/update_iteration.html
func (c *Client) UpdateIteration(ctx context.Context, req *model.UpdateIterationRequest) (*model.Iteration, error) {
	data, err := c.doPost(ctx, "/iterations", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Iteration](data, "Iteration")
}

// CountIterations 查询迭代数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations_count.html
func (c *Client) CountIterations(ctx context.Context, req *model.CountIterationsRequest) (int, error) {
	data, err := c.doGet(ctx, "/iterations/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// LockIteration 锁定迭代
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/lock_iteration.html
func (c *Client) LockIteration(ctx context.Context, req *model.LockIterationRequest) error {
	_, err := c.doPost(ctx, "/iterations/lock_iteration", req.ToParams())
	return err
}

// UnlockIteration 解锁迭代
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/unlock_iteration.html
func (c *Client) UnlockIteration(ctx context.Context, req *model.UnlockIterationRequest) error {
	_, err := c.doPost(ctx, "/iterations/unlock_iteration", req.ToParams())
	return err
}

// GetCustomDashBoardContent 获取迭代仪表盘自定义卡片内容
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/get_custom_dash_board_content.html
func (c *Client) GetCustomDashBoardContent(ctx context.Context, req *model.GetCustomDashBoardContentRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/iterations/get_custom_dash_board_content", req.ToParams())
}

// UpdateCustomDashBoardContent 修改迭代仪表盘自定义卡片内容
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/update_custom_dash_board_content.html
func (c *Client) UpdateCustomDashBoardContent(ctx context.Context, req *model.UpdateCustomDashBoardContentRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/iterations/update_custom_dash_board_content", req.ToParams())
}

// GetIterationTemplateList 获取迭代模板列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/template_list.html
func (c *Client) GetIterationTemplateList(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.WorkitemTemplate, error) {
	data, err := c.doGet(ctx, "/iterations/template_list", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplate](data, "WorkitemTemplate")
}

// GetIterationTemplateFields 获取迭代模板字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/template_fields.html
func (c *Client) GetIterationTemplateFields(ctx context.Context, req *model.GetIterationTemplateFieldsRequest) ([]model.WorkitemTemplateField, error) {
	data, err := c.doGet(ctx, "/iterations/template_fields", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplateField](data, "WorkitemTemplateField")
}

// GetIterationWorkitemTypes 获取迭代类别列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/workitem_types.html
func (c *Client) GetIterationWorkitemTypes(ctx context.Context, req *model.WorkspaceIDRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/iterations/workitem_types", req.ToParams())
}

// GetDefaultTemplateFieldsByWorkitemTypeID 获取迭代类别默认模板字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/default_template_fields_by_workitem_type_id.html
func (c *Client) GetDefaultTemplateFieldsByWorkitemTypeID(ctx context.Context, req *model.GetDefaultTemplateFieldsByWorkitemTypeIDRequest) ([]model.WorkitemTemplateField, error) {
	data, err := c.doGet(ctx, "/iterations/default_template_fields_by_workitem_type_id", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplateField](data, "WorkitemTemplateField")
}

// GetPlanApps 获取计划应用列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/get_plan_apps.html
func (c *Client) GetPlanApps(ctx context.Context, req *model.GetPlanAppsRequest) ([]model.PlanApp, error) {
	data, err := c.doGet(ctx, "/plan_apps", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.PlanApp](data, "PlanApp")
}

// CountPlanApps 获取计划应用数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/get_plan_apps_count.html
func (c *Client) CountPlanApps(ctx context.Context, req *model.CountPlanAppsRequest) (int, error) {
	data, err := c.doGet(ctx, "/plan_apps/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}
