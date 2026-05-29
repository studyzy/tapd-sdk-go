// Package model 中的 pipeline.go 定义了流水线关联相关 API 的请求/响应类型
package model

// ThirdRelation 表示 TAPD 业务对象与流水线构建记录的关联关系
type ThirdRelation struct {
	ID              string `json:"id,omitempty"`
	WorkspaceID     string `json:"workspace_id,omitempty"`
	SourceAppID     string `json:"source_app_id,omitempty"`
	SourceProjectID string `json:"source_project_id,omitempty"`
	SourceID        string `json:"source_id,omitempty"`
	SourceType      string `json:"source_type,omitempty"`
	SourceData      string `json:"source_data,omitempty"`
	SourceTag       string `json:"source_tag,omitempty"`
	TapdID          string `json:"tapd_id,omitempty"`
	WorkitemTypeID  string `json:"workitem_type_id,omitempty"`
	TapdType        string `json:"tapd_type,omitempty"`
	Created         string `json:"created,omitempty"`
	Modified        string `json:"modified,omitempty"`
	Status          string `json:"status,omitempty"`
	Operator        string `json:"operator,omitempty"`
}

// AddThirdRelationRequest 创建流水线关联关系的请求参数
type AddThirdRelationRequest struct {
	SourceType      string // 必填：资源类型（目前可选值：build）
	SourceProjectID string // 必填：流水线 ID
	SourceID        string // 必填：流水线构建记录 ID
	WorkspaceID     string // 必填：TAPD 项目 ID
	TapdID          string // 必填：TAPD 业务对象 ID
	TapdType        string // 必填：TAPD 业务对象类型（task/story/bug）
	Operator        string // 必填：操作人（英文名）
}

func (r *AddThirdRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"source_type":       r.SourceType,
		"source_project_id": r.SourceProjectID,
		"source_id":         r.SourceID,
		"workspace_id":      r.WorkspaceID,
		"tapd_id":           r.TapdID,
		"tapd_type":         r.TapdType,
		"operator":          r.Operator,
	}
}

// GetThirdRelationsRequest 获取流水线关联关系的请求参数
type GetThirdRelationsRequest struct {
	SourceType  string // 必填：资源类型（目前可选值：build）
	WorkspaceID string // 必填：TAPD 项目 ID
	TapdID      string // 必填：TAPD 对象 ID
	TapdType    string // 必填：TAPD 对象类型（story/task/bug）
}

func (r *GetThirdRelationsRequest) ToParams() map[string]string {
	return map[string]string{
		"source_type":  r.SourceType,
		"workspace_id": r.WorkspaceID,
		"tapd_id":      r.TapdID,
		"tapd_type":    r.TapdType,
	}
}

// DeleteThirdRelationRequest 删除流水线关联关系的请求参数
type DeleteThirdRelationRequest struct {
	ID       string // 必填：关联 ID（列表接口返回的 id）
	Operator string // 必填：操作人（英文名）
}

func (r *DeleteThirdRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"id":       r.ID,
		"operator": r.Operator,
	}
}
