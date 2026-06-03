package tapd

import (
	"context"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetStoryChanges 查询需求变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes.html
func (c *Client) GetStoryChanges(ctx context.Context, req *model.GetStoryChangesRequest) ([]model.WorkitemChange, error) {
	data, err := c.doGet(ctx, "/story_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.WorkitemChange](data, "WorkitemChange")
}

// CountStoryChanges 查询需求变更数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes_count.html
func (c *Client) CountStoryChanges(ctx context.Context, req *model.CountStoryChangesRequest) (int, error) {
	data, err := c.doGet(ctx, "/story_changes/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// GetBugChanges 查询缺陷变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_changes.html
func (c *Client) GetBugChanges(ctx context.Context, req *model.GetBugChangesRequest) ([]model.BugChange, error) {
	data, err := c.doGet(ctx, "/bug_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.BugChange](data, "BugChange")
}

// CountBugChanges 查询缺陷变更数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_changes_count.html
func (c *Client) CountBugChanges(ctx context.Context, req *model.CountBugChangesRequest) (int, error) {
	data, err := c.doGet(ctx, "/bug_changes/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// GetTaskChanges 查询任务变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes.html
// 注意：task_id 与 created 参数二选一必填
func (c *Client) GetTaskChanges(ctx context.Context, req *model.GetTaskChangesRequest) ([]model.WorkitemChange, error) {
	if req.TaskID == "" && req.Created == "" {
		return nil, fmt.Errorf("get_task_changes: task_id and created cannot both be empty, at least one is required")
	}
	data, err := c.doGet(ctx, "/task_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.WorkitemChange](data, "WorkitemChange")
}

// CountTaskChanges 查询任务变更数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes_count.html
func (c *Client) CountTaskChanges(ctx context.Context, req *model.CountTaskChangesRequest) (int, error) {
	data, err := c.doGet(ctx, "/task_changes/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// GetIterationChanges 查询迭代变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iteration_changes.html
func (c *Client) GetIterationChanges(ctx context.Context, req *model.GetIterationChangesRequest) ([]model.IterationChange, error) {
	data, err := c.doGet(ctx, "/iteration_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.IterationChange](data, "IterationChange")
}
