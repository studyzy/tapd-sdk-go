// Package model 中的 request.go 定义了杂项 API 的请求参数结构体
package model

import "encoding/json"

// GetCustomFieldsRequest 获取自定义字段配置的请求参数
type GetCustomFieldsRequest struct {
	WorkspaceID string // 必填：项目 ID
	EntityType  string // 必填：实体类型（stories/tasks/iterations/tcases）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetCustomFieldsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
	}
}

// WorkspaceIDRequest 仅需项目 ID 的通用请求参数
// 适用于 GetStoryFieldsLabel、GetStoryFieldsInfo、GetWorkitemTypes、
// ListCategories 等仅需 workspace_id 的接口
type WorkspaceIDRequest struct {
	WorkspaceID string // 必填：项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *WorkspaceIDRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
	}
}

// WorkflowRequest 工作流相关接口的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/
type WorkflowRequest struct {
	WorkspaceID    string // 必填：项目 ID
	System         string // 可选：系统名（story/bug），用于 status_map/first_step/last_steps 等接口
	WorkitemTypeID string // 可选：需求类别 ID
	ID             string // 可选：工作流 ID
	Name           string // 可选：工作流名称
	Description    string // 可选：工作流描述
	SystemName     string // 可选：系统名称（用于 get_workflows 接口，参数名 system_name）
	Creator        string // 可选：创建人
	Created        string // 可选：创建时间
	Limit          string // 可选：返回数量限制
	Page           string // 可选：页码
	Type           string // 可选：工作流类型
	GroupKey       string // 可选：分组键（用于 all_last_steps 等接口）
	Fields         string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *WorkflowRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "system", r.System)
	setOptional(params, "workitem_type_id", r.WorkitemTypeID)
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "system_name", r.SystemName)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "limit", r.Limit)
	setOptional(params, "page", r.Page)
	setOptional(params, "type", r.Type)
	setOptional(params, "group_key", r.GroupKey)
	setOptional(params, "fields", r.Fields)
	return params
}

// AddNewStepForBugRequest 新增 Bug 工作流步骤的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/add_new_step_for_bug.html
type AddNewStepForBugRequest struct {
	WorkspaceID string // 必填：项目 ID
	WorkflowID  string // 必填：工作流 ID
	StepName    string // 必填：步骤名称
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *AddNewStepForBugRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"workflow_id":  r.WorkflowID,
		"step_name":    r.StepName,
	}
}

// GetCommitMsgRequest 获取源码提交关键字的请求参数
type GetCommitMsgRequest struct {
	WorkspaceID string // 必填：项目 ID
	ObjectID    string // 必填：条目 ID
	Type        string // 必填：条目类型（story/task/bug）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetCommitMsgRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"object_id":    r.ObjectID,
		"type":         r.Type,
	}
}

// GetTodoRequest 获取用户待办事项的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/user/get_user_todo_story.html
type GetTodoRequest struct {
	WorkspaceID string // 必填：项目 ID
	EntityType  string // 必填：对象类型（story/bug/task）
	User        string // 可选：用户
	Limit       int    // 可选：返回数量限制（默认 30，最大 200）
	Page        int    // 可选：页码（默认 1）
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段列表（逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetTodoRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "user", r.User)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// GetRelatedBugsRequest 获取需求关联缺陷的请求参数
type GetRelatedBugsRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetRelatedBugsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
	}
}

// CreateRelationRequest 创建实体关联关系的请求参数
type CreateRelationRequest struct {
	WorkspaceID string // 必填：项目 ID
	SourceType  string // 必填：源实体类型（story/bug/task）
	TargetType  string // 必填：目标实体类型（story/bug/task）
	SourceID    string // 必填：源实体 ID
	TargetID    string // 必填：目标实体 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"source_type":  r.SourceType,
		"target_type":  r.TargetType,
		"source_id":    r.SourceID,
		"target_id":    r.TargetID,
	}
}

// GetSubWorkspacesRequest 获取子项目信息的请求参数
type GetSubWorkspacesRequest struct {
	WorkspaceID string // 必填：项目 ID
	TemplateID  string // 可选：模板 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetSubWorkspacesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "template_id", r.TemplateID)
	return params
}

