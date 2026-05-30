package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListWorkspaces 获取当前用户参与的项目列表，自动过滤 category 为 organization 的条目
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/user_participant_projects.html
func (c *Client) ListWorkspaces(ctx context.Context, companyID string) ([]model.Workspace, error) {
	params := map[string]string{
		"company_id": companyID,
	}
	if nick := c.GetNick(); nick != "" {
		params["nick"] = nick
	}
	data, err := c.doGet(ctx, "/workspaces/user_participant_projects", params)
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: [{"Workspace": {...}}, ...]
	all, err := parseList[model.Workspace](data, "Workspace")
	if err != nil {
		return nil, err
	}

	workspaces := make([]model.Workspace, 0, len(all))
	for _, ws := range all {
		if ws.Category != "organization" {
			workspaces = append(workspaces, ws)
		}
	}
	return workspaces, nil
}

// GetWorkspaceInfo 获取指定工作区的详细信息
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_workspace_info.html
func (c *Client) GetWorkspaceInfo(ctx context.Context, workspaceID string) (*model.Workspace, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
	}
	data, err := c.doGet(ctx, "/workspaces/get_workspace_info", params)
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: {"Workspace": {...}} (单个对象)
	var wrapper map[string]model.Workspace
	if err := json.Unmarshal(data, &wrapper); err != nil {
		// 尝试数组格式: [{"Workspace": {...}}]
		var rawList []map[string]model.Workspace
		if err2 := json.Unmarshal(data, &rawList); err2 != nil {
			return nil, fmt.Errorf("failed to parse workspace info: %w", err)
		}
		if len(rawList) == 0 {
			return nil, &TAPDError{ExitCode: 2, Message: "workspace not found"}
		}
		if ws, ok := rawList[0]["Workspace"]; ok {
			return &ws, nil
		}
		return nil, fmt.Errorf("unexpected response format")
	}

	if ws, ok := wrapper["Workspace"]; ok {
		return &ws, nil
	}
	return nil, &TAPDError{ExitCode: 2, Message: "workspace not found"}
}

// GetSubWorkspaces 获取子项目信息
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/sub_workspaces.html
func (c *Client) GetSubWorkspaces(ctx context.Context, req *model.GetSubWorkspacesRequest) (*model.Workspace, error) {
	data, err := c.doGet(ctx, "/workspaces/sub_workspaces", req.ToParams())
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: {"Workspace": {...}}
	var wrapper map[string]model.Workspace
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse sub workspaces: %w", err)
	}
	if ws, ok := wrapper["Workspace"]; ok {
		return &ws, nil
	}
	return nil, &TAPDError{ExitCode: 2, Message: "sub workspace not found"}
}

// ListCompanyProjects 获取公司项目列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/projects.html
func (c *Client) ListCompanyProjects(ctx context.Context, req *model.ListCompanyProjectsRequest) ([]model.Workspace, error) {
	data, err := c.doGet(ctx, "/workspaces/projects", req.ToParams())
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: [{"Workspace": {...}}, ...]
	return parseList[model.Workspace](data, "Workspace")
}

// GetWorkspaceUsers 获取指定项目成员
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/users.html
func (c *Client) GetWorkspaceUsers(ctx context.Context, req *model.GetWorkspaceUsersRequest) ([]model.UserWorkspace, error) {
	data, err := c.doGet(ctx, "/workspaces/users", req.ToParams())
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: [{"UserWorkspace": {...}}, ...]
	return parseList[model.UserWorkspace](data, "UserWorkspace")
}

// AddWorkspaceMember 添加项目成员
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/add_workspace_member.html
func (c *Client) AddWorkspaceMember(ctx context.Context, req *model.AddWorkspaceMemberRequest) (*model.SuccessResponse, error) {
	data, err := c.doPost(ctx, "/workspaces/add_workspace_member", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.SuccessResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse add member response: %w", err)
	}
	return &result, nil
}

// GetRoles 获取用户组ID对照关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/roles.html
func (c *Client) GetRoles(ctx context.Context, workspaceID string) (map[string]string, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
	}
	data, err := c.doGet(ctx, "/roles", params)
	if err != nil {
		return nil, err
	}

	var roles map[string]string
	if err := json.Unmarshal(data, &roles); err != nil {
		return nil, fmt.Errorf("failed to parse roles: %w", err)
	}
	return roles, nil
}

// CreateMiniProject 新建空间
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/workspace/create_mini_project.html
func (c *Client) CreateMiniProject(ctx context.Context, req *model.CreateMiniProjectRequest) (*model.CreateMiniProjectResponse, error) {
	data, err := c.doPost(ctx, "/workspaces/create_mini_project", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.CreateMiniProjectResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse create mini project response: %w", err)
	}
	return &result, nil
}

