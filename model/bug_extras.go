// Package model 中的 bug_extras.go 定义了缺陷关联、模板及自定义字段配置接口的请求/响应类型
package model

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