// ListCompanyProjectsRequest 获取公司项目列表的请求参数
type ListCompanyProjectsRequest struct {
	CompanyID   string // 必填：公司 ID
	Category    string // 可选：项目类别
	WithExtends string // 可选：是否返回扩展信息
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListCompanyProjectsRequest) ToParams() map[string]string {
	params := map[string]string{
		"company_id": r.CompanyID,
	}
	setOptional(params, "category", r.Category)
	setOptional(params, "with_extends", r.WithExtends)
	return params
}

// GetWorkspaceUsersRequest 获取指定项目成员的请求参数
type GetWorkspaceUsersRequest struct {
	WorkspaceID string // 必填：项目 ID
	User        string // 可选：用户账号
	Fields      string // 可选：返回字段列表（逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetWorkspaceUsersRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "user", r.User)
	setOptional(params, "fields", r.Fields)
	return params
}

// AddWorkspaceMemberRequest 添加项目成员的请求参数
type AddWorkspaceMemberRequest struct {
	WorkspaceID string // 必填：项目 ID
	Nick        string // 必填：用户昵称
	CompanyID   string // 可选：公司 ID
	RoleIDs     string // 可选：用户组 ID（逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *AddWorkspaceMemberRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"nick":         r.Nick,
	}
	setOptional(params, "company_id", r.CompanyID)
	setOptional(params, "role_ids", r.RoleIDs)
	return params
}

// CreateMiniProjectRequest 新建空间的请求参数
type CreateMiniProjectRequest struct {
	CompanyID   string // 必填：公司ID
	Name        string // 必填：空间名称
	Creator     string // 必填：创建人（必须属于当前公司）
	Description string // 可选：空间描述
	TemplateID  string // 可选：模板ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateMiniProjectRequest) ToParams() map[string]string {
	params := map[string]string{
		"company_id": r.CompanyID,
		"name":       r.Name,
		"creator":    r.Creator,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "template_id", r.TemplateID)
	return params
}

// GetMiniProjectListRequest 获取用户所有参与空间的请求参数
type GetMiniProjectListRequest struct {
	User      string // 必填：用户名
	CompanyID string // 必填：公司ID
	CanEdit   string // 可选：是否过滤有编辑权限的空间
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetMiniProjectListRequest) ToParams() map[string]string {
	params := map[string]string{
		"user":       r.User,
		"company_id": r.CompanyID,
	}
	setOptional(params, "can_edit", r.CanEdit)
	return params
}

// MiniProject 表示 Mini 空间简要信息
type MiniProject struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Status  string `json:"status,omitempty"`
	Creator string `json:"creator,omitempty"`
	Created string `json:"created,omitempty"`
}

// CreateMiniProjectResponse 新建空间的响应
type CreateMiniProjectResponse struct {
	WorkspaceID  string `json:"workspace_id,omitempty"`
	WorkspaceURL string `json:"workspace_url,omitempty"`
}

// UploadAttachmentRequest 上传附件的请求参数（不含文件本身，文件通过 multipart 传输）
type UploadAttachmentRequest struct {
	WorkspaceID string // 必填：空间ID
	Type        string // 必填：固定值 story_custom_field
	CustomField string // 必填：自定义字段英文名
	EntryID     string // 必填：工作项ID
	Owner       string // 可选：附件创建者
}

// UploadImageBase64Request 上传 base64 图片的请求参数
type UploadImageBase64Request struct {
	WorkspaceID string // 必填：空间ID
	Base64Data  string // 必填：base64格式的图片数据
	Type        string // 必填：固定值 story_custom_field
	CustomField string // 必填：自定义字段英文名
	EntryID     string // 必填：工作项ID
	Owner       string // 可选：附件创建者
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UploadImageBase64Request) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"base64_data":  r.Base64Data,
		"type":         r.Type,
		"custom_field": r.CustomField,
		"entry_id":     r.EntryID,
	}
	setOptional(params, "owner", r.Owner)
	return params
}

// DownloadAttachmentRequest 获取单个附件下载链接的请求参数
type DownloadAttachmentRequest struct {
	WorkspaceID string // 必填：空间ID
	ID          string // 必填：附件ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *DownloadAttachmentRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
}

