// Package model 中的 release.go 定义了 TAPD 发布计划数据模型
package model

// Release 表示 TAPD 发布计划
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/
type Release struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startdate,omitempty"`
	EndDate     string `json:"enddate,omitempty"`
	Status      string `json:"status,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

// CreateReleaseRequest 创建发布计划的请求参数
type CreateReleaseRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：发布计划名称
	Description string // 可选：描述
	StartDate   string // 可选：开始日期
	EndDate     string // 可选：结束日期
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateReleaseRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "startdate", r.StartDate)
	setOptional(params, "enddate", r.EndDate)
	return params
}

// UpdateReleaseRequest 更新发布计划的请求参数
type UpdateReleaseRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：发布计划 ID
	Name        string // 可选：发布计划名称
	Description string // 可选：描述
	StartDate   string // 可选：开始日期
	EndDate     string // 可选：结束日期
	Status      string // 可选：状态
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateReleaseRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "startdate", r.StartDate)
	setOptional(params, "enddate", r.EndDate)
	setOptional(params, "status", r.Status)
	return params
}

// LaunchForm 表示发布评审单
type LaunchForm struct {
	ID             string `json:"id,omitempty"`
	Title          string `json:"title,omitempty"`
	Name           string `json:"name,omitempty"`
	Creator        string `json:"creator,omitempty"`
	Created        string `json:"created,omitempty"`
	WorkspaceID    string `json:"workspace_id,omitempty"`
	Status         string `json:"status,omitempty"`
	VersionType    string `json:"version_type,omitempty"`
	Baseline       string `json:"baseline,omitempty"`
	ReleaseModel   string `json:"release_model,omitempty"`
	RoadmapVersion string `json:"roadmap_version,omitempty"`
	ReleaseType    string `json:"release_type,omitempty"`
	ChangeType     string `json:"change_type,omitempty"`
	SignedBy       string `json:"signed_by,omitempty"`
	ArchivedBy     string `json:"archived_by,omitempty"`
	CC             string `json:"cc,omitempty"`
	ChangeNotifier string `json:"change_notifier,omitempty"`
	Signed         string `json:"signed,omitempty"`
	Archived       string `json:"archived,omitempty"`
	SignerResult   string `json:"signer_result,omitempty"`
	SignerComment  string `json:"signer_comment,omitempty"`
	ReleaseResult  string `json:"release_result,omitempty"`
	ReleaseComment string `json:"release_comment,omitempty"`
	TestPath       string `json:"test_path,omitempty"`
	CreatedPath    string `json:"created_path,omitempty"`
	Remark         string `json:"remark,omitempty"`
	Participator   string `json:"participator,omitempty"`
	TemplateID     string `json:"template_id,omitempty"`
	IterationID    string `json:"iteration_id,omitempty"`
	ReleaseID      string `json:"release_id,omitempty"`
	Flows          string `json:"flows,omitempty"`
}

// GetLaunchFormsRequest 获取发布评审列表的请求参数
type GetLaunchFormsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：发布评审 ID
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间（支持时间查询）
	Title       string // 可选：标题
	Status      string // 可选：状态
	VersionType string // 可选：版本类型
	Baseline    string // 可选：基线
	ReleaseID   string // 可选：发布计划 ID
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetLaunchFormsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "title", r.Title)
	setOptional(params, "status", r.Status)
	setOptional(params, "version_type", r.VersionType)
	setOptional(params, "baseline", r.Baseline)
	setOptional(params, "release_id", r.ReleaseID)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// UpdateLaunchFormRequest 更新发布评审单的请求参数
type UpdateLaunchFormRequest struct {
	WorkspaceID    string            // 必填：项目 ID
	ID             string            // 必填：发布评审 ID
	Title          string            // 可选：标题
	Status         string            // 可选：状态
	VersionType    string            // 可选：版本类型
	Baseline       string            // 可选：基线
	ReleaseModel   string            // 可选：发布模块
	RoadmapVersion string            // 可选：路标版本
	ReleaseType    string            // 可选：发布类型
	ChangeType     string            // 可选：变更类型
	CustomFields   map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateLaunchFormRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "title", r.Title)
	setOptional(params, "status", r.Status)
	setOptional(params, "version_type", r.VersionType)
	setOptional(params, "baseline", r.Baseline)
	setOptional(params, "release_model", r.ReleaseModel)
	setOptional(params, "roadmap_version", r.RoadmapVersion)
	setOptional(params, "release_type", r.ReleaseType)
	setOptional(params, "change_type", r.ChangeType)
	for k, v := range r.CustomFields {
		setOptional(params, k, v)
	}
	return params
}
