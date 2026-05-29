// Package model 中的 setting.go 定义了 TAPD 设置相关数据模型（模块/版本/基线/特性）
package model

// Module 表示 TAPD 模块
type Module struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Owner       string `json:"owner,omitempty"`
}

// CreateModuleRequest 创建模块的请求参数
type CreateModuleRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：模块名称
	Description string // 可选：模块描述
	Owner       string // 可选：负责人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateModuleRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	return params
}

// UpdateModuleRequest 更新模块的请求参数
type UpdateModuleRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：模块 ID
	Name        string // 可选：模块名称
	Description string // 可选：模块描述
	Owner       string // 可选：负责人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateModuleRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	return params
}

// GetModulesRequest 获取模块列表的请求参数
type GetModulesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：模块 ID（支持多 ID 查询，逗号分隔）
	Name        string // 可选：标题（支持模糊匹配）
	Description string // 可选：详细描述
	Owner       string // 可选：负责人
	Created     string // 可选：创建时间（支持时间查询）
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段，多个以英文逗号隔开
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetModulesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountModulesRequest 获取模块数量的请求参数
type CountModulesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：模块 ID（支持多 ID 查询，逗号分隔）
	Name        string // 可选：标题（支持模糊匹配）
	Description string // 可选：详细描述
	Owner       string // 可选：负责人
	Created     string // 可选：创建时间（支持时间查询）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountModulesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "created", r.Created)
	return params
}

// Version 表示 TAPD 版本
type Version struct {
	ID             string `json:"id,omitempty"`
	WorkspaceID    string `json:"workspace_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Creator        string `json:"creator,omitempty"`
	Created        string `json:"created,omitempty"`
	Modified       string `json:"modified,omitempty"`
	Due            string `json:"due,omitempty"`
	Completed      string `json:"completed,omitempty"`
	Default        string `json:"default,omitempty"`
	ParentID       string `json:"parent_id,omitempty"`
	Path           string `json:"path,omitempty"`
	ModuleID       string `json:"module_id,omitempty"`
	Start          string `json:"start,omitempty"`
	RealEnd        string `json:"realend,omitempty"`
	TestTime       string `json:"testtime,omitempty"`
	RealBegin      string `json:"realbegin,omitempty"`
	ReleaseTime    string `json:"releasetime,omitempty"`
	BusinessModule string `json:"business_module,omitempty"`
	VersionType    string `json:"version_type,omitempty"`
	Modifier       string `json:"modifier,omitempty"`
	ModifiedTime   string `json:"modified_time,omitempty"`
	Status         string `json:"status,omitempty"`
	Owner          string `json:"owner,omitempty"`
}

// CreateVersionRequest 创建版本的请求参数
type CreateVersionRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：版本名称
	Creator     string // 必填：创建人
	ID          string // 可选：版本 ID
	Description string // 可选：版本描述
	Owner       string // 可选：负责人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateVersionRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
		"creator":      r.Creator,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	return params
}

// UpdateVersionRequest 更新版本的请求参数
type UpdateVersionRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：版本 ID
	Modifier    string // 必填：当前处理人
	Creator     string // 可选：提交人
	Name        string // 可选：版本名称
	Description string // 可选：版本描述
	Owner       string // 可选：负责人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateVersionRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
		"modifier":     r.Modifier,
	}
	setOptional(params, "creator", r.Creator)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	return params
}

// GetVersionsRequest 获取版本列表的请求参数
type GetVersionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：版本 ID（支持多 ID 查询，逗号分隔）
	Owner       string // 可选：负责人
	Creator     string // 可选：提交人
	Name        string // 可选：版本标题（支持模糊匹配）
	Created     string // 可选：创建时间（支持时间查询）
	Status      string // 可选：状态（Closed/Unclosed）
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段，多个以英文逗号隔开
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetVersionsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "name", r.Name)
	setOptional(params, "created", r.Created)
	setOptional(params, "status", r.Status)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountVersionsRequest 获取版本数量的请求参数
type CountVersionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：版本 ID（支持多 ID 查询，逗号分隔）
	Owner       string // 可选：负责人
	Name        string // 可选：版本标题（支持模糊匹配）
	Description string // 可选：详细描述
	Created     string // 可选：创建时间（支持时间查询）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountVersionsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "created", r.Created)
	return params
}

// Baseline 表示 TAPD 基线
type Baseline struct {
	ID           string `json:"id,omitempty"`
	WorkspaceID  string `json:"workspace_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Created      string `json:"created,omitempty"`
	Modified     string `json:"modified,omitempty"`
	Owner        string `json:"owner,omitempty"`
	Due          string `json:"due,omitempty"`
	Completed    string `json:"completed,omitempty"`
	Default      string `json:"default,omitempty"`
	VersionID    string `json:"version_id,omitempty"`
	ModuleID     string `json:"module_id,omitempty"`
	SvnTag       string `json:"svn_tag,omitempty"`
	SvnProjectID string `json:"svn_project_id,omitempty"`
	SvnPathID    string `json:"svn_path_id,omitempty"`
	SvnSyncType  string `json:"svn_sync_type,omitempty"`
}

