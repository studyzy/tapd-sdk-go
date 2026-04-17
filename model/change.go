// Package model 中的 change.go 定义了 TAPD 变更历史数据模型及请求参数结构体
package model

// WorkitemChange 表示需求/任务的变更记录，适用于 Story 和 Task 的变更历史
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_changes.html
type WorkitemChange struct {
	ID               string `json:"id,omitempty"`
	WorkspaceID      string `json:"workspace_id,omitempty"`
	AppID            string `json:"app_id,omitempty"`
	WorkitemTypeID   string `json:"workitem_type_id,omitempty"`
	Creator          string `json:"creator,omitempty"`
	Created          string `json:"created,omitempty"`
	ChangeSummary    string `json:"change_summary,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Changes          string `json:"changes,omitempty"`
	EntityType       string `json:"entity_type,omitempty"`
	ChangeType       string `json:"change_type,omitempty"`
	ChangeTypeDetail string `json:"change_type_detail,omitempty"`
	ChangeTypeText   string `json:"change_type_text,omitempty"`
	Updated          string `json:"updated,omitempty"`
	StoryID          string `json:"story_id,omitempty"`
	TaskID           string `json:"task_id,omitempty"`
	MiniItemID       string `json:"mini_item_id,omitempty"`
	FieldChanges     string `json:"field_changes,omitempty"`
}

// BugChange 表示缺陷的变更记录
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_bug_changes.html
type BugChange struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	BugID       string `json:"bug_id,omitempty"`
	Author      string `json:"author,omitempty"`
	Field       string `json:"field,omitempty"`
	OldValue    string `json:"old_value,omitempty"`
	NewValue    string `json:"new_value,omitempty"`
	Memo        string `json:"memo,omitempty"`
	Created     string `json:"created,omitempty"`
}

// IterationChange 表示迭代的变更记录
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/get_iteration_changes.html
type IterationChange struct {
	ID           string `json:"id,omitempty"`
	WorkspaceID  string `json:"workspace_id,omitempty"`
	IterationID  string `json:"iteration_id,omitempty"`
	Author       string `json:"author,omitempty"`
	Field        string `json:"field,omitempty"`
	OldValue     string `json:"old_value,omitempty"`
	NewValue     string `json:"new_value,omitempty"`
	Memo         string `json:"memo,omitempty"`
	Created      string `json:"created,omitempty"`
	OperaterType string `json:"operater_type,omitempty"`
}

// GetStoryChangesRequest 查询需求变更历史的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_changes.html
type GetStoryChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 可选：需求 ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
	ChangeType  string // 可选：变更类型
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段列表
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetStoryChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "story_id", r.StoryID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "change_type", r.ChangeType)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountStoryChangesRequest 查询需求变更数量的请求参数
type CountStoryChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 可选：需求 ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountStoryChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "story_id", r.StoryID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	return params
}

// GetBugChangesRequest 查询缺陷变更历史的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_bug_changes.html
type GetBugChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	BugID       string // 可选：缺陷 ID
	Author      string // 可选：变更人
	Created     string // 可选：创建时间
	Field       string // 可选：变更字段
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段列表
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBugChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "bug_id", r.BugID)
	setOptional(params, "author", r.Author)
	setOptional(params, "created", r.Created)
	setOptional(params, "field", r.Field)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountBugChangesRequest 查询缺陷变更数量的请求参数
type CountBugChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	BugID       string // 可选：缺陷 ID
	Author      string // 可选：变更人
	Created     string // 可选：创建时间
	Field       string // 可选：变更字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountBugChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "bug_id", r.BugID)
	setOptional(params, "author", r.Author)
	setOptional(params, "created", r.Created)
	setOptional(params, "field", r.Field)
	return params
}

// GetTaskChangesRequest 查询任务变更历史的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/task/get_task_changes.html
type GetTaskChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	TaskID      string // 可选：任务 ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段列表
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetTaskChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "task_id", r.TaskID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountTaskChangesRequest 查询任务变更数量的请求参数
type CountTaskChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	TaskID      string // 可选：任务 ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountTaskChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "task_id", r.TaskID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	return params
}

// GetIterationChangesRequest 查询迭代变更历史的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/get_iteration_changes.html
type GetIterationChangesRequest struct {
	WorkspaceID string // 必填：项目 ID
	IterationID string // 必填：迭代 ID
	Author      string // 可选：变更人
	Created     string // 可选：创建时间
	Field       string // 可选：变更字段
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
	Fields      string // 可选：返回字段列表
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetIterationChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"iteration_id": r.IterationID,
	}
	setOptional(params, "author", r.Author)
	setOptional(params, "created", r.Created)
	setOptional(params, "field", r.Field)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "fields", r.Fields)
	return params
}
