// Package model 中的 workspace.go 定义了项目（workspace）相关 API 的请求/响应类型
package model

import "strings"

// EnableWorkCalendarRequest 启用工作日历的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/enable_work_calendar.html
type EnableWorkCalendarRequest struct {
	WorkspaceID string // 必填：项目 ID
	Type        string // 必填：启用类型，可选范围 system 或 custom
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *EnableWorkCalendarRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"type":         r.Type,
	}
}

// GetCustomWorkCalendarRequest 获取自定义工作日历详情的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_custom_work_calendar.html
type GetCustomWorkCalendarRequest struct {
	WorkspaceID string // 必填：项目 ID
	Year        string // 必填：年份
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetCustomWorkCalendarRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"year":         r.Year,
	}
}

// CustomWorkCalendar 自定义工作日历详情
type CustomWorkCalendar struct {
	Weekdays []string `json:"weekdays,omitempty"` // 一周内被设置为工作日的日期（1-7）
	Holidays []string `json:"holidays,omitempty"` // 假日
	Workdays []string `json:"workdays,omitempty"` // 工作日
}

// WorkCalendarSetting 工作日历设置项
type WorkCalendarSetting struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Enable bool   `json:"enable,omitempty"`
}

// SetCustomWorkCalendarRequest 设置自定义工作日历的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/set_custom_work_calendar.html
type SetCustomWorkCalendarRequest struct {
	WorkspaceID string   // 必填：项目 ID
	Year        string   // 必填：年份
	Weekdays    []string // 可选：一周内的工作日，取值 1-7
	Holidays    []string // 可选：额外节假日（YYYY-MM-DD）
	Workdays    []string // 可选：额外工作日（YYYY-MM-DD）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
// 数组参数序列化为 JSON 数组格式 [a,b,c]
func (r *SetCustomWorkCalendarRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"year":         r.Year,
	}
	if len(r.Weekdays) > 0 {
		params["weekdays"] = "[" + strings.Join(r.Weekdays, ",") + "]"
	}
	if len(r.Holidays) > 0 {
		params["holidays"] = "[" + quoteJoin(r.Holidays) + "]"
	}
	if len(r.Workdays) > 0 {
		params["workdays"] = "[" + quoteJoin(r.Workdays) + "]"
	}
	return params
}

// quoteJoin 将字符串数组用双引号包裹后逗号拼接
func quoteJoin(items []string) string {
	parts := make([]string, len(items))
	for i, s := range items {
		parts[i] = "\"" + s + "\""
	}
	return strings.Join(parts, ",")
}

// GetWorkitemsLongIDByShortIDsRequest 通过工作项短 id 换长 id 的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_workitems_long_id_by_short_ids.html
// short_ids 与 long_ids 不允许都不传
type GetWorkitemsLongIDByShortIDsRequest struct {
	WorkspaceID string // 必填：项目 ID
	EntityType  string // 必填：业务对象类型，候选值 story, task, bug
	ShortIDs    string // 可选：短 ID，多个以 ; 分隔
	LongIDs     string // 可选：长 ID，多个以 ; 分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetWorkitemsLongIDByShortIDsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"entity_type":  r.EntityType,
	}
	if r.ShortIDs != "" {
		params["short_ids"] = r.ShortIDs
	}
	if r.LongIDs != "" {
		params["long_ids"] = r.LongIDs
	}
	return params
}

// WorkitemIDMap 短 id 与长 id 的映射条目
type WorkitemIDMap struct {
	ShortID     string `json:"short_id,omitempty"`
	LongID      string `json:"long_id,omitempty"`
	EntityType  string `json:"entity_type,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	CompanyID   string `json:"company_id,omitempty"`
}

// GetWorkitemsLongIDByShortIDsResponse 通过短 id 换长 id 接口响应
type GetWorkitemsLongIDByShortIDsResponse struct {
	ValidIDMap      []WorkitemIDMap `json:"valid_id_map,omitempty"`
	InvalidLongIDs  []string        `json:"invalid_long_ids,omitempty"`
	InvalidShortIDs []string        `json:"invalid_short_ids,omitempty"`
}

// GetWorkspaceDocumentsRequest 获取项目文档的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workspace/get_workspace_documents.html
type GetWorkspaceDocumentsRequest struct {
	WorkspaceID string // 必填：项目 ID
	Limit       string // 可选：每页数量（默认 30，最大 200）
	Page        string // 可选：页码（默认 1）
	Fields      string // 可选：返回字段列表（逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetWorkspaceDocumentsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	if r.Limit != "" {
		params["limit"] = r.Limit
	}
	if r.Page != "" {
		params["page"] = r.Page
	}
	if r.Fields != "" {
		params["fields"] = r.Fields
	}
	return params
}

// Document 项目文档信息
type Document struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	FolderID    string `json:"folder_id,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Modifier    string `json:"modifier,omitempty"`
	Status      string `json:"status,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
}

// UpdateWorkspaceInfoRequest 更新项目信息的请求参数
// 只能通过 field+value 模式更新单个字段，field 可选值：description/begin_date/end_date/begin_end
type UpdateWorkspaceInfoRequest struct {
	WorkspaceID string // 必填：项目 ID
	Field       string // 必填：字段名（description/begin_date/end_date/begin_end）
	Value       string // 必填：字段值
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateWorkspaceInfoRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"field":        r.Field,
		"value":        r.Value,
	}
	return params
}