// CustomFieldConfig 表示 TAPD 自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_custom_fields_settings.html
type CustomFieldConfig struct {
	ID              string `json:"id,omitempty"`
	WorkspaceID     string `json:"workspace_id,omitempty"`
	AppID           string `json:"app_id,omitempty"`
	EntryType       string `json:"entry_type,omitempty"`        // 所属实体类型（story/bug/task 等）
	CustomField     string `json:"custom_field,omitempty"`      // 自定义字段标识（如 custom_field_17）
	Type            string `json:"type,omitempty"`              // 输入类型（text/select/cascade_radio 等）
	Name            string `json:"name,omitempty"`              // 显示名称
	Options         string `json:"options,omitempty"`           // 可选值（JSON 字符串，需二次解析）
	ExtraConfig     string `json:"extra_config,omitempty"`      // 额外配置
	Enabled         string `json:"enabled,omitempty"`           // 是否启用（1=启用，0=禁用）
	Creator         string `json:"creator,omitempty"`           // 创建人
	Created         string `json:"created,omitempty"`           // 创建时间
	Modified        string `json:"modified,omitempty"`          // 最后修改时间
	Freeze          string `json:"freeze,omitempty"`            // 是否冻结
	Sort            string `json:"sort,omitempty"`              // 显示排序
	Memo            string `json:"memo,omitempty"`              // 备注
	OpenExtensionID string `json:"open_extension_id,omitempty"` // 插件扩展字段标识
	IsOut           int    `json:"is_out,omitempty"`            // 已弃用
	IsUninstall     int    `json:"is_uninstall,omitempty"`      // 应用是否未安装（0=已安装，1=未安装）
	AppName         string `json:"app_name,omitempty"`          // 关联应用名称
	CreatedFrom     string `json:"created_from,omitempty"`      // 创建来源
}

// FieldInfo 表示需求字段的详细信息（含候选值）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_info.html
type FieldInfo struct {
	Name         string        `json:"name,omitempty"`          // 字段英文名
	Label        string        `json:"label,omitempty"`         // 字段中文名
	HTMLType     string        `json:"html_type,omitempty"`     // 控件类型（input/select/rich_edit 等）
	Options      interface{}   `json:"options,omitempty"`       // 候选值（object 或 array，因字段类型而异）
	PureOptions  interface{}   `json:"pure_options,omitempty"`  // 详细候选值（含层级元信息）
	Readonly     int           `json:"readonly,omitempty"`      // 是否只读（0=可编辑，1=只读）
	ColorOptions []ColorOption `json:"color_options,omitempty"` // 带颜色的候选值（如优先级）
	UserOptions  interface{}   `json:"user_options,omitempty"`  // 用户选择器的用户列表
	Memo         string        `json:"memo,omitempty"`          // 自定义字段备注
	EnableColor  int           `json:"enable_color,omitempty"`  // 是否启用颜色标记
}

// ColorOption 表示带颜色的候选值（如优先级字段）
type ColorOption struct {
	Value string `json:"value,omitempty"` // 选项值
	Color string `json:"color,omitempty"` // 颜色（十六进制色值）
	Label string `json:"label,omitempty"` // 显示名称
}

// WorkitemType 表示 TAPD 需求类别
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_workitem_types.html
type WorkitemType struct {
	ID             string `json:"id,omitempty"`
	WorkspaceID    string `json:"workspace_id,omitempty"`
	AppID          string `json:"app_id,omitempty"`
	EntityType     string `json:"entity_type,omitempty"`      // 类别别名
	Name           string `json:"name,omitempty"`             // 类别名称
	EnglishName    string `json:"english_name,omitempty"`     // 英文名称
	Status         string `json:"status,omitempty"`           // 状态（1=未完成，2=未启用，3=已启用）
	Color          string `json:"color,omitempty"`            // 颜色
	WorkflowID     string `json:"workflow_id,omitempty"`      // 关联工作流 ID
	ChildrenIDs    string `json:"children_ids,omitempty"`     // 允许的子需求类别
	ParentIDs      string `json:"parent_ids,omitempty"`       // 允许的父需求类别
	Icon           string `json:"icon,omitempty"`             // 图标路径
	IconSmall      string `json:"icon_small,omitempty"`       // 小图标路径
	Creator        string `json:"creator,omitempty"`          // 创建人
	Created        string `json:"created,omitempty"`          // 创建时间
	ModifiedBy     string `json:"modified_by,omitempty"`      // 最后修改人
	Modified       string `json:"modified,omitempty"`         // 最后修改时间
	IconViper      string `json:"icon_viper,omitempty"`       // 图标完整 URL
	IconSmallViper string `json:"icon_small_viper,omitempty"` // 小图标完整 URL
}

