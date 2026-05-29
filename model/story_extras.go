// Package model 中的 story_extras.go 定义了需求扩展接口（分类、模板、关联、前后置、批量等）
// 涉及的 API 文档位于 docs/api_reference/story/。
package model

import (
	"encoding/json"
	"fmt"
)

// CreateStoryCategoryRequest 创建需求分类的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/add_story_category.html
type CreateStoryCategoryRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：分类名称
	Description string // 可选：分类描述
	ParentID    string // 可选：父分类 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateStoryCategoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	setOptional(params, "parent_id", r.ParentID)
	return params
}

// UpdateStoryCategoryRequest 更新需求分类的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/update_story_category.html
type UpdateStoryCategoryRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：分类 ID
	Name        string // 可选：分类名称
	Description string // 可选：分类描述
	ParentID    string // 可选：父分类 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateStoryCategoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "parent_id", r.ParentID)
	return params
}

// CountStoryCategoriesRequest 查询需求分类数量的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_categories_count.html
type CountStoryCategoriesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：分类 ID（多 ID 用逗号分隔）
	Name        string // 可选：分类名称（模糊匹配）
	Description string // 可选：分类描述
	ParentID    string // 可选：父分类 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountStoryCategoriesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	setOptional(params, "parent_id", r.ParentID)
	return params
}

// AddStoryLinkRelationsRequest 创建需求关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/add_story_link_relations.html
type AddStoryLinkRelationsRequest struct {
	WorkspaceID   string // 必填：项目 ID
	SrcStoryID    string // 必填：源需求 ID
	TargetStoryID string // 必填：目标需求 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *AddStoryLinkRelationsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id":    r.WorkspaceID,
		"src_story_id":    r.SrcStoryID,
		"target_story_id": r.TargetStoryID,
	}
}

// StoryLinkRelation 表示需求与其它需求的关联关系
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_link_stories.html
type StoryLinkRelation struct {
	Type              string `json:"type,omitempty"`                // 关系类型：sync_copy/copy/derivation/direct_relate/sync_relate
	ID                string `json:"id,omitempty"`                  // 关联的需求 ID
	StoryID           string `json:"story_id,omitempty"`            // 原需求 ID
	WorkspaceID       string `json:"workspace_id,omitempty"`        // 项目 ID
	Actas             string `json:"actas,omitempty"`               // 角色（target=操作发起方）
	Created           string `json:"created,omitempty"`             // 创建时间
	Modified          string `json:"modified,omitempty"`            // 最后修改时间
	LinkedWorkspaceID int    `json:"linked_workspace_id,omitempty"` // 关联项目 ID
}

// GetLinkStoriesRequest 获取需求与其它需求关联关系的请求参数
type GetLinkStoriesRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetLinkStoriesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
	}
}

// AddStoryTcaseRequest 创建需求与测试用例关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/create_story_tcase.html
type AddStoryTcaseRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
	TcaseID     string // 必填：测试用例 ID（多 ID 用逗号分隔，最多 20 个）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *AddStoryTcaseRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
		"tcase_id":     r.TcaseID,
	}
}

// GetStoryTcaseRequest 获取需求与测试用例关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_tcase.html
type GetStoryTcaseRequest struct {
	WorkspaceID     string // 必填：项目 ID
	StoryID         string // 必填：需求 ID
	IncludeTestPlan string // 可选：是否包含测试计划（0/1，默认 1）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetStoryTcaseRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
	}
	setOptional(params, "include_test_plan", r.IncludeTestPlan)
	return params
}

// TestPlanStoryTcaseRelation 表示测试计划-需求-测试用例关联关系
type TestPlanStoryTcaseRelation struct {
	ID          string `json:"id,omitempty"`           // 关系 ID
	WorkspaceID string `json:"workspace_id,omitempty"` // 项目 ID
	TestPlanID  string `json:"test_plan_id,omitempty"` // 测试计划 ID
	StoryID     string `json:"story_id,omitempty"`     // 需求 ID
	TcaseID     string `json:"tcase_id,omitempty"`     // 测试用例 ID
	Sort        string `json:"sort,omitempty"`         // 显示排序
	Creator     string `json:"creator,omitempty"`      // 创建人
	Created     string `json:"created,omitempty"`      // 创建时间
}

// RemoveStoryBugRelationsRequest 解除需求与缺陷关联关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/remove_story_bug_raletions.html
type RemoveStoryBugRelationsRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
	BugID       string // 必填：缺陷 ID
	CurrentUser string // 可选：操作人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *RemoveStoryBugRelationsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
		"bug_id":       r.BugID,
	}
	setOptional(params, "current_user", r.CurrentUser)
	return params
}

