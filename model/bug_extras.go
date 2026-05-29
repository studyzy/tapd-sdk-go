// Package model 中的 bug_extras.go 定义了缺陷关联、模板及自定义字段配置接口的请求/响应类型
package model

import "encoding/json"

// LinkBugsRequest 关联缺陷的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/link_bugs.html
type LinkBugsRequest struct {
	WorkspaceID string // 必填：项目 ID
	BugID       string // 必填：原始缺陷 ID
	RelateBugs  string // 必填：关联缺陷 ID，多个以英文逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *LinkBugsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"bug_id":       r.BugID,
		"relate_bugs":  r.RelateBugs,
	}
}

// DeleteLinkBugsRequest 取消关联缺陷的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/delete_link_bugs.html
type DeleteLinkBugsRequest struct {
	WorkspaceID string // 必填：项目 ID
	BugID       string // 必填：缺陷 ID
	LinkIDs     string // 必填：link_id，多个以英文逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *DeleteLinkBugsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"bug_id":       r.BugID,
		"link_ids":     r.LinkIDs,
	}
}

// GetLinkBugsRequest 获取缺陷关联缺陷的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_link_bugs.html
type GetLinkBugsRequest struct {
	WorkspaceID string // 必填：项目 ID
	BugID       string // 必填：缺陷 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetLinkBugsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"bug_id":       r.BugID,
	}
}

// BugLink 表示缺陷间的关联关系
type BugLink struct {
	Type              string `json:"type,omitempty"` // sync_copy/copy/repeat/direct_relate/sync_relate
	ID                string `json:"id,omitempty"`
	WorkspaceID       string `json:"workspace_id,omitempty"`
	Actas             string `json:"actas,omitempty"` // target 为操作发起方
	LinkedWorkspaceID int    `json:"linked_workspace_id,omitempty"`
	LinkID            string `json:"link_id,omitempty"`
}

// GetBugRelatedStoriesRequest 获取缺陷关联需求 ID 的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_related_stories.html
type GetBugRelatedStoriesRequest struct {
	WorkspaceID string // 必填：项目 ID
	BugID       string // 必填：缺陷 ID，支持多 ID 查询
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBugRelatedStoriesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"bug_id":       r.BugID,
	}
}

// BugStoryRelation 表示缺陷关联的需求
type BugStoryRelation struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
	BugID       string `json:"bug_id,omitempty"`
	StoryID     string `json:"story_id,omitempty"`
}

// GetDefaultBugTemplateRequest 获取缺陷模板字段的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_default_bug_template.html
type GetDefaultBugTemplateRequest struct {
	WorkspaceID      string // 必填：项目 ID
	TemplateID       string // 必填：模板 ID
	UsePriorityLabel string // 可选：是否替换优先级字段为 priority_label，取值 0/1
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetDefaultBugTemplateRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"template_id":  r.TemplateID,
	}
	setOptional(params, "use_priority_label", r.UsePriorityLabel)
	return params
}

// GetBugFieldsInfoRequest 获取缺陷所有字段及候选值的请求参数
type GetBugFieldsInfoRequest struct {
	WorkspaceID string // 必填：项目 ID
	AllOptions  string // 可选：是否返回所有候选值
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBugFieldsInfoRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "all_options", r.AllOptions)
	return params
}

// CopyBugRequest 复制缺陷的请求参数
type CopyBugRequest struct {
	WorkspaceID    string // 必填：源项目 ID
	SrcBugID       string // 必填：源缺陷 ID
	DstWorkspaceID string // 必填：目标项目 ID
	SyncFields     string // 可选：需要同步的字段（逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CopyBugRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id":     r.WorkspaceID,
		"src_bug_id":       r.SrcBugID,
		"dst_workspace_id": r.DstWorkspaceID,
	}
	setOptional(params, "sync_fields", r.SyncFields)
	return params
}

