// Package model 中的 label.go 定义了标签相关 API 的请求/响应类型
package model

// LabelPool 表示 TAPD 标签
type LabelPool struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	ColorValue  string `json:"color_value,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Modifier    string `json:"modifier,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

// AddLabelRequest 创建标签的请求参数
type AddLabelRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：标签名称（不能包括英文竖线）
	Color       string // 可选：颜色标识，枚举值 [1,2,3,4]
	Creator     string // 可选：创建人
}

func (r *AddLabelRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "color", r.Color)
	setOptional(params, "creator", r.Creator)
	return params
}

// UpdateLabelRequest 更新标签的请求参数（不支持更新标签名称）
type UpdateLabelRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：标签 ID
	Color       string // 可选：颜色标识，枚举值 [1,2,3,4]
	Modifier    string // 可选：更新人
}

func (r *UpdateLabelRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "color", r.Color)
	setOptional(params, "modifier", r.Modifier)
	return params
}

// QueryLabelRequest 查询标签列表的请求参数
type QueryLabelRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：标签 ID（支持多 ID 查询）
	Name        string // 可选：标签名称（支持模糊匹配）
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间（支持时间查询）
	Limit       int    // 可选：返回数量限制（默认 30，最大 200）
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
}

func (r *QueryLabelRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	return params
}

// CountLabelRequest 查询标签数量的请求参数
type CountLabelRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：标签 ID（支持多 ID 查询）
	Name        string // 可选：标签名称（支持模糊匹配）
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间（支持时间查询）
}

func (r *CountLabelRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	return params
}