// GetMiniProjectList 获取用户所有参与的空间
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/workspace/get_mini_project_list_with_permission.html
func (c *Client) GetMiniProjectList(ctx context.Context, req *model.GetMiniProjectListRequest) ([]model.MiniProject, error) {
	data, err := c.doGet(ctx, "/workspaces/get_mini_project_list_with_permission", req.ToParams())
	if err != nil {
		return nil, err
	}

	var projects []model.MiniProject
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, fmt.Errorf("failed to parse mini project list: %w", err)
	}
	return projects, nil
}

// EnableWorkCalendar 设置启用工作日历
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/enable_work_calendar.html
func (c *Client) EnableWorkCalendar(ctx context.Context, req *model.EnableWorkCalendarRequest) (*model.SuccessResponse, error) {
	data, err := c.doPost(ctx, "/workspaces/enable_work_calendar", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.SuccessResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse enable work calendar response: %w", err)
	}
	return &result, nil
}

// GetCustomWorkCalendar 获取自定义工作日历详情
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_custom_work_calendar.html
func (c *Client) GetCustomWorkCalendar(ctx context.Context, req *model.GetCustomWorkCalendarRequest) (*model.CustomWorkCalendar, error) {
	data, err := c.doGet(ctx, "/workspaces/get_custom_work_calendar", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.CustomWorkCalendar
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse custom work calendar: %w", err)
	}
	return &result, nil
}

// SetCustomWorkCalendar 设置自定义工作日历
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/set_custom_work_calendar.html
func (c *Client) SetCustomWorkCalendar(ctx context.Context, req *model.SetCustomWorkCalendarRequest) (*model.SuccessResponse, error) {
	data, err := c.doPost(ctx, "/workspaces/set_custom_work_calendar", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.SuccessResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse set custom work calendar response: %w", err)
	}
	return &result, nil
}

// GetWorkCalendarSettings 获取工作日历设置列表及启用选项
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_work_calendar_settings.html
func (c *Client) GetWorkCalendarSettings(ctx context.Context, workspaceID string) ([]model.WorkCalendarSetting, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
	}
	data, err := c.doGet(ctx, "/workspaces/get_work_calendar_settings", params)
	if err != nil {
		return nil, err
	}

	var settings []model.WorkCalendarSetting
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse work calendar settings: %w", err)
	}
	return settings, nil
}

// GetWorkitemsLongIDByShortIDs 通过工作项短 id 换长 id（也支持反向校验）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_workitems_long_id_by_short_ids.html
func (c *Client) GetWorkitemsLongIDByShortIDs(ctx context.Context, req *model.GetWorkitemsLongIDByShortIDsRequest) (*model.GetWorkitemsLongIDByShortIDsResponse, error) {
	data, err := c.doGet(ctx, "/workspaces/get_workitems_long_id_by_short_ids", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.GetWorkitemsLongIDByShortIDsResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse workitems id map: %w", err)
	}
	return &result, nil
}

// GetWorkspaceDocuments 获取项目文档
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_workspace_documents.html
func (c *Client) GetWorkspaceDocuments(ctx context.Context, req *model.GetWorkspaceDocumentsRequest) ([]model.Document, error) {
	data, err := c.doGet(ctx, "/documents/get_workspace_documents", req.ToParams())
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: [{"Document": {...}}, ...]
	return parseList[model.Document](data, "Document")
}

// UpdateWorkspaceInfo 更新项目信息
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/update_workspace_info.html
func (c *Client) UpdateWorkspaceInfo(ctx context.Context, req *model.UpdateWorkspaceInfoRequest) (string, error) {
	data, err := c.doPost(ctx, "/workspaces/update_workspace_info", req.ToParams())
	if err != nil {
		return "", err
	}
	var msg string
	if err := json.Unmarshal(data, &msg); err != nil {
		return "", err
	}
	return msg, nil
}

// GetWorkspaceCustomFieldSettings 获取项目自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/workspace_custom_field_settings.html
func (c *Client) GetWorkspaceCustomFieldSettings(ctx context.Context, workspaceID string) ([]model.CustomFieldConfig, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
	}
	data, err := c.doGet(ctx, "/workspaces/workspace_custom_field_settings", params)
	if err != nil {
		return nil, err
	}

	// TAPD 返回格式: [{"CustomFieldConfig": {...}}, ...]
	return parseList[model.CustomFieldConfig](data, "CustomFieldConfig")
}