// StoryBugRelation 表示需求与缺陷的关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_related_bugs.html
type StoryBugRelation struct {
	WorkspaceID int    `json:"workspace_id,omitempty"` // 项目 ID
	StoryID     string `json:"story_id,omitempty"`     // 需求 ID
	BugID       string `json:"bug_id,omitempty"`       // 缺陷 ID
}

// GetLifeTimesRequest 获取状态流转时间的请求参数
type GetLifeTimesRequest struct {
	WorkspaceID string // 必填：项目 ID
	EntityType  string // 必填：对象类型（story/bug/task）
	EntityID    string // 必填：实体 ID
	Created     string // 可选：创建时间
	Limit       string // 可选：返回数量限制
	Page        string // 可选：页码
	Order       string // 可选：排序
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetLifeTimesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"entity_type":  r.EntityType,
		"entity_id":    r.EntityID,
	}
	setOptional(params, "created", r.Created)
	setOptional(params, "limit", r.Limit)
	setOptional(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// AddCodeCommitInfoRequest 保存 Commit 提交数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/add_code_commit_info.html
type AddCodeCommitInfoRequest struct {
	WorkspaceID string   // 必填：项目 ID
	Message     string   // 必填：提交信息
	Author      string   // 必填：提交人
	CommitID    string   // 必填：提交 ID
	Files       []string // 必填：变更文件
	Repo        string   // 必填：仓库名
	RepoID      string   // 必填：仓库 ID
	CommitTime  string   // 必填：提交时间
	HookURL     string   // 可选：Hook URL（保留兼容）
	Ref         string   // 可选：分支引用（保留兼容）
	GitEnv      string   // 可选：信息来源
	RepoURL     string   // 可选：仓库链接
	CommitURL   string   // 可选：提交链接
}

// ToJSON 将请求结构体序列化为 JSON 字节（该 API 使用 JSON Body 提交）
func (r *AddCodeCommitInfoRequest) ToJSON() ([]byte, error) {
	m := map[string]interface{}{
		"workspace_id": r.WorkspaceID,
		"message":      r.Message,
		"author":       r.Author,
		"commit_id":    r.CommitID,
		"repo":         r.Repo,
		"repo_id":      r.RepoID,
		"commit_time":  r.CommitTime,
	}
	if len(r.Files) > 0 {
		m["files"] = r.Files
	}
	if r.HookURL != "" {
		m["hook_url"] = r.HookURL
	}
	if r.Ref != "" {
		m["ref"] = r.Ref
	}
	if r.GitEnv != "" {
		m["git_env"] = r.GitEnv
	}
	if r.RepoURL != "" {
		m["repo_url"] = r.RepoURL
	}
	if r.CommitURL != "" {
		m["commit_url"] = r.CommitURL
	}
	return json.Marshal(m)
}

// GetCodeCommitInfosRequest 获取 GIT 关联提交数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/get_code_commit_infos.html
type GetCodeCommitInfosRequest struct {
	WorkspaceID string // 必填：项目 ID
	Type        string // 必填：TAPD 业务对象类型
	ObjectID    string // 必填：TAPD 业务对象 ID
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Created     string // 可选：创建时间（保留兼容）
	CommitTime  string // 可选：提交时间
	RelatedType string // 可选：关联类型
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetCodeCommitInfosRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"type":         r.Type,
		"object_id":    r.ObjectID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "created", r.Created)
	setOptional(params, "commit_time", r.CommitTime)
	setOptional(params, "related_type", r.RelatedType)
	return params
}

// GetOneAttachmentRequest 获取单个附件下载链接的请求参数
type GetOneAttachmentRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：附件 ID
	FileName    string // 可选：文件名
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetOneAttachmentRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "file_name", r.FileName)
	return params
}

// DownloadDocumentRequest 获取单个文档下载链接的请求参数
type DownloadDocumentRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：文档 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *DownloadDocumentRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
}

// Report 表示 TAPD 项目报告
type Report struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Created     string `json:"created,omitempty"`
}

// LifeTime 表示状态流转时间
type LifeTime struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	EntityType  string `json:"entity_type,omitempty"`
	EntityID    string `json:"entity_id,omitempty"`
	Status      string `json:"status,omitempty"`
	BeginDate   string `json:"begin_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	TimeCost    string `json:"time_cost,omitempty"`
	Owner       string `json:"owner,omitempty"`
	IsRepeated  string `json:"is_repeated,omitempty"`
	Created     string `json:"created,omitempty"`
	Operator    string `json:"operator,omitempty"`
	ChangeFrom  string `json:"change_from,omitempty"`
}
