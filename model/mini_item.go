// Package model 中的 mini_item.go 定义了 TAPD Mini 工作项数据模型
package model

import "encoding/json"

// MiniItem 表示 TAPD Mini 工作项
// 自定义字段（custom_field_*）通过 CustomFields map 保留，不会丢失
// 参考：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_items.html
type MiniItem struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	WorkspaceID         string `json:"workspace_id,omitempty"`
	CategoryID          string `json:"category_id,omitempty"`
	Status              string `json:"status,omitempty"`
	Owner               string `json:"owner,omitempty"`
	Begin               string `json:"begin,omitempty"`
	Due                 string `json:"due,omitempty"`
	Priority            string `json:"priority,omitempty"`
	Label               string `json:"label,omitempty"`
	DescriptionType     string `json:"description_type,omitempty"`
	Description         string `json:"description,omitempty"`
	MarkdownDescription string `json:"markdown_description,omitempty"`
	AncestorID          string `json:"ancestor_id,omitempty"`
	ParentID            string `json:"parent_id,omitempty"`
	ChildrenID          string `json:"children_id,omitempty"`
	Level               string `json:"level,omitempty"`
	Creator             string `json:"creator,omitempty"`
	Created             string `json:"created,omitempty"`
	Modifier            string `json:"modifier,omitempty"`
	Modified            string `json:"modified,omitempty"`
	Completed           string `json:"completed,omitempty"`
	HasAttachment       string `json:"has_attachment,omitempty"`
	Sort                string `json:"sort,omitempty"`
	IsArchived          string `json:"is_archived,omitempty"`
	ProgressManual      string `json:"progress_manual,omitempty"`
	Participator        string `json:"participator,omitempty"`

	// 自定义字段，key 为 custom_field_one、custom_field_9 等
	CustomFields map[string]string `json:"-"`
}

// UnmarshalJSON 自定义反序列化，在解析标准字段的同时收集 custom_field_* 字段
func (m *MiniItem) UnmarshalJSON(data []byte) error {
	type Alias MiniItem
	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*m = MiniItem(alias)

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	m.CustomFields = ExtractCustomFields(raw)
	return nil
}

// MarshalJSON 自定义序列化，将 CustomFields 中的键值对合并到输出 JSON
func (m MiniItem) MarshalJSON() ([]byte, error) {
	type Alias MiniItem
	b, err := json.Marshal(Alias(m))
	if err != nil {
		return nil, err
	}
	if len(m.CustomFields) == 0 {
		return b, nil
	}

	var base map[string]json.RawMessage
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}
	for k, v := range m.CustomFields {
		raw, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		base[k] = raw
	}
	return json.Marshal(base)
}

// MiniItemRelation 表示工作项与其他业务对象的关联关系
type MiniItemRelation struct {
	Type              string `json:"type,omitempty"`
	ID                string `json:"id,omitempty"`
	WorkspaceID       string `json:"workspace_id,omitempty"`
	LinkedWorkspaceID int    `json:"linked_workspace_id,omitempty"`
	Actas             string `json:"actas,omitempty"`
	Created           string `json:"created,omitempty"`
	Modified          string `json:"modified,omitempty"`
}

// RemovedMiniItem 表示回收站中的工作项
type RemovedMiniItem struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Creator       string `json:"creator,omitempty"`
	Created       string `json:"created,omitempty"`
	OperationUser string `json:"operation_user,omitempty"`
	Deleted       string `json:"deleted,omitempty"`
}

// CreateMiniItemRequest 创建工作项的请求参数
type CreateMiniItemRequest struct {
	WorkspaceID  string            // 必填：空间ID
	Name         string            // 必填：标题
	Priority     string            // 可选：优先级
	Owner        string            // 可选：处理人
	Creator      string            // 可选：创建人
	IsArchived   string            // 可选：是否归档
	Begin        string            // 可选：预计开始
	Due          string            // 可选：预计结束
	ParentID     string            // 可选：父工作项ID
	CategoryID   string            // 可选：分组
	Description  string            // 可选：详细描述
	Label        string            // 可选：标签，多个用|分隔
	CustomFields map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateMiniItemRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "priority", r.Priority)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "is_archived", r.IsArchived)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "parent_id", r.ParentID)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "description", r.Description)
	setOptional(params, "label", r.Label)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// UpdateMiniItemRequest 更新工作项的请求参数
