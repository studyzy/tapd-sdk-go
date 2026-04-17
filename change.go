package tapd

import (
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetStoryChanges 查询需求变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes.html
func (c *Client) GetStoryChanges(req *model.GetStoryChangesRequest) ([]model.WorkitemChange, error) {
	data, err := c.doGet("/story_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse story changes list: %w", err)
	}

	var results []model.WorkitemChange
	for _, item := range rawList {
		if raw, ok := item["WorkitemChange"]; ok {
			var change model.WorkitemChange
			if err := json.Unmarshal(raw, &change); err == nil {
				results = append(results, change)
			}
		}
	}
	return results, nil
}

// CountStoryChanges 查询需求变更数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes_count.html
func (c *Client) CountStoryChanges(req *model.CountStoryChangesRequest) (int, error) {
	data, err := c.doGet("/story_changes/count", req.ToParams())
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

// GetBugChanges 查询缺陷变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_changes.html
func (c *Client) GetBugChanges(req *model.GetBugChangesRequest) ([]model.BugChange, error) {
	data, err := c.doGet("/bug_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse bug changes list: %w", err)
	}

	var results []model.BugChange
	for _, item := range rawList {
		if raw, ok := item["BugChange"]; ok {
			var change model.BugChange
			if err := json.Unmarshal(raw, &change); err == nil {
				results = append(results, change)
			}
		}
	}
	return results, nil
}

// CountBugChanges 查询缺陷变更数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_changes_count.html
func (c *Client) CountBugChanges(req *model.CountBugChangesRequest) (int, error) {
	data, err := c.doGet("/bug_changes/count", req.ToParams())
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

// GetTaskChanges 查询任务变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes.html
func (c *Client) GetTaskChanges(req *model.GetTaskChangesRequest) ([]model.WorkitemChange, error) {
	data, err := c.doGet("/task_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse task changes list: %w", err)
	}

	var results []model.WorkitemChange
	for _, item := range rawList {
		if raw, ok := item["WorkitemChange"]; ok {
			var change model.WorkitemChange
			if err := json.Unmarshal(raw, &change); err == nil {
				results = append(results, change)
			}
		}
	}
	return results, nil
}

// CountTaskChanges 查询任务变更数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes_count.html
func (c *Client) CountTaskChanges(req *model.CountTaskChangesRequest) (int, error) {
	data, err := c.doGet("/task_changes/count", req.ToParams())
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

// GetIterationChanges 查询迭代变更历史列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iteration_changes.html
func (c *Client) GetIterationChanges(req *model.GetIterationChangesRequest) ([]model.IterationChange, error) {
	data, err := c.doGet("/iteration_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse iteration changes list: %w", err)
	}

	var results []model.IterationChange
	for _, item := range rawList {
		if raw, ok := item["IterationChange"]; ok {
			var change model.IterationChange
			if err := json.Unmarshal(raw, &change); err == nil {
				results = append(results, change)
			}
		}
	}
	return results, nil
}
