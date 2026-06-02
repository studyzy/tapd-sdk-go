// Package model 中的 story.go 定义了 TAPD 需求数据模型
package model

import "encoding/json"

// Story 表示 TAPD 需求/工作项，字段覆盖 TAPD API 返回的所有常用字段
// 自定义字段（custom_field_*、custom_plan_field_*）通过 CustomFields map 保留，不会丢失
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/story.html
type Story struct {
	// 基本信息
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	WorkspaceID    string `json:"workspace_id,omitempty"`
	Status         string `json:"status,omitempty"`
	Step           string `json:"step,omitempty"`
	Type           string `json:"type,omitempty"`
	Source         string `json:"source,omitempty"`
	Flows          string `json:"flows,omitempty"`
	CreatedFrom    string `json:"created_from,omitempty"`
	WorkitemTypeID string `json:"workitem_type_id,omitempty"`
	TemplatedID    string `json:"templated_id,omitempty"`

	// 优先级
	Priority      string `json:"priority,omitempty"`
	PriorityLabel string `json:"priority_label,omitempty"`
	BusinessValue string `json:"business_value,omitempty"`

	// 人员相关
	Owner     string `json:"owner,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Developer string `json:"developer,omitempty"`
	CC        string `json:"cc,omitempty"`

	// 时间相关
	Created   string `json:"created,omitempty"`
	Modified  string `json:"modified,omitempty"`
	Completed string `json:"completed,omitempty"`
	Begin     string `json:"begin,omitempty"`
	Due       string `json:"due,omitempty"`

	// 关联与分类
	IterationID string `json:"iteration_id,omitempty"`
	Module      string `json:"module,omitempty"`
	Feature     string `json:"feature,omitempty"`
	Label       string `json:"label,omitempty"`
	CategoryID  string `json:"category_id,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
	ChildrenID  string `json:"children_id,omitempty"`
	AncestorID  string `json:"ancestor_id,omitempty"`
	Path        string `json:"path,omitempty"`
	Level       string `json:"level,omitempty"`
	ReleaseID   string `json:"release_id,omitempty"`
	BugID       string `json:"bug_id,omitempty"`
	Version     string `json:"version,omitempty"`

	// 规模与工时
	Size            string `json:"size,omitempty"`
	Effort          string `json:"effort,omitempty"`
	EffortCompleted string `json:"effort_completed,omitempty"`
	Remain          string `json:"remain,omitempty"`
	Exceed          string `json:"exceed,omitempty"`

	// 进度与风险
	Progress       string `json:"progress,omitempty"`
	ProgressManual string `json:"progress_manual,omitempty"`
	TechRisk       string `json:"tech_risk,omitempty"`
	TestFocus      string `json:"test_focus,omitempty"`
	IsArchived     string `json:"is_archived,omitempty"`

	// 保密
	SecretRootID string `json:"secret_root_id,omitempty"`

	// 附加信息
	Sort                string `json:"sort,omitempty"`
	Participator        string `json:"participator,omitempty"`
	Follower            string `json:"follower,omitempty"`
	Confidential        string `json:"confidential,omitempty"`
	EntityType          string `json:"entity_type,omitempty"`
	AttachmentCount     string `json:"attachment_count,omitempty"`
	DescriptionType     string `json:"description_type,omitempty"`
	MarkdownDescription string `json:"markdown_description,omitempty"`
	PredecessorCount    string `json:"predecessor_count,omitempty"`
	SuccessorCount      string `json:"successor_count,omitempty"`
	SyncType            string `json:"sync_type,omitempty"`
	LastModify          string `json:"lastmodify,omitempty"`
	URL                 string `json:"url,omitempty"`

	// 自定义字段，key 为 custom_field_one、custom_field_9 等
	CustomFields map[string]string `json:"-"`
}

// UnmarshalJSON 自定义反序列化，在解析标准字段的同时收集 custom_field_* 和 custom_plan_field_* 字段
func (s *Story) UnmarshalJSON(data []byte) error {
	type Alias Story
	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*s = Story(alias)

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.CustomFields = ExtractCustomFields(raw)
	return nil
}

// MarshalJSON 自定义序列化，将 CustomFields 中的键值对合并到输出 JSON
func (s Story) MarshalJSON() ([]byte, error) {
	type Alias Story
	b, err := json.Marshal(Alias(s))
	if err != nil {
		return nil, err
	}
	if len(s.CustomFields) == 0 {
		return b, nil
	}

	// 将自定义字段合并到 JSON 对象中
	var base map[string]json.RawMessage
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}
	for k, v := range s.CustomFields {
		raw, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		base[k] = raw
	}
	return json.Marshal(base)
}

