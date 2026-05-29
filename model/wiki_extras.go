// Package model 中的 wiki_extras.go 定义了 Wiki 附属资源（附件、drawio、关注人、标签、可访问范围）的模型与请求参数
package model

// WikiDrawio 表示 Wiki 内嵌 drawio 数据
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_drawios.html
type WikiDrawio struct {
	ID     string `json:"id,omitempty"`
	Values string `json:"values,omitempty"` // drawio 的 xml 数据
}

// WikiEntityPermission 表示 Wiki 可访问范围（用户或用户组）
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_entity_permissions.html
type WikiEntityPermission struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	EntryType   string `json:"entry_type,omitempty"`  // 固定值 wiki
	TargetType  string `json:"target_type,omitempty"` // role_id 表示用户组，nick 表示用户昵称
	TargetID    string `json:"target_id,omitempty"`   // 用户昵称或用户组 ID
	WikiID      string `json:"wiki_id,omitempty"`
}

// WikiFollower 表示 Wiki 关注人
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_followers.html
type WikiFollower struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Created     string `json:"created,omitempty"`
	User        string `json:"user,omitempty"`
	WikiID      string `json:"wiki_id,omitempty"`
}

// WikiTag 表示 Wiki 标签信息
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_tags.html
type WikiTag struct {
	Tag     string `json:"tag,omitempty"`
	WikiID  string `json:"wiki_id,omitempty"`
	Creator string `json:"creator,omitempty"`
	Created string `json:"created,omitempty"`
}

// CountWikiAttachmentsRequest 获取 Wiki 附件数量的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_attachments_count.html
type CountWikiAttachmentsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：附件 id
	Filename    string // 可选：文件名
	Size        string // 可选：文件大小，字节
	Owner       string // 可选：上传者
	Created     string // 可选：创建时间，支持时间查询
	Modified    string // 可选：最后修改时间，支持时间查询
	WikiID      string // 可选：关联的 wiki id
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountWikiAttachmentsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "filename", r.Filename)
	setOptional(params, "size", r.Size)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "wiki_id", r.WikiID)
	return params
}

// GetWikiDrawioRequest 获取 Wiki drawio 数据的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_drawios.html
type GetWikiDrawioRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：drawio 数据 id（在 wiki 内容里）
	Token       string // 可选：验证用 token（如果 wiki 内容里有 token 必须传）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetWikiDrawioRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "token", r.Token)
	return params
}

// ListWikiEntityPermissionsRequest 获取 Wiki 可访问范围的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_entity_permissions.html
type ListWikiEntityPermissionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	WikiID      string // 必填：wiki ID
	TargetType  string // 可选：可访问的类型（role_id 表示用户组，user_id 表示用户）
	TargetID    string // 可选：用户 ID 或用户组 ID
	Limit       int    // 可选：返回数量限制，默认 30，最大 200
	Page        int    // 可选：页码，默认 1
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段列表
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListWikiEntityPermissionsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"wiki_id":      r.WikiID,
	}
	setOptional(params, "target_type", r.TargetType)
	setOptional(params, "target_id", r.TargetID)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// ListWikiFollowersRequest 获取 Wiki 关注人的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_followers.html
type ListWikiFollowersRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：记录 id
	WikiID      string // 可选：关联的 wiki id
	User        string // 可选：关注者昵称
	Created     string // 可选：创建时间，支持时间查询
	Limit       int    // 可选：返回数量限制，默认 30，最大 200
	Page        int    // 可选：页码，默认 1
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段列表
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListWikiFollowersRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "wiki_id", r.WikiID)
	setOptional(params, "user", r.User)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountWikiFollowersRequest 获取 Wiki 关注人数量的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_followers_count.html
type CountWikiFollowersRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：记录 id
	WikiID      string // 可选：关联的 wiki id
	User        string // 可选：关注者昵称
	Created     string // 可选：创建时间，支持时间查询
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountWikiFollowersRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "wiki_id", r.WikiID)
	setOptional(params, "user", r.User)
	setOptional(params, "created", r.Created)
	return params
}

// ListWikiTagsRequest 获取 Wiki 标签信息的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_tags.html
type ListWikiTagsRequest struct {
	WorkspaceID string // 必填：项目 ID
	WikiID      string // 可选：wiki id（不传则取项目下所有 wiki）
	Tag         string // 可选：标签
	Creator     string // 可选：标签创建人 nick
	Created     string // 可选：标签创建时间，支持时间查询
	Limit       int    // 可选：返回数量限制，默认 30，最大 200
	Page        int    // 可选：页码，默认 1
	Order       string // 可选：排序规则
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListWikiTagsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "wiki_id", r.WikiID)
	setOptional(params, "tag", r.Tag)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	return params
}

// CountWikiTagsRequest 获取 Wiki 标签信息数量的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/wiki/get_tapd_wikis_tags_count.html
type CountWikiTagsRequest struct {
	WorkspaceID string // 必填：项目 ID
	WikiID      string // 可选：wiki id（不传则取项目下所有 wiki）
	Tag         string // 可选：标签
	Creator     string // 可选：标签创建人 nick
	Created     string // 可选：标签创建时间，支持时间查询
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountWikiTagsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "wiki_id", r.WikiID)
	setOptional(params, "tag", r.Tag)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	return params
}