// CreateBaselineRequest 创建基线的请求参数
type CreateBaselineRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 可选：基线名称（文档标注为可选）
	VersionID   string // 可选：关联版本 ID
	Description string // 可选：基线描述
	Owner       string // 可选：处理人
	Completed   string // 可选：状态（0 未完成 / 1 完成）
	Due         string // 可选：预计结束日期
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateBaselineRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "version_id", r.VersionID)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "due", r.Due)
	return params
}

// UpdateBaselineRequest 更新基线的请求参数
type UpdateBaselineRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：基线 ID
	Name        string // 可选：基线名称
	VersionID   string // 可选：关联版本 ID
	Description string // 可选：基线描述
	Owner       string // 可选：处理人
	Completed   string // 可选：状态（0 未完成 / 1 完成）
	Due         string // 可选：预计结束日期
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateBaselineRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "version_id", r.VersionID)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "due", r.Due)
	return params
}

// GetBaselinesRequest 获取基线列表的请求参数
type GetBaselinesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：基线 ID（支持多 ID 查询，逗号分隔）
	Name        string // 可选：基线名称（支持模糊匹配）
	VersionID   string // 可选：关联版本 ID
	Description string // 可选：详细描述
	Owner       string // 可选：处理人
	Completed   string // 可选：状态（0/1）
	Due         string // 可选：预计结束日期
	Created     string // 可选：创建时间（支持时间查询）
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段，多个以英文逗号隔开
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBaselinesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "version_id", r.VersionID)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "due", r.Due)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountBaselinesRequest 获取基线数量的请求参数
type CountBaselinesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：基线 ID（支持多 ID 查询，逗号分隔）
	Name        string // 可选：基线名称（支持模糊匹配）
	Description string // 可选：详细描述
	Owner       string // 可选：负责人
	Created     string // 可选：创建时间（支持时间查询）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountBaselinesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "created", r.Created)
	return params
}

// Feature 表示 TAPD 特性
type Feature struct {
	ID                string `json:"id,omitempty"`
	WorkspaceID       string `json:"workspace_id,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Creator           string `json:"creator,omitempty"`
	Created           string `json:"created,omitempty"`
	Modified          string `json:"modified,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Due               string `json:"due,omitempty"`
	Completed         string `json:"completed,omitempty"`
	Default           string `json:"default,omitempty"`
	ModuleID          string `json:"module_id,omitempty"`
	ReleaseID         string `json:"release_id,omitempty"`
	ReleaseName       string `json:"release_name,omitempty"`
	Priority          string `json:"priority,omitempty"`
	Estimate          string `json:"estimate,omitempty"`
	EstimateCompleted string `json:"estimate_completed,omitempty"`
}

// CreateFeatureRequest 创建特性的请求参数
type CreateFeatureRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：特性名称
	Description string // 可选：特性描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateFeatureRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	return params
}

// UpdateFeatureRequest 更新特性的请求参数
type UpdateFeatureRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：特性 ID
	Name        string // 可选：特性名称
	Description string // 可选：特性描述
	Owner       string // 可选：负责人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateFeatureRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	return params
}