// WorkitemTimeRelation 表示需求前后置关系
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_time_relative_stories.html
type WorkitemTimeRelation struct {
	ID              string `json:"id,omitempty"`                // 关系 ID
	WorkspaceID     string `json:"workspace_id,omitempty"`      // 源项目 ID
	WorkitemType    string `json:"workitem_type,omitempty"`     // 业务对象类型（固定 story）
	WorkitemID      string `json:"workitem_id,omitempty"`       // 源需求 ID
	SrcField        string `json:"src_field,omitempty"`         // 源字段（begin/due）
	DstWorkspaceID  string `json:"dst_workspace_id,omitempty"`  // 目标项目 ID
	DstWorkitemType string `json:"dst_workitem_type,omitempty"` // 目标业务对象类型
	DstWorkitemID   string `json:"dst_workitem_id,omitempty"`   // 目标需求 ID
	DstField        string `json:"dst_field,omitempty"`         // 目标字段（begin/due）
	RelationType    string `json:"relation_type,omitempty"`     // 依赖类型（before/after）
}

// GetTimeRelativeStoriesRequest 获取需求前后置关系的请求参数
type GetTimeRelativeStoriesRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetTimeRelativeStoriesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
	}
}

// SaveTimeRelationItem 表示一条待新增/修改的需求前后置关系
type SaveTimeRelationItem struct {
	WorkitemID    string // 起点需求 ID
	DstWorkitemID string // 终点需求 ID
	SrcField      string // 起点字段（begin/due）
	DstField      string // 终点字段（begin/due）
}

// SaveTimeRelationsRequest 批量新增或修改需求前后置关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/save_time_relations.html
type SaveTimeRelationsRequest struct {
	WorkspaceID string                 // 必填：项目 ID
	Relations   []SaveTimeRelationItem // 必填：待新增/修改的关系列表
	CurrentUser string                 // 必填：执行此操作的用户 nick
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *SaveTimeRelationsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"current_user": r.CurrentUser,
	}
	for i, rel := range r.Relations {
		setOptional(params, fmt.Sprintf("relations[%d][workitem_id]", i), rel.WorkitemID)
		setOptional(params, fmt.Sprintf("relations[%d][dst_workitem_id]", i), rel.DstWorkitemID)
		setOptional(params, fmt.Sprintf("relations[%d][src_field]", i), rel.SrcField)
		setOptional(params, fmt.Sprintf("relations[%d][dst_field]", i), rel.DstField)
	}
	return params
}

// DeleteTimeRelationItem 表示按节点删除的关系（起点+终点）
type DeleteTimeRelationItem struct {
	WorkitemID    string // 起点需求 ID
	DstWorkitemID string // 终点需求 ID
}

// DeleteTimeRelationsRequest 批量删除需求前后置关系的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/delete_time_relations.html
type DeleteTimeRelationsRequest struct {
	WorkspaceID string                   // 必填：项目 ID
	Relations   []DeleteTimeRelationItem // 可选：按节点删除
	RelationIDs []string                 // 可选：按 ID 删除
	CurrentUser string                   // 必填：执行此操作的用户 nick
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *DeleteTimeRelationsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"current_user": r.CurrentUser,
	}
	for i, rel := range r.Relations {
		setOptional(params, fmt.Sprintf("relations[%d][workitem_id]", i), rel.WorkitemID)
		setOptional(params, fmt.Sprintf("relations[%d][dst_workitem_id]", i), rel.DstWorkitemID)
	}
	for i, id := range r.RelationIDs {
		setOptional(params, fmt.Sprintf("relation_ids[%d]", i), id)
	}
	return params
}

// WorkitemTemplate 表示需求模板
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_template_list.html
type WorkitemTemplate struct {
	ID          string `json:"id,omitempty"`          // 模板 ID
	Name        string `json:"name,omitempty"`        // 模板名称
	Description string `json:"description,omitempty"` // 描述
	Sort        string `json:"sort,omitempty"`        // 排序
	Default     string `json:"default,omitempty"`     // 是否启用
	Creator     string `json:"creator,omitempty"`     // 提交人
	EditorType  string `json:"editor_type,omitempty"` // 详细描述类型
}

// GetStoryTemplateListRequest 获取需求模板列表的请求参数
type GetStoryTemplateListRequest struct {
	WorkspaceID    string // 必填：项目 ID
	WorkitemTypeID string // 可选：需求类别 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetStoryTemplateListRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "workitem_type_id", r.WorkitemTypeID)
	return params
}