// ListStoriesRequest 查询需求列表的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_stories.html
type ListStoriesRequest struct {
	WorkspaceID         string            // 必填：项目 ID
	ID                  string            // 可选：需求 ID
	Name                string            // 可选：标题（支持模糊匹配）
	Priority            string            // 可选：优先级（数字）
	PriorityLabel       string            // 可选：优先级标签
	BusinessValue       string            // 可选：业务价值
	VStatus             string            // 可选：中文状态名
	WithVStatus         string            // 可选：带中文状态名返回
	Status              string            // 可选：状态
	Owner               string            // 可选：处理人
	Creator             string            // 可选：创建人
	Developer           string            // 可选：开发人员
	CC                  string            // 可选：抄送人
	IterationID         string            // 可选：迭代 ID
	CategoryID          string            // 可选：需求分类
	Label               string            // 可选：标签
	Version             string            // 可选：版本
	Module              string            // 可选：模块
	TestFocus           string            // 可选：测试重点
	Size                string            // 可选：规模
	Begin               string            // 可选：预计开始日期
	Due                 string            // 可选：预计结束日期
	Created             string            // 可选：创建时间
	Modified            string            // 可选：修改时间
	Completed           string            // 可选：完成时间
	Type                string            // 可选：需求类型
	Source              string            // 可选：需求来源
	Feature             string            // 可选：特性
	ParentID            string            // 可选：父需求 ID
	WorkitemTypeID      string            // 可选：需求类别 ID
	ReleaseID           string            // 可选：发布计划 ID
	TechRisk            string            // 可选：技术风险
	Effort              string            // 可选：预估工时
	EffortCompleted     string            // 可选：已完成工时
	Remain              string            // 可选：剩余工时
	Exceed              string            // 可选：超出工时
	AncestorID          string            // 可选：祖先需求 ID
	ChildrenID          string            // 可选：子需求 ID
	Description         string            // 可选：详细描述
	IncludeSubIteration string            // 可选：是否包含子迭代
	IncludeSubCategory  string            // 可选：是否包含子分类
	IncludeLeafStories  string            // 可选：是否包含叶子需求
	Fields              string            // 可选：返回字段列表
	Limit               int               // 可选：返回数量限制
	Page                int               // 可选：页码
	Order               string            // 可选：排序规则
	CustomFields        map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListStoriesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "priority_label", r.PriorityLabel)
	setOptional(params, "business_value", r.BusinessValue)
	setOptional(params, "v_status", r.VStatus)
	setOptional(params, "with_v_status", r.WithVStatus)
	setOptional(params, "status", r.Status)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "developer", r.Developer)
	setOptional(params, "cc", r.CC)
	setOptional(params, "iteration_id", r.IterationID)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "label", r.Label)
	setOptional(params, "version", r.Version)
	setOptional(params, "module", r.Module)
	setOptional(params, "test_focus", r.TestFocus)
	setOptional(params, "size", r.Size)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "type", r.Type)
	setOptional(params, "source", r.Source)
	setOptional(params, "feature", r.Feature)
	setOptional(params, "parent_id", r.ParentID)
	setOptional(params, "workitem_type_id", r.WorkitemTypeID)
	setOptional(params, "release_id", r.ReleaseID)
	setOptional(params, "tech_risk", r.TechRisk)
	setOptional(params, "effort", r.Effort)
	setOptional(params, "effort_completed", r.EffortCompleted)
	setOptional(params, "remain", r.Remain)
	setOptional(params, "exceed", r.Exceed)
	setOptional(params, "ancestor_id", r.AncestorID)
	setOptional(params, "children_id", r.ChildrenID)
	setOptional(params, "description", r.Description)
	setOptional(params, "include_sub_iteration", r.IncludeSubIteration)
	setOptional(params, "include_sub_category", r.IncludeSubCategory)
	setOptional(params, "include_leaf_stories", r.IncludeLeafStories)
	setOptional(params, "fields", r.Fields)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// CreateStoryRequest 创建需求的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/add_story.html
