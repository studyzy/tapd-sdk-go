// Package model 中的 board.go 定义了 TAPD 看板数据模型及请求参数结构体
package model

// BoardCard 表示 TAPD 看板工作项
type BoardCard struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	BoardID     string `json:"b_board_id,omitempty"`
	ColumnID    string `json:"b_column_id,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       string `json:"owner,omitempty"`
	CC          string `json:"cc,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Status      string `json:"status,omitempty"`
	Begin       string `json:"begin,omitempty"`
	Due         string `json:"due,omitempty"`
	BLabel      string `json:"b_label,omitempty"`
	BSort       string `json:"b_sort,omitempty"`
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
	Status      string `json:"status,omitempty"`
	Sort        string `json:"sort,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
}

// CreateBoardCardRequest 新建看板工作项的请求参数
type CreateBoardCardRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：工作项标题
	BoardID     string // 必填：看板 ID（API 字段 b_board_id）
	ColumnID    string // 必填：板块 ID（API 字段 b_column_id）
	Owner       string // 可选：负责人
	CC          string // 可选：参与人
	Status      string // 可选：状态
	Begin       string // 可选：开始时间
	Due         string // 可选：截止时间
	BLabel      string // 可选：标签 ID
	Description string // 可选：详细描述
	Priority    string // 可选：优先级（文档未列出，保留兼容）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateBoardCardRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
		"b_board_id":   r.BoardID,
		"b_column_id":  r.ColumnID,
	}
	setOptional(params, "owner", r.Owner)
	setOptional(params, "cc", r.CC)
	setOptional(params, "status", r.Status)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "b_label", r.BLabel)
	setOptional(params, "description", r.Description)
	setOptional(params, "priority", r.Priority)
	return params
}

// GetBoardCardsRequest 获取看板工作项的请求参数
type GetBoardCardsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：工作项 ID（支持多 ID）
	BoardID     string // 可选：看板 ID（API 字段 b_board_id）
	ColumnID    string // 可选：板块 ID（API 字段 b_column_id）
	Owner       string // 可选：负责人
	CC          string // 可选：参与人
	Status      string // 可选：状态
	Name        string // 可选：工作项标题
	Created     string // 可选：创建时间
	Begin       string // 可选：开始时间
	Due         string // 可选：截止时间
	BLabel      string // 可选：标签 ID
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBoardCardsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "b_board_id", r.BoardID)
	setOptional(params, "b_column_id", r.ColumnID)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "cc", r.CC)
	setOptional(params, "status", r.Status)
	setOptional(params, "name", r.Name)
	setOptional(params, "created", r.Created)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "b_label", r.BLabel)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "fields", r.Fields)
	return params
}

// UpdateBoardCardRequest 更新看板工作项的请求参数
type UpdateBoardCardRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：工作项 ID
	Name        string // 可选：工作项标题
	Owner       string // 可选：负责人
	CC          string // 可选：参与人
	Status      string // 可选：状态
	Begin       string // 可选：开始时间
	Due         string // 可选：截止时间
	BLabel      string // 可选：标签 ID
	Description string // 可选：详细描述
	ColumnID    string // 可选：板块 ID（文档未列出，保留兼容用于移动卡片）
	Priority    string // 可选：优先级（文档未列出，保留兼容）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateBoardCardRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "cc", r.CC)
	setOptional(params, "status", r.Status)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "b_label", r.BLabel)
	setOptional(params, "description", r.Description)
	setOptional(params, "b_column_id", r.ColumnID)
	setOptional(params, "priority", r.Priority)
	return params
}

// GetBoardColumnsRequest 获取看板板块的请求参数
type GetBoardColumnsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：板块 ID
	Name        string // 可选：板块名称
	BoardID     string // 可选：看板 ID
	Status      string // 可选：状态
	Created     string // 可选：创建时间
	Creator     string // 可选：创建人
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBoardColumnsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "board_id", r.BoardID)
	setOptional(params, "status", r.Status)
	setOptional(params, "created", r.Created)
	setOptional(params, "creator", r.Creator)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}
