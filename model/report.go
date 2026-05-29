// Package model 中的 report.go 定义了 TAPD 项目报告数据模型
package model

// WorkspaceReport 表示 TAPD 项目报告
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/report/get_workspace_reports.html
type WorkspaceReport struct {
	ID                      string `json:"id,omitempty"`
	Title                   string `json:"title,omitempty"`
	WorkspaceID             string `json:"workspace_id,omitempty"`
	ReportType              string `json:"report_type,omitempty"`
	Receiver                string `json:"receiver,omitempty"`
	CC                      string `json:"cc,omitempty"`
	ReceiverOrganizationIDs string `json:"receiver_organization_ids,omitempty"`
	CCOrganizationIDs       string `json:"cc_organization_ids,omitempty"`
	Sender                  string `json:"sender,omitempty"`
	SendTime                string `json:"send_time,omitempty"`
	Author                  string `json:"author,omitempty"`
	Created                 string `json:"created,omitempty"`
	Status                  string `json:"status,omitempty"`
	Modified                string `json:"modified,omitempty"`
	LastModify              string `json:"last_modify,omitempty"`
}

// GetWorkspaceReportsRequest 获取项目报告的请求参数
type GetWorkspaceReportsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：报告 ID
	Title       string // 可选：标题
	Author      string // 可选：创建人
	Created     string // 可选：创建时间
	Limit       int    // 可选：返回数量限制（默认 30，最大 200）
	Page        int    // 可选：页码（默认 1）
	Fields      string // 可选：返回字段列表（逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetWorkspaceReportsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "title", r.Title)
	setOptional(params, "author", r.Author)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "fields", r.Fields)
	return params
}