// BatchUpdateBugItem 批量更新中的单条缺陷字段
type BatchUpdateBugItem struct {
	ID            string            `json:"id,omitempty"`             // 必填：缺陷 ID
	Title         string            `json:"title,omitempty"`          // 可选：标题
	Description   string            `json:"description,omitempty"`    // 可选：详细描述
	Priority      string            `json:"priority,omitempty"`       // 可选：优先级
	PriorityLabel string            `json:"priority_label,omitempty"` // 可选：优先级（推荐）
	Severity      string            `json:"severity,omitempty"`       // 可选：严重程度
	Status        string            `json:"status,omitempty"`         // 可选：状态
	VStatus       string            `json:"v_status,omitempty"`       // 可选：中文状态名
	BugType       string            `json:"bugtype,omitempty"`        // 可选：缺陷类型
	CurrentOwner  string            `json:"current_owner,omitempty"`  // 可选：处理人
	CurrentUser   string            `json:"current_user,omitempty"`   // 可选：变更人
	IterationID   string            `json:"iteration_id,omitempty"`   // 可选：迭代 ID
	Module        string            `json:"module,omitempty"`         // 可选：模块
	Label         string            `json:"label,omitempty"`          // 可选：标签
	Begin         string            `json:"begin,omitempty"`          // 可选：预计开始
	Due           string            `json:"due,omitempty"`            // 可选：预计结束
	CustomFields  map[string]string `json:"-"`                        // 可选：自定义字段
}

// MarshalJSON 自定义序列化，将 CustomFields 合并到输出 JSON
func (i BatchUpdateBugItem) MarshalJSON() ([]byte, error) {
	type Alias BatchUpdateBugItem
	b, err := json.Marshal(Alias(i))
	if err != nil {
		return nil, err
	}
	if len(i.CustomFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range i.CustomFields {
		if v == "" {
			continue
		}
		raw, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		m[k] = raw
	}
	return json.Marshal(m)
}

// BatchUpdateBugRequest 批量更新缺陷的请求参数
type BatchUpdateBugRequest struct {
	WorkspaceID string               // 必填：项目 ID
	Workitems   []BatchUpdateBugItem // 必填：要更新的缺陷列表（单次最多 50 条）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *BatchUpdateBugRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	if len(r.Workitems) > 0 {
		raw, _ := json.Marshal(r.Workitems)
		params["workitems"] = string(raw)
	}
	return params
}

// RemovedBug 表示回收站中的缺陷
type RemovedBug struct {
	ID            string `json:"id,omitempty"`             // 缺陷 ID
	Name          string `json:"name,omitempty"`           // 标题
	Creator       string `json:"creator,omitempty"`        // 创建人
	Created       string `json:"created,omitempty"`        // 创建时间
	OperationUser string `json:"operation_user,omitempty"` // 删除人
	Modified      string `json:"modified,omitempty"`       // 最后修改时间
}

// GetRemovedBugsRequest 获取回收站缺陷的请求参数
type GetRemovedBugsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：缺陷 ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
	Modified    string // 可选：最后修改时间
	IncludeAll  string // 可选：是否包含全部
	Limit       int    // 可选：返回数量限制（默认 30，最大 200）
	Page        int    // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetRemovedBugsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "include_all", r.IncludeAll)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// GetBugsByViewConfIDRequest 获取视图对应的缺陷列表的请求参数
type GetBugsByViewConfIDRequest struct {
	WorkspaceID string // 必填：项目 ID
	ViewConfID  string // 必填：视图 ID
	CurrentUser string // 可选：当前登录用户
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBugsByViewConfIDRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"view_conf_id": r.ViewConfID,
	}
	setOptional(params, "current_user", r.CurrentUser)
	setOptional(params, "fields", r.Fields)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// BugIDsToQueryTokenRequest 缺陷 ID 转换成 QueryToken 的请求参数
type BugIDsToQueryTokenRequest struct {
	WorkspaceID string // 必填：项目 ID
	IDs         string // 必填：缺陷 ID 列表，逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *BugIDsToQueryTokenRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"ids":          r.IDs,
	}
}

// UpdateBugSystemSelectFieldOptionsRequest 更新缺陷系统下拉字段选项的请求参数
type UpdateBugSystemSelectFieldOptionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	Field       string // 必填：字段名
	Options     string // 必填：选项列表（JSON 格式）
	Value       string // 可选：默认值
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateBugSystemSelectFieldOptionsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"field":        r.Field,
		"options":      r.Options,
	}
	setOptional(params, "value", r.Value)
	return params
}