type CreateStoryRequest struct {
	WorkspaceID                 string            // 必填：项目 ID
	Name                        string            // 必填：标题
	Description                 string            // 可选：详细描述
	Priority                    string            // 可选：优先级（数字）
	PriorityLabel               string            // 可选：优先级标签
	BusinessValue               string            // 可选：业务价值
	Owner                       string            // 可选：处理人
	Creator                     string            // 可选：创建人
	Developer                   string            // 可选：开发人员
	CC                          string            // 可选：抄送人
	IterationID                 string            // 可选：迭代 ID
	ParentID                    string            // 可选：父需求 ID（创建子需求时使用）
	CategoryID                  string            // 可选：需求分类 ID
	Type                        string            // 可选：需求类型
	Source                      string            // 可选：需求来源
	Begin                       string            // 可选：预计开始日期
	Due                         string            // 可选：预计结束日期
	Label                       string            // 可选：标签
	Version                     string            // 可选：版本
	Module                      string            // 可选：模块
	TestFocus                   string            // 可选：测试重点
	Size                        string            // 可选：规模
	Effort                      string            // 可选：预估工时
	EffortCompleted             string            // 可选：已完成工时
	Remain                      string            // 可选：剩余工时
	Exceed                      string            // 可选：超出工时
	ReleaseID                   string            // 可选：发布计划 ID
	Feature                     string            // 可选：特性
	TemplatedID                 string            // 可选：模板 ID
	WorkitemTypeID              string            // 可选：需求类别 ID
	TechRisk                    string            // 可选：技术风险
	IsApplyTemplateDefaultValue string            // 可选：是否应用模板默认值
	ApplyTemplate               string            // 可选：应用模板
	CustomFields                map[string]string // 可选：自定义字段，key 如 custom_field_one、custom_field_9
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateStoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "priority_label", r.PriorityLabel)
	setOptional(params, "business_value", r.BusinessValue)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "developer", r.Developer)
	setOptional(params, "cc", r.CC)
	setOptional(params, "iteration_id", r.IterationID)
	setOptional(params, "parent_id", r.ParentID)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "type", r.Type)
	setOptional(params, "source", r.Source)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "label", r.Label)
	setOptional(params, "version", r.Version)
	setOptional(params, "module", r.Module)
	setOptional(params, "test_focus", r.TestFocus)
	setOptional(params, "size", r.Size)
	setOptional(params, "effort", r.Effort)
	setOptional(params, "effort_completed", r.EffortCompleted)
	setOptional(params, "remain", r.Remain)
	setOptional(params, "exceed", r.Exceed)
	setOptional(params, "release_id", r.ReleaseID)
	setOptional(params, "feature", r.Feature)
	setOptional(params, "templated_id", r.TemplatedID)
	setOptional(params, "workitem_type_id", r.WorkitemTypeID)
	setOptional(params, "tech_risk", r.TechRisk)
	setOptional(params, "is_apply_template_default_value", r.IsApplyTemplateDefaultValue)
	setOptional(params, "apply_template", r.ApplyTemplate)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// UpdateStoryRequest 更新需求的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/update_story.html
type UpdateStoryRequest struct {
	WorkspaceID     string            // 必填：项目 ID
	ID              string            // 必填：需求 ID
	Name            string            // 可选：标题
	Status          string            // 可选：状态
	VStatus         string            // 可选：中文状态名
	Priority        string            // 可选：优先级（数字）
	PriorityLabel   string            // 可选：优先级标签
	BusinessValue   string            // 可选：业务价值
	Owner           string            // 可选：处理人
	CurrentUser     string            // 可选：变更人
	Developer       string            // 可选：开发人员
	CC              string            // 可选：抄送人
	Description     string            // 可选：详细描述
	IterationID     string            // 可选：迭代 ID
	CategoryID      string            // 可选：需求分类 ID
	Begin           string            // 可选：预计开始日期
	Due             string            // 可选：预计结束日期
	Label           string            // 可选：标签
	Version         string            // 可选：版本
	Module          string            // 可选：模块
	TestFocus       string            // 可选：测试重点
	Size            string            // 可选：规模
	Effort          string            // 可选：预估工时
	EffortCompleted string            // 可选：已完成工时
	Remain          string            // 可选：剩余工时
	Exceed          string            // 可选：超出工时
	ReleaseID       string            // 可选：发布计划 ID
	Type            string            // 可选：需求类型
	Source          string            // 可选：需求来源
	IsAutoCloseTask string            // 可选：是否自动关闭关联任务
	CustomFields    map[string]string // 可选：自定义字段，key 如 custom_field_one、custom_field_9
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateStoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "status", r.Status)
	setOptional(params, "v_status", r.VStatus)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "priority_label", r.PriorityLabel)
	setOptional(params, "business_value", r.BusinessValue)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "current_user", r.CurrentUser)
	setOptional(params, "developer", r.Developer)
	setOptional(params, "cc", r.CC)
	setOptional(params, "description", r.Description)
	setOptional(params, "iteration_id", r.IterationID)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "label", r.Label)
	setOptional(params, "version", r.Version)
	setOptional(params, "module", r.Module)
	setOptional(params, "test_focus", r.TestFocus)
	setOptional(params, "size", r.Size)
	setOptional(params, "effort", r.Effort)
	setOptional(params, "effort_completed", r.EffortCompleted)
	setOptional(params, "remain", r.Remain)
	setOptional(params, "exceed", r.Exceed)
	setOptional(params, "release_id", r.ReleaseID)
	setOptional(params, "type", r.Type)
	setOptional(params, "source", r.Source)
	setOptional(params, "is_auto_close_task", r.IsAutoCloseTask)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// CountStoriesRequest 查询需求数量的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_stories_count.html
