package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetWorkflowTransitions 获取工作流状态流转细则，返回强类型 []model.WorkflowTransition
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflow_all_transitions.html
func (c *Client) GetWorkflowTransitions(ctx context.Context, req *model.WorkflowRequest) ([]model.WorkflowTransition, error) {
	data, err := c.doGet(ctx, "/workflows/all_transitions", req.ToParams())
	if err != nil {
		return nil, err
	}

	var transitions []model.WorkflowTransition
	if err := json.Unmarshal(data, &transitions); err != nil {
		return nil, fmt.Errorf("failed to parse workflow transitions: %w", err)
	}
	return transitions, nil
}

// GetWorkflowStatusMap 获取工作流状态中英文映射，返回 map[string]string（英文状态名→中文状态名）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflow_status_map.html
func (c *Client) GetWorkflowStatusMap(ctx context.Context, req *model.WorkflowRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/workflows/status_map", req.ToParams())
	if err != nil {
		return nil, err
	}

	var statusMap map[string]string
	if err := json.Unmarshal(data, &statusMap); err != nil {
		return nil, fmt.Errorf("failed to parse workflow status map: %w", err)
	}
	return statusMap, nil
}

// GetWorkflowLastSteps 获取工作流结束状态，返回 map[string]string（英文状态名→中文状态名）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflow_last_steps.html
func (c *Client) GetWorkflowLastSteps(ctx context.Context, req *model.WorkflowRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/workflows/last_steps", req.ToParams())
	if err != nil {
		return nil, err
	}

	var lastSteps map[string]string
	if err := json.Unmarshal(data, &lastSteps); err != nil {
		return nil, fmt.Errorf("failed to parse workflow last steps: %w", err)
	}
	return lastSteps, nil
}

// GetWorkflowAllLastSteps 获取所有结束状态，返回 map[string]string（英文状态名→中文状态名）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflow_all_last_steps.html
func (c *Client) GetWorkflowAllLastSteps(ctx context.Context, req *model.WorkflowRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/workflows/all_last_steps", req.ToParams())
	if err != nil {
		return nil, err
	}

	var allLastSteps map[string]string
	if err := json.Unmarshal(data, &allLastSteps); err != nil {
		return nil, fmt.Errorf("failed to parse workflow all last steps: %w", err)
	}
	return allLastSteps, nil
}

// GetWorkflowFirstStep 获取工作流起始状态，返回起始状态字符串
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflow_first_step.html
func (c *Client) GetWorkflowFirstStep(ctx context.Context, req *model.WorkflowRequest) (string, error) {
	data, err := c.doGet(ctx, "/workflows/first_step", req.ToParams())
	if err != nil {
		return "", err
	}

	var firstStep string
	if err := json.Unmarshal(data, &firstStep); err != nil {
		return "", fmt.Errorf("failed to parse workflow first step: %w", err)
	}
	return firstStep, nil
}

// GetWorkflows 获取项目下的工作流列表，返回 json.RawMessage（结构复杂，不做强类型解析）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflows.html
func (c *Client) GetWorkflows(ctx context.Context, req *model.WorkflowRequest) (json.RawMessage, error) {
	data, err := c.doGet(ctx, "/workflows", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}
