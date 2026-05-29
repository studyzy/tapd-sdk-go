// Package model 中的 tcase_test_plan.go 定义 TAPD 测试计划相关数据模型与请求参数
package model

import "encoding/json"

// TestPlan 表示 TAPD 测试计划
// 自定义字段（custom_field_*）通过 CustomFields map 保留，不会丢失
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_test_plans.html
type TestPlan struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Status      string `json:"status,omitempty"`
	Type        string `json:"type,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Modifier    string `json:"modifier,omitempty"`
	CreatedFrom string `json:"created_from,omitempty"`
	IterationID string `json:"iteration_id,omitempty"`
	TemplateID  string `json:"template_id,omitempty"`

	// 自定义字段，key 为 custom_field_1、custom_field_2 等
	CustomFields map[string]string `json:"-"`
}

// UnmarshalJSON 自定义反序列化，在解析标准字段的同时收集 custom_field_* 字段
func (tp *TestPlan) UnmarshalJSON(data []byte) error {
	type Alias TestPlan
	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*tp = TestPlan(alias)

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	tp.CustomFields = ExtractCustomFields(raw)
	return nil
}

// MarshalJSON 自定义序列化，将 CustomFields 中的键值对合并到输出 JSON
func (tp TestPlan) MarshalJSON() ([]byte, error) {
	type Alias TestPlan
	b, err := json.Marshal(Alias(tp))
	if err != nil {
		return nil, err
	}
	if len(tp.CustomFields) == 0 {
		return b, nil
	}

	var base map[string]json.RawMessage
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}
	for k, v := range tp.CustomFields {
		raw, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		base[k] = raw
	}
	return json.Marshal(base)
}

// TestPlanProgress 测试计划执行进度
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_test_plan_progress.html
type TestPlanProgress struct {
	StoryCount    int            `json:"story_count,omitempty"`
	TCaseCount    int            `json:"tcase_count,omitempty"`
	StatusCounter map[string]any `json:"status_counter,omitempty"`
	ExecutedRate  string         `json:"executed_rate,omitempty"`
}

// CreateTestPlanRequest 创建测试计划的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/add_test_plan.html
type CreateTestPlanRequest struct {
	WorkspaceID  string            // 必填：项目 ID
	Name         string            // 必填：测试计划标题
	Description  string            // 可选：详细描述
	Creator      string            // 可选：创建人
	Modifier     string            // 可选：修改人
	Owner        string            // 可选：负责人
	StartDate    string            // 可选：预计开始
	EndDate      string            // 可选：预计结束
	IterationID  string            // 可选：关联迭代 ID
	Version      string            // 可选：版本号
	Type         string            // 可选：测试类型
	Status       string            // 可选：状态
	CustomFields map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateTestPlanRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "modifier", r.Modifier)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "start_date", r.StartDate)
	setOptional(params, "end_date", r.EndDate)
	setOptional(params, "iteration_id", r.IterationID)
	setOptional(params, "version", r.Version)
	setOptional(params, "type", r.Type)
	setOptional(params, "status", r.Status)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// UpdateTestPlanRequest 编辑测试计划的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/update_test_plan.html
type UpdateTestPlanRequest struct {
	WorkspaceID  string            // 必填：项目 ID
	ID           string            // 必填：测试计划 ID
	Name         string            // 可选：标题
	Description  string            // 可选：详细描述
	Modifier     string            // 可选：修改人
	Owner        string            // 可选：负责人
	StartDate    string            // 可选：预计开始
	EndDate      string            // 可选：预计结束
	Version      string            // 可选：版本号
	Type         string            // 可选：测试类型
	Status       string            // 可选：状态
	TemplateID   string            // 可选：模板 ID
	CustomFields map[string]string // 可选：自定义字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateTestPlanRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "modifier", r.Modifier)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "start_date", r.StartDate)
	setOptional(params, "end_date", r.EndDate)
	setOptional(params, "version", r.Version)
	setOptional(params, "type", r.Type)
	setOptional(params, "status", r.Status)
	setOptional(params, "template_id", r.TemplateID)
	MergeCustomFields(params, r.CustomFields)
	return params
}

// ListTestPlansRequest 查询测试计划列表/数量的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_test_plans.html
type ListTestPlansRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：测试计划 ID
	Name        string // 可选：标题
	Description string // 可选：详细描述
	Creator     string // 可选：创建人
	Created     string // 可选：创建时间
	Modifier    string // 可选：修改人
	Modified    string // 可选：最后修改时间
	Owner       string // 可选：负责人
	StartDate   string // 可选：预计开始
	EndDate     string // 可选：预计结束
	IterationID string // 可选：关联迭代 ID
	Version     string // 可选：版本号
	Type        string // 可选：测试类型
	Status      string // 可选：状态
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ListTestPlansRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "created", r.Created)
	setOptional(params, "modifier", r.Modifier)
	setOptional(params, "modified", r.Modified)
	setOptional(params, "owner", r.Owner)
	setOptional(params, "start_date", r.StartDate)
	setOptional(params, "end_date", r.EndDate)
	setOptional(params, "iteration_id", r.IterationID)
	setOptional(params, "version", r.Version)
	setOptional(params, "type", r.Type)
	setOptional(params, "status", r.Status)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}

// TestPlanIDRequest 仅需 workspace_id + 测试计划 id 的请求参数
// 适用于 GetTestPlanDetails、GetTestPlanProgress、GetTestPlanBugs
type TestPlanIDRequest struct {
	WorkspaceID   string // 必填：项目 ID
	ID            string // 必填：测试计划 ID
	LastExecutor  string // 可选：最后执行人（仅 details 接口）
	IncludeRepeat int    // 可选：=1 获取所有数据（仅 details 接口）
	Limit         int    // 可选：返回数量限制（仅 bugs 接口）
	Page          int    // 可选：页码（仅 bugs 接口）
	Order         string // 可选：排序规则（仅 bugs 接口）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TestPlanIDRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "last_executor", r.LastExecutor)
	setOptionalInt(params, "include_repeat", r.IncludeRepeat)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	return params
}

// TestPlanRelativeStoriesRequest 获取测试计划关联需求的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_test_plan_relative_stories.html
type TestPlanRelativeStoriesRequest struct {
	WorkspaceID string // 必填：项目 ID
	TestPlanID  string // 必填：测试计划 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TestPlanRelativeStoriesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"test_plan_id": r.TestPlanID,
	}
}

// TestPlanTCasesRequest 获取测试计划与测试用例关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_test_plan_tcases.html
type TestPlanTCasesRequest struct {
	WorkspaceID string // 必填：项目 ID
	TestPlanID  string // 必填：测试计划 ID
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TestPlanTCasesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"test_plan_id": r.TestPlanID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	return params
}

// TestPlansByIterationIDRequest 获取迭代下测试计划的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_by_iteration_id.html
type TestPlansByIterationIDRequest struct {
	WorkspaceID string // 必填：项目 ID
	IterationID string // 必填：迭代 ID
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Order       string // 可选：排序规则
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TestPlansByIterationIDRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"iteration_id": r.IterationID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	return params
}

// TestPlanStoryRelationRequest 创建/解除测试计划与需求关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/create_story_relation.html
//
//	https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/delete_story_relation.html
type TestPlanStoryRelationRequest struct {
	WorkspaceID string // 必填：项目 ID
	PlanID      string // 必填：测试计划 ID
	StoryIDs    string // 必填：需求 ID，多个用 , 隔开（最多 10）
	Creator     string // 必填：创建人/操作人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TestPlanStoryRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"plan_id":      r.PlanID,
		"story_ids":    r.StoryIDs,
		"creator":      r.Creator,
	}
}

// TestPlanTCaseRelationRequest 创建测试计划与测试用例关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/create_tcase_relation.html
type TestPlanTCaseRelationRequest struct {
	WorkspaceID string // 必填：项目 ID
	TestPlanID  string // 必填：测试计划 ID
	TCaseIDs    string // 必填：测试用例 ID，多个用 , 隔开（最多 10）
	Creator     string // 必填：创建人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TestPlanTCaseRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"test_plan_id": r.TestPlanID,
		"tcase_ids":    r.TCaseIDs,
		"creator":      r.Creator,
	}
}
