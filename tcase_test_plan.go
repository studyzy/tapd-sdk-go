package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateTestPlan 创建测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_test_plan.html
func (c *Client) CreateTestPlan(ctx context.Context, req *model.CreateTestPlanRequest) (*model.TestPlan, error) {
	data, err := c.doPost(ctx, "/test_plans", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.TestPlan](data, "TestPlan")
}

// UpdateTestPlan 编辑测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/update_test_plan.html
func (c *Client) UpdateTestPlan(ctx context.Context, req *model.UpdateTestPlanRequest) (*model.TestPlan, error) {
	data, err := c.doPost(ctx, "/test_plans", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.TestPlan](data, "TestPlan")
}

// ListTestPlans 查询测试计划列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plans.html
func (c *Client) ListTestPlans(ctx context.Context, req *model.ListTestPlansRequest) ([]model.TestPlan, error) {
	data, err := c.doGet(ctx, "/test_plans", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.TestPlan](data, "TestPlan")
}

// CountTestPlans 查询测试计划数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plans_count.html
func (c *Client) CountTestPlans(ctx context.Context, req *model.ListTestPlansRequest) (int, error) {
	data, err := c.doGet(ctx, "/test_plans/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// GetTestPlanDetails 获取测试计划测试结果
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_details.html
func (c *Client) GetTestPlanDetails(ctx context.Context, req *model.TestPlanIDRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/test_plans/details", req.ToParams())
}

// GetTestPlanFieldsInfo 获取测试计划所有字段及候选值
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_fields_info.html
func (c *Client) GetTestPlanFieldsInfo(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]model.FieldInfo, error) {
	data, err := c.doGet(ctx, "/test_plans/get_fields_info", req.ToParams())
	if err != nil {
		return nil, err
	}

	var fields map[string]model.FieldInfo
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

// GetTestPlanProgress 获取测试计划执行进度
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_progress.html
func (c *Client) GetTestPlanProgress(ctx context.Context, req *model.TestPlanIDRequest) (*model.TestPlanProgress, error) {
	data, err := c.doGet(ctx, "/test_plans/progress", req.ToParams())
	if err != nil {
		return nil, err
	}

	var progress model.TestPlanProgress
	if err := json.Unmarshal(data, &progress); err != nil {
		return nil, err
	}
	return &progress, nil
}

// GetTestPlanRelativeStories 获取测试计划关联的需求
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_relative_stories.html
func (c *Client) GetTestPlanRelativeStories(ctx context.Context, req *model.TestPlanRelativeStoriesRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/test_plans/get_relative_stories", req.ToParams())
}

// ListTestPlanTCaseRelations 获取测试计划与测试用例关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_tcases.html
func (c *Client) ListTestPlanTCaseRelations(ctx context.Context, req *model.TestPlanTCasesRequest) ([]model.TestPlanStoryTcaseRelation, error) {
	data, err := c.doGet(ctx, "/test_plans/get_test_plan_tcase", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.TestPlanStoryTcaseRelation](data, "TestPlanStoryTcaseRelation")
}

// GetTestPlanBugs 获取测试计划关联 bug
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_bugs.html
func (c *Client) GetTestPlanBugs(ctx context.Context, req *model.TestPlanIDRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/test_plans/result_relation_bugs", req.ToParams())
}

// ListTestPlansByIterationID 获取迭代下测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_by_iteration_id.html
func (c *Client) ListTestPlansByIterationID(ctx context.Context, req *model.TestPlansByIterationIDRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/test_plans/get_by_iteration_id", req.ToParams())
}

// CreateTestPlanStoryRelation 创建测试计划和需求关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/create_story_relation.html
func (c *Client) CreateTestPlanStoryRelation(ctx context.Context, req *model.TestPlanStoryRelationRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/test_plans/create_story_relation", req.ToParams())
}

// DeleteTestPlanStoryRelation 解除测试计划和需求关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/delete_story_relation.html
func (c *Client) DeleteTestPlanStoryRelation(ctx context.Context, req *model.TestPlanStoryRelationRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/test_plans/delete_story_relation", req.ToParams())
}

// CreateTestPlanTCaseRelation 创建测试计划和测试用例关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/create_tcase_relation.html
func (c *Client) CreateTestPlanTCaseRelation(ctx context.Context, req *model.TestPlanTCaseRelationRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/test_plans/create_tcase_relation", req.ToParams())
}