type UpdateMiniItemRequest struct {
	ID             string            // 必填：工作项ID
	WorkspaceID    string            // 必填：空间ID
	Name           string            // 可选：标题
	Priority       string            // 可选：优先级
	Status         string            // 可选：状态
	ProgressManual string            // 可选：进度
	Owner          string            // 可选：处理人
	Begin          string            // 可选：预计开始
	Due            string            // 可选：预计结束
	CategoryID     string            // 可选：分组ID
	Description    string            // 可选：详细描述
	Label          string            // 可选：标签，多个用|分隔
	CustomFields   map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateMiniItemRequest) ToParams() map[string]string {
	params := map[string]string{
		"id":           r.ID,
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "status", r.Status)
	setOptional(params, "progress_manual", r.ProgressManual)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "description", r.Description)
	setOptional(params, "label", r.Label)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// ListMiniItemsRequest 查询工作项列表的请求参数
type ListMiniItemsRequest struct {
	WorkspaceID    string            // 必填：空间ID
	ID             string            // 可选：工作项ID，支持多ID查询
	Name           string            // 可选：标题，支持模糊匹配
	Priority       string            // 可选：优先级，支持枚举查询
	Status         string            // 可选：状态，支持枚举查询
	Label          string            // 可选：标签，支持枚举查询
	Owner          string            // 可选：处理人，支持模糊匹配
	IsArchived     string            // 可选：是否归档
	Creator        string            // 可选：创建人，支持多人查询
	Begin          string            // 可选：预计开始，支持时间范围查询
	Due            string            // 可选：预计结束，支持时间范围查询
	Created        string            // 可选：创建时间，支持时间范围查询
	Modified       string            // 可选：最后修改时间，支持时间范围查询
	Completed      string            // 可选：完成时间，支持时间范围查询
	ProgressManual string            // 可选：进度
	CategoryID     string            // 可选：分组，支持枚举查询
	ParentID       string            // 可选：父工作项
	ChildrenID     string            // 可选：子工作项
	Description    string            // 可选：详细描述，支持模糊匹配
	Limit          int               // 可选：每页数量，默认30，最大200
	Page           int               // 可选：页码，默认1
	Order          string            // 可选：排序规则
	Fields         string            // 可选：返回字段，逗号分隔
	CustomFields   map[string]string // 可选：自定义字段过滤，key 为 custom_field_* 或 cus_别名
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListMiniItemsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "status", r.Status)
	setOptional(params, "label", r.Label)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "is_archived", r.IsArchived)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "progress_manual", r.ProgressManual)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "parent_id", r.ParentID)
	setOptional(params, "children_id", r.ChildrenID)
	setOptional(params, "description", r.Description)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// CountMiniItemsRequest 查询工作项数量的请求参数
type CountMiniItemsRequest struct {
	WorkspaceID    string            // 必填：空间ID
	ID             string            // 可选：工作项ID，支持多ID查询
	Name           string            // 可选：标题，支持模糊匹配
	Priority       string            // 可选：优先级，支持枚举查询
	Status         string            // 可选：状态，支持枚举查询
	Label          string            // 可选：标签，支持枚举查询
	Owner          string            // 可选：处理人，支持模糊匹配
	Creator        string            // 可选：创建人，支持模糊匹配
	Begin          string            // 可选：预计开始，支持时间范围查询
	Due            string            // 可选：预计结束，支持时间范围查询
	Created        string            // 可选：创建时间，支持时间范围查询
	Modified       string            // 可选：最后修改时间，支持时间范围查询
	Completed      string            // 可选：完成时间，支持时间范围查询
	ProgressManual string            // 可选：进度
	CategoryID     string            // 可选：分组，支持枚举查询
	ParentID       string            // 可选：父工作项
	ChildrenID     string            // 可选：子工作项
	Description    string            // 可选：详细描述，支持模糊匹配
	IsArchived     string            // 可选：是否归档
	CustomFields   map[string]string // 可选：自定义字段过滤，key 为 custom_field_* 或 cus_别名
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountMiniItemsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "status", r.Status)
	setOptional(params, "label", r.Label)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "progress_manual", r.ProgressManual)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "parent_id", r.ParentID)
	setOptional(params, "children_id", r.ChildrenID)
	setOptional(params, "description", r.Description)
	setOptional(params, "is_archived", r.IsArchived)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// CreateMiniItemCategoryRequest 创建工作项分组的请求参数
type CreateMiniItemCategoryRequest struct {
	WorkspaceID string // 必填：空间ID
	Name        string // 必填：分组名称
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateMiniItemCategoryRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
}

// UpdateMiniItemCategoryRequest 更新工作项分组的请求参数
type UpdateMiniItemCategoryRequest struct {
	WorkspaceID string // 必填：空间ID
	ID          string // 必填：分组ID
	Name        string // 可选：分组名称
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateMiniItemCategoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	return params
}