type CountStoriesRequest struct {
	WorkspaceID         string            // 必填：项目 ID
	ID                  string            // 可选：需求 ID
	Name                string            // 可选：标题（支持模糊匹配）
	Priority            string            // 可选：优先级（数字）
	PriorityLabel       string            // 可选：优先级标签
	BusinessValue       string            // 可选：业务价值
	VStatus             string            // 可选：中文状态名
	WithVStatus         string            // 可选：带中文状态名返回
	Status              string            // 可选：状态
	Owner               string            // 可选：处理人
	Creator             string            // 可选：创建人
	Developer           string            // 可选：开发人员
	CC                  string            // 可选：抄送人
	IterationID         string            // 可选：迭代 ID
	CategoryID          string            // 可选：需求分类
	Label               string            // 可选：标签
	Version             string            // 可选：版本
	Module              string            // 可选：模块
	TestFocus           string            // 可选：测试重点
	Size                string            // 可选：规模
	Begin               string            // 可选：预计开始日期
	Due                 string            // 可选：预计结束日期
	Created             string            // 可选：创建时间
	Modified            string            // 可选：修改时间
	Completed           string            // 可选：完成时间
	Type                string            // 可选：需求类型
	Source              string            // 可选：需求来源
	Feature             string            // 可选：特性
	ParentID            string            // 可选：父需求 ID
	WorkitemTypeID      string            // 可选：需求类别 ID
	ReleaseID           string            // 可选：发布计划 ID
	TechRisk            string            // 可选：技术风险
	Effort              string            // 可选：预估工时
	EffortCompleted     string            // 可选：已完成工时
	Remain              string            // 可选：剩余工时
	Exceed              string            // 可选：超出工时
	AncestorID          string            // 可选：祖先需求 ID
	ChildrenID          string            // 可选：子需求 ID
	Description         string            // 可选：详细描述
	IncludeSubCategory  string            // 可选：是否包含子分类
	IncludeSubIteration string            // 可选：是否包含子迭代
	IncludeLeafStories  string            // 可选：是否包含叶子需求
	Limit               int               // 可选：返回数量限制
	Page                int               // 可选：页码
	Order               string            // 可选：排序规则
	Fields              string            // 可选：返回字段
	CustomFields        map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountStoriesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "priority", r.Priority)
	setOptional(params, "priority_label", r.PriorityLabel)
	setOptional(params, "business_value", r.BusinessValue)
	setOptional(params, "v_status", r.VStatus)
	setOptional(params, "with_v_status", r.WithVStatus)
	setOptional(params, "status", r.Status)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "developer", r.Developer)
	setOptional(params, "cc", r.CC)
	setOptional(params, "iteration_id", r.IterationID)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "label", r.Label)
	setOptional(params, "version", r.Version)
	setOptional(params, "module", r.Module)
	setOptional(params, "test_focus", r.TestFocus)
	setOptional(params, "size", r.Size)
	setOptional(params, "begin", r.Begin)
	setOptional(params, "due", r.Due)
	setOptional(params, "created", r.Created)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "completed", r.Completed)
	setOptional(params, "type", r.Type)
	setOptional(params, "source", r.Source)
	setOptional(params, "feature", r.Feature)
	setOptional(params, "parent_id", r.ParentID)
	setOptional(params, "workitem_type_id", r.WorkitemTypeID)
	setOptional(params, "release_id", r.ReleaseID)
	setOptional(params, "tech_risk", r.TechRisk)
	setOptional(params, "effort", r.Effort)
	setOptional(params, "effort_completed", r.EffortCompleted)
	setOptional(params, "remain", r.Remain)
	setOptional(params, "exceed", r.Exceed)
	setOptional(params, "ancestor_id", r.AncestorID)
	setOptional(params, "children_id", r.ChildrenID)
	setOptional(params, "description", r.Description)
	setOptional(params, "include_sub_category", r.IncludeSubCategory)
	setOptional(params, "include_sub_iteration", r.IncludeSubIteration)
	setOptional(params, "include_leaf_stories", r.IncludeLeafStories)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	MergeCustomFields(params, r.CustomFields)
	return params
}
