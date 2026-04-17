// Package model 中的 board.go 定义了 TAPD 看板数据模型及请求参数结构体
package model

// BoardCard 表示 TAPD 看板工作项
type BoardCard struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	BoardID     string `json:"board_id,omitempty"`
	ColumnID    string `json:"column_id,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Status      string `json:"status,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

// BoardColumn 表示 TAPD 看板板块
type BoardColumn struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	BoardID     string `json:"board_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Sort        string `json:"sort,omitempty"`
}

// CreateBoardCardRequest 新建看板工作项的请求参数
type CreateBoardCardRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：名称
	BoardID     string // 必填：看板 ID
	ColumnID    string // 必填：板块 ID
	Description string // 可选：描述
	Owner       string // 可选：处理人
	Priority    string // 可选：优先级
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateBoardCardRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
		"board_id":     r.BoardID,
		"column_id":    r.ColumnID,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "priority", r.Priority)
	return params
}

// GetBoardCardsRequest 获取看板工作项的请求参数
type GetBoardCardsRequest struct {
	WorkspaceID string // 必填：项目 ID
	BoardID     string // 可选：看板 ID
	ColumnID    string // 可选：板块 ID
	Owner       string // 可选：处理人
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBoardCardsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "board_id", r.BoardID)
	setOptional(params, "column_id", r.ColumnID)
	setOptional(params, "owner", r.Owner)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// UpdateBoardCardRequest 更新看板工作项的请求参数
type UpdateBoardCardRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：工作项 ID
	Name        string // 可选：名称
	ColumnID    string // 可选：板块 ID
	Description string // 可选：描述
	Owner       string // 可选：处理人
	Priority    string // 可选：优先级
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateBoardCardRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "column_id", r.ColumnID)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "priority", r.Priority)
	return params
}

// GetBoardColumnsRequest 获取看板板块的请求参数
type GetBoardColumnsRequest struct {
	WorkspaceID string // 必填：项目 ID
	BoardID     string // 必填：看板 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBoardColumnsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"board_id":     r.BoardID,
	}
}