// GetFeaturesRequest 获取特性列表的请求参数
type GetFeaturesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：特性 ID（支持多 ID 查询，逗号分隔）
	Name        string // 可选：标题（支持模糊匹配）
	Description string // 可选：详细描述
	Owner       string // 可选：负责人
	Created     string // 可选：创建时间（支持时间查询）
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段，多个以英文逗号隔开
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetFeaturesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "created", r.Created)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// CountFeaturesRequest 获取特性数量的请求参数
type CountFeaturesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：特性 ID（支持多 ID 查询，逗号分隔）
	Name        string // 可选：标题（支持模糊匹配）
	Description string // 可选：详细描述
	Owner       string // 可选：负责人
	Created     string // 可选：创建时间（支持时间查询）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountFeaturesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "created", r.Created)
	return params
}

// IterationLockRequest 锁定/解锁迭代的请求参数
type IterationLockRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：迭代 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *IterationLockRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
}

// AddCustomFieldConfigRequest 创建自定义字段（需求/缺陷/任务/测试用例）的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/add_custom_field_config.html
type AddCustomFieldConfigRequest struct {
	WorkspaceID string // 必填：项目 ID
	EntryType   string // 必填：字段分类（story / bug / task / tcase）
	Name        string // 必填：字段名称
	Type        string // 必填：字段类型（select / multi_select / text / checkbox / radio / textarea / user_chooser / dateinput / datetime / float / integer / cascade_checkbox / cascade_radio）
	Memo        string // 可选：字段备注
	Option      string // 可选：字段值选项（select/multi_select/radio/checkbox 时为 "AA|BB|CC"，cascade 时为 JSON 字符串）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *AddCustomFieldConfigRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"entry_type":   r.EntryType,
		"name":         r.Name,
		"type":         r.Type,
	}
	setOptional(params, "memo", r.Memo)
	setOptional(params, "option", r.Option)
	return params
}

// UpdateSelectFieldOptionsRequest 更新下拉类型自定义字段候选值的请求参数（用于 bug / story）
// API 文档：
//   - https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_bug_select_field_options.html
//   - https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_story_select_field_options.html
type UpdateSelectFieldOptionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：自定义字段配置 ID（19 位长度）
	Options     string // 必填：候选值，以英文竖线 "|" 隔开
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateSelectFieldOptionsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
		"options":      r.Options,
	}
}

// UpdateCascadeFieldOptionsRequest 更新级联自定义字段候选值的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_cascade_field_options.html
type UpdateCascadeFieldOptionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：自定义字段配置 ID（19 位长度）
	Options     string // 必填：候选值，json 字符串结构，children 表示子项
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateCascadeFieldOptionsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
		"options":      r.Options,
	}
}

// GetWorkspaceSettingRequest 获取项目配置开关的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/get_workspace_setting.html
type GetWorkspaceSettingRequest struct {
	WorkspaceID string // 必填：项目 ID
	Type        string // 必填：配置名称（is_enabled_story_category / workspace_metrology）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetWorkspaceSettingRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"type":         r.Type,
	}
}

// UpdateSelectFieldOptionsUnifiedRequest 更新下拉类型自定义字段候选值（统一接口）的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/setting/update_select_field_options.html
type UpdateSelectFieldOptionsUnifiedRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：自定义字段配置 ID（19 位长度）
	Options     string // 必填：候选值，以英文竖线 "|" 隔开
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateSelectFieldOptionsUnifiedRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
		"options":      r.Options,
	}
}

// CopyWorkitemTypeSettingRequest 复制需求类别配置的请求参数
type CopyWorkitemTypeSettingRequest struct {
	SrcWorkspaceID    string // 必填：源项目 ID
	SrcWorkitemTypeID string // 必填：源需求类型 ID
	WorkspaceID       string // 必填：目标项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CopyWorkitemTypeSettingRequest) ToParams() map[string]string {
	return map[string]string{
		"src_workspace_id":     r.SrcWorkspaceID,
		"src_workitem_type_id": r.SrcWorkitemTypeID,
		"workspace_id":         r.WorkspaceID,
	}
}

// CopyBugSettingRequest 复制缺陷配置的请求参数
type CopyBugSettingRequest struct {
	SrcWorkspaceID    string // 必填：源项目 ID
	TargetWorkspaceID string // 必填：目标项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CopyBugSettingRequest) ToParams() map[string]string {
	return map[string]string{
		"src_workspace_id":    r.SrcWorkspaceID,
		"target_workspace_id": r.TargetWorkspaceID,
	}
}
