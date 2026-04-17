// Package model 中的 release.go 定义了 TAPD 发布计划数据模型
package model

// Release 表示 TAPD 发布计划
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/
type Release struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startdate,omitempty"`
	EndDate     string `json:"enddate,omitempty"`
	Status      string `json:"status,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

// CreateReleaseRequest 创建发布计划的请求参数
type CreateReleaseRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：发布计划名称
	Description string // 可选：描述
	StartDate   string // 可选：开始日期
	EndDate     string // 可选：结束日期
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateReleaseRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "startdate", r.StartDate)
	setOptional(params, "enddate", r.EndDate)
	return params
}

// UpdateReleaseRequest 更新发布计划的请求参数
type UpdateReleaseRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：发布计划 ID
	Name        string // 可选：发布计划名称
	Description string // 可选：描述
	StartDate   string // 可选：开始日期
	EndDate     string // 可选：结束日期
	Status      string // 可选：状态
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateReleaseRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "startdate", r.StartDate)
	setOptional(params, "enddate", r.EndDate)
	setOptional(params, "status", r.Status)
	return params
}