// WorkitemTemplateField 表示需求模板字段
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_default_story_template.html
type WorkitemTemplateField struct {
	ID           string `json:"id,omitempty"`            // 模板字段 ID
	WorkspaceID  string `json:"workspace_id,omitempty"`  // 项目 ID
	Type         string `json:"type,omitempty"`          // 类型
	TemplateID   string `json:"template_id,omitempty"`   // 模板 ID
	Field        string `json:"field,omitempty"`         // 字段名
	Value        string `json:"value,omitempty"`         // 值
	Required     string `json:"required,omitempty"`      // 是否必填
	Sort         string `json:"sort,omitempty"`          // 排序
	LinkageRules string `json:"linkage_rules,omitempty"` // 显示规则
	DefaultValue string `json:"default_value,omitempty"` // 默认值
}

// GetDefaultStoryTemplateRequest 获取需求模板字段的请求参数
type GetDefaultStoryTemplateRequest struct {
	WorkspaceID      string // 必填：项目 ID
	TemplateID       string // 必填：模板 ID
	UsePriorityLabel string // 可选：是否将 priority 替换为 priority_label（0/1，默认 0）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetDefaultStoryTemplateRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"template_id":  r.TemplateID,
	}
	setOptional(params, "use_priority_label", r.UsePriorityLabel)
	return params
}

// WorkitemStepInfo 表示并行工作流的需求节点
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_steps.html
type WorkitemStepInfo struct {
	ID           string `json:"id,omitempty"`            // 节点 ID
	WorkspaceID  string `json:"workspace_id,omitempty"`  // 项目 ID
	EntityType   string `json:"entity_type,omitempty"`   // 实体类型
	WorkitemID   string `json:"workitem_id,omitempty"`   // 需求 ID
	Step         string `json:"step,omitempty"`          // 节点原名
	Status       string `json:"status,omitempty"`        // 节点状态
	Owner        string `json:"owner,omitempty"`         // 节点负责人
	Begin        string `json:"begin,omitempty"`         // 节点预计开始
	Due          string `json:"due,omitempty"`           // 节点预计结束
	Effort       string `json:"effort,omitempty"`        // 节点预估工时
	IterationID  string `json:"iteration_id,omitempty"`  // 节点迭代
	BeginTime    string `json:"begin_time,omitempty"`    // 实际开始时间
	CompleteTime string `json:"complete_time,omitempty"` // 实际完成时间
	TimeCost     string `json:"time_cost,omitempty"`     // 节点停留时长
	Completer    string `json:"completer,omitempty"`     // 操作完成人
}

// GetStoryStepsRequest 获取需求节点的请求参数
type GetStoryStepsRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetStoryStepsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
	}
}

// RemovedStory 表示回收站中的需求
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_removed_stories.html
type RemovedStory struct {
	ID            string `json:"id,omitempty"`             // 需求 ID
	Name          string `json:"name,omitempty"`           // 标题
	Creator       string `json:"creator,omitempty"`        // 创建人
	Created       string `json:"created,omitempty"`        // 创建时间
	OperationUser string `json:"operation_user,omitempty"` // 删除人
	Deleted       string `json:"deleted,omitempty"`        // 删除时间
	IsArchived    string `json:"is_archived,omitempty"`    // 是否归档
}

// GetRemovedStoriesRequest 获取回收站需求的请求参数
type GetRemovedStoriesRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：需求 ID
	Creator     string // 可选：创建人
	IsArchived  string // 可选：是否归档（默认 0）
	Created     string // 可选：创建时间
	Deleted     string // 可选：删除时间
	Limit       int    // 可选：返回数量限制（默认 30，最大 200）
	Page        int    // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetRemovedStoriesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "creator", r.Creator)
	setOptional(params, "is_archived", r.IsArchived)
	setOptional(params, "created", r.Created)
	setOptional(params, "deleted", r.Deleted)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// CopyStoryRequest 复制需求的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/copy_story.html
type CopyStoryRequest struct {
	WorkspaceID       string // 必填：源项目 ID
	SrcStoryID        string // 必填：源需求 ID
	DstWorkspaceID    string // 必填：目标项目 ID
	SyncFields        string // 可选：需要同步的字段（逗号分隔）
	DstWorkitemTypeID string // 可选：目标需求类别 ID
	NewCreator        string // 可选：新需求创建人
	NewStatus         string // 可选：新需求状态
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CopyStoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id":     r.WorkspaceID,
		"src_story_id":     r.SrcStoryID,
		"dst_workspace_id": r.DstWorkspaceID,
	}
	setOptional(params, "sync_fields", r.SyncFields)
	setOptional(params, "dst_workitem_type_id", r.DstWorkitemTypeID)
	setOptional(params, "new_creator", r.NewCreator)
	setOptional(params, "new_status", r.NewStatus)
	return params
}