// ListMiniItemCategoriesRequest 查询工作项分组列表的请求参数
type ListMiniItemCategoriesRequest struct {
	WorkspaceID string // 必填：空间ID
	ID          string // 可选：分组ID，支持多ID查询
	Name        string // 可选：分组名称，支持模糊匹配
	Created     string // 可选：创建时间，支持时间查询
	Modified    string // 可选：最后修改时间，支持时间查询
	Limit       int    // 可选：每页数量，默认30，最大200
	Page        int    // 可选：页码，默认1
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段，逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListMiniItemCategoriesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountMiniItemCategoriesRequest 查询工作项分组数量的请求参数
type CountMiniItemCategoriesRequest struct {
	WorkspaceID string // 必填：空间ID
	ID          string // 可选：分组ID，支持多ID查询
	Name        string // 可选：分组名称，支持模糊匹配
	Created     string // 可选：创建时间，支持时间查询
	Modified    string // 可选：最后修改时间，支持时间查询
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountMiniItemCategoriesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	return params
}

// GetMiniItemChangesRequest 查询工作项动态的请求参数
// created 和 mini_item_id 至少提供一个
type GetMiniItemChangesRequest struct {
	WorkspaceID      string // 必填：空间ID
	ID               string // 可选：动态ID，支持多ID查询
	MiniItemID       string // 可选：工作项ID，支持多ID查询
	Creator          string // 可选：创建人/操作人
	Created          string // 可选：创建时间/变更时间，支持时间查询
	ChangeType       string // 可选：变更类型：api, manual_update
	ChangeSummary    string // 可选：变更描述
	Comment          string // 可选：评论
	ChangeField      string // 可选：指定变更字段
	NeedParseChanges string // 可选：是否返回field_changes，默认1
	Limit            int    // 可选：每页数量，默认30，最大100
	Page             int    // 可选：页码，默认1
	Order            string // 可选：排序规则
	Fields           string // 可选：返回字段，逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetMiniItemChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "mini_item_id", r.MiniItemID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "change_type", r.ChangeType)
	setOptional(params, "change_summary", r.ChangeSummary)
	setOptional(params, "comment", r.Comment)
	setOptional(params, "change_field", r.ChangeField)
	setOptional(params, "need_parse_changes", r.NeedParseChanges)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountMiniItemChangesRequest 查询工作项动态数量的请求参数
type CountMiniItemChangesRequest struct {
	WorkspaceID   string // 必填：空间ID
	ID            string // 可选：动态ID
	MiniItemID    string // 可选：工作项ID
	Creator       string // 可选：创建人/操作人
	Created       string // 可选：创建时间/变更时间，支持时间查询
	ChangeSummary string // 可选：变更描述
	Comment       string // 可选：评论
	Changes       string // 可选：变更详细记录
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountMiniItemChangesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "mini_item_id", r.MiniItemID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "change_summary", r.ChangeSummary)
	setOptional(params, "comment", r.Comment)
	setOptional(params, "changes", r.Changes)
	return params
}

// CreateMiniItemRelationRequest 添加工作项与其他业务对象关联关系的请求参数
type CreateMiniItemRelationRequest struct {
	WorkspaceID string // 必填：空间ID
	TargetType  string // 必填：目标对象类型：story, bug, mini_item
	SourceID    string // 必填：源对象ID
	TargetID    string // 必填：目标对象ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateMiniItemRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"target_type":  r.TargetType,
		"source_id":    r.SourceID,
		"target_id":    r.TargetID,
	}
}

// RemoveMiniItemRelationRequest 解除工作项与其他业务对象关联关系的请求参数
type RemoveMiniItemRelationRequest struct {
	WorkspaceID string // 必填：空间ID
	MiniItemID  string // 必填：工作项ID
	TargetType  string // 必填：业务对象类型
	TargetID    string // 必填：业务对象ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *RemoveMiniItemRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"mini_item_id": r.MiniItemID,
		"target_type":  r.TargetType,
		"target_id":    r.TargetID,
	}
}

// GetMiniItemLinkedStoriesRequest 获取工作项关联需求的请求参数
type GetMiniItemLinkedStoriesRequest struct {
	WorkspaceID string // 必填：空间ID
	MiniItemID  string // 必填：工作项ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetMiniItemLinkedStoriesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"mini_item_id": r.MiniItemID,
	}
}

// GetMiniItemRelatedBugsRequest 获取工作项关联缺陷的请求参数
type GetMiniItemRelatedBugsRequest struct {
	WorkspaceID string // 必填：空间ID
	MiniItemID  string // 必填：工作项ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetMiniItemRelatedBugsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"mini_item_id": r.MiniItemID,
	}
}

// GetRemovedMiniItemsRequest 获取回收站内工作项的请求参数
type GetRemovedMiniItemsRequest struct {
	WorkspaceID string // 必填：空间ID
	ID          string // 可选：工作项ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
	Deleted     string // 可选：删除时间
	Limit       int    // 可选：每页数量，默认30
	Page        int    // 可选：页码，默认1
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetRemovedMiniItemsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "deleted", r.Deleted)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}
