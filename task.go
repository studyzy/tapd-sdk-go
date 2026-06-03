package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListTasks 查询任务列表，返回强类型 Task 切片，自定义字段通过 CustomFields map 保留
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks.html
func (c *Client) ListTasks(ctx context.Context, req *model.ListTasksRequest) ([]model.Task, error) {
	data, err := c.doGet(ctx, "/tasks", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Task](data, "Task")
}

// GetTask 获取单个任务详情，description 字段保留原始 HTML
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks.html
func (c *Client) GetTask(ctx context.Context, workspaceID, id string) (*model.Task, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
		"id":           id,
	}

	data, err := c.doGet(ctx, "/tasks", params)
	if err != nil {
		return nil, err
	}

	tasks, err := parseList[model.Task](data, "Task")
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, &TAPDError{ExitCode: 2, Message: fmt.Sprintf("task %s not found", id)}
	}

	task := tasks[0]
	task.URL = fmt.Sprintf("%s/%s/prong/tasks/view/%s", c.webURL, workspaceID, id)

	return &task, nil
}

// CreateTask 创建任务，返回创建后的完整 Task 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/add_task.html
func (c *Client) CreateTask(ctx context.Context, req *model.CreateTaskRequest) (*model.Task, error) {
	data, err := c.doPost(ctx, "/tasks", req.ToParams())
	if err != nil {
		return nil, err
	}

	task, err := parseOne[model.Task](data, "Task")
	if err != nil {
		return nil, err
	}

	task.URL = fmt.Sprintf("%s/%s/prong/tasks/view/%s", c.webURL, req.WorkspaceID, task.ID)

	return task, nil
}

// UpdateTask 更新任务
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/update_task.html
func (c *Client) UpdateTask(ctx context.Context, req *model.UpdateTaskRequest) (*model.Task, error) {
	data, err := c.doPost(ctx, "/tasks", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Task](data, "Task")
}

// CountTasks 查询任务数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks_count.html
func (c *Client) CountTasks(ctx context.Context, req *model.CountTasksRequest) (int, error) {
	data, err := c.doGet(ctx, "/tasks/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// BatchUpdateTask 批量更新任务，最多 50 条，返回服务端 msg
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/batch_update_task.html
func (c *Client) BatchUpdateTask(ctx context.Context, req *model.BatchUpdateTaskRequest) (string, error) {
	if len(req.Workitems) > 50 {
		return "", fmt.Errorf("batch_update_task: workitems count %d exceeds limit 50", len(req.Workitems))
	}
	data, err := c.doPost(ctx, "/tasks/batch_update_task", req.ToParams())
	if err != nil {
		return "", err
	}

	var result struct {
		Msg string `json:"msg"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("failed to parse batch update response: %w", err)
	}
	return result.Msg, nil
}

// GetRemovedTasks 获取回收站中的任务
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_removed_tasks.html
func (c *Client) GetRemovedTasks(ctx context.Context, req *model.GetRemovedTasksRequest) ([]model.RemovedTask, error) {
	data, err := c.doGet(ctx, "/tasks/get_removed_tasks", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.RemovedTask](data, "RemovedTask")
}

// GetTasksByViewConfID 获取视图对应的任务列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks_by_view_conf_id.html
func (c *Client) GetTasksByViewConfID(ctx context.Context, req *model.GetTasksByViewConfIDRequest) ([]model.Task, error) {
	data, err := c.doGet(ctx, "/tasks/get_tasks_by_view_conf_id", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Task](data, "Task")
}

// GetTaskFieldsInfo 获取任务所有字段及候选值
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/task/get_task_fields_info.html
func (c *Client) GetTaskFieldsInfo(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]model.FieldInfo, error) {
	data, err := c.doGet(ctx, "/tasks/get_fields_info", req.ToParams())
	if err != nil {
		return nil, err
	}
	var fields map[string]model.FieldInfo
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, fmt.Errorf("failed to parse task fields info: %w", err)
	}
	return fields, nil
}

// GetTaskCustomFieldsSettings 获取任务自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/task/get_task_custom_fields_settings.html
func (c *Client) GetTaskCustomFieldsSettings(ctx context.Context, workspaceID string) ([]model.CustomFieldConfig, error) {
	return c.GetCustomFields(ctx, &model.GetCustomFieldsRequest{
		WorkspaceID: workspaceID,
		EntityType:  "tasks",
	})
}