// UpdateStoryParentRequest 更新父需求的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/update_story_parent.html
type UpdateStoryParentRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
	ParentID    string // 必填：父需求 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateStoryParentRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
		"parent_id":    r.ParentID,
	}
}

// ChangeWorkitemTypeRequest 更新需求的需求类别的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/change_workitem_type.html
type ChangeWorkitemTypeRequest struct {
	WorkspaceID    string // 必填：项目 ID
	StoryID        string // 必填：需求 ID
	WorkitemTypeID string // 必填：目标需求类别 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ChangeWorkitemTypeRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id":     r.WorkspaceID,
		"story_id":         r.StoryID,
		"workitem_type_id": r.WorkitemTypeID,
	}
}

// GetStoriesByViewConfIDRequest 获取视图对应的需求列表的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_stories_by_view_conf_id.html
type GetStoriesByViewConfIDRequest struct {
	WorkspaceID string // 必填：项目 ID
	ViewConfID  string // 必填：视图 ID
	CurrentUser string // 可选：当前登录用户
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
	Fields      string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetStoriesByViewConfIDRequest) ToParams() map[string]string {
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

// BatchUpdateStoryItem 批量更新中的单条需求字段
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/batch_update_story.html
type BatchUpdateStoryItem struct {
	ID              string            `json:"id,omitempty"`                 // 必填：需求 ID
	Name            string            `json:"name,omitempty"`               // 可选：标题
	Priority        string            `json:"priority,omitempty"`           // 可选：优先级（建议使用 PriorityLabel）
	PriorityLabel   string            `json:"priority_label,omitempty"`     // 可选：优先级
	BusinessValue   string            `json:"business_value,omitempty"`     // 可选：业务价值
	Status          string            `json:"status,omitempty"`             // 可选：状态
	VStatus         string            `json:"v_status,omitempty"`           // 可选：中文状态名
	Version         string            `json:"version,omitempty"`            // 可选：版本
	Module          string            `json:"module,omitempty"`             // 可选：模块
	TestFocus       string            `json:"test_focus,omitempty"`         // 可选：测试重点
	Size            string            `json:"size,omitempty"`               // 可选：规模
	Owner           string            `json:"owner,omitempty"`              // 可选：处理人
	CurrentUser     string            `json:"current_user,omitempty"`       // 可选：变更人
	CC              string            `json:"cc,omitempty"`                 // 可选：抄送人
	Developer       string            `json:"developer,omitempty"`          // 可选：开发人员
	Begin           string            `json:"begin,omitempty"`              // 可选：预计开始
	Due             string            `json:"due,omitempty"`                // 可选：预计结束
	IterationID     string            `json:"iteration_id,omitempty"`       // 可选：迭代 ID
	Effort          string            `json:"effort,omitempty"`             // 可选：预估工时
	EffortCompleted string            `json:"effort_completed,omitempty"`   // 可选：完成工时
	Remain          string            `json:"remain,omitempty"`             // 可选：剩余工时
	Exceed          string            `json:"exceed,omitempty"`             // 可选：超出工时
	CategoryID      string            `json:"category_id,omitempty"`        // 可选：需求分类
	ReleaseID       string            `json:"release_id,omitempty"`         // 可选：发布计划
	Source          string            `json:"source,omitempty"`             // 可选：来源
	Type            string            `json:"type,omitempty"`               // 可选：类型
	Description     string            `json:"description,omitempty"`        // 可选：详细描述
	IsAutoCloseTask string            `json:"is_auto_close_task,omitempty"` // 可选：是否自动关闭关联任务
	Label           string            `json:"label,omitempty"`              // 可选：标签
	CustomFields    map[string]string `json:"-"`                            // 可选：自定义字段
}

// MarshalJSON 自定义序列化，将 CustomFields 合并到输出 JSON
func (i BatchUpdateStoryItem) MarshalJSON() ([]byte, error) {
	type Alias BatchUpdateStoryItem
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

// BatchUpdateStoryRequest 批量更新需求的请求参数
type BatchUpdateStoryRequest struct {
	WorkspaceID string                 // 必填：项目 ID
	Workitems   []BatchUpdateStoryItem // 必填：要更新的需求列表（单次最多 50 条）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *BatchUpdateStoryRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	if len(r.Workitems) > 0 {
		raw, _ := json.Marshal(r.Workitems)
		params["workitems"] = string(raw)
	}
	return params
}
