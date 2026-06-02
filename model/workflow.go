// Package model 中的 workflow.go 定义了 TAPD 工作流相关数据模型

package model

// WorkflowStep 工作流步骤
type WorkflowStep struct {
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
}

// WorkflowStepGroup 工作流步骤组（状态与步骤的对应关系）
type WorkflowStepGroup struct {
	Name  string         `json:"name,omitempty"`
	Label string         `json:"label,omitempty"`
	Steps []WorkflowStep `json:"steps,omitempty"`
}

// WorkflowTransition 表示工作流状态流转规则
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/workflow/get_workflow_all_transitions.html
type WorkflowTransition struct {
	Name           string                `json:"Name,omitempty"`           // 流转名称（格式：前状态-后状态）
	StepPrevious   string                `json:"StepPrevious,omitempty"`   // 当前状态
	StepNext       string                `json:"StepNext,omitempty"`       // 目标状态
	Inform         []WorkflowInform      `json:"Inform,omitempty"`         // 通知配置
	Appendfield    []WorkflowAppendField `json:"Appendfield,omitempty"`    // 流转时需补充的附加字段
	AuthorizedUser interface{}           `json:"AuthorizedUser,omitempty"` // 流转权限设置
}

// WorkflowInform 表示工作流流转通知配置
type WorkflowInform struct {
	InformType string `json:"InformType,omitempty"` // 通知类型（如 RTX、Email）
	InformId   string `json:"InformId,omitempty"`   // 通知规则 ID
}

// WorkflowAppendField 表示工作流流转时需补充的附加字段
type WorkflowAppendField struct {
	DBModel      string                      `json:"DBModel,omitempty"`      // 数据模型（如 Bug、Story）
	FieldName    string                      `json:"FieldName,omitempty"`    // 字段名称
	Notnull      string                      `json:"Notnull,omitempty"`      // 是否必填（yes/no）
	Sort         string                      `json:"Sort,omitempty"`         // 显示排序号
	DefaultValue []WorkflowFieldDefaultValue `json:"DefaultValue,omitempty"` // 默认值设置
}

// WorkflowFieldDefaultValue 表示附加字段的默认值设置
type WorkflowFieldDefaultValue struct {
	Type    string `json:"Type,omitempty"`    // 默认值类型（default_value/record_value）
	Value   string `json:"Value,omitempty"`   // 固定默认值（Type 为 default_value 时使用）
	DBModel string `json:"DBModel,omitempty"` // 引用的数据模型（Type 为 record_value 时使用）
	Field   string `json:"Field,omitempty"`   // 引用的字段名（Type 为 record_value 时使用）
}

// Workflow 表示工作流定义
type Workflow struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SystemName  string `json:"system_name,omitempty"`
	IsDefault   string `json:"is_default,omitempty"`
	Created     string `json:"created,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Modified    string `json:"modified,omitempty"`
	Modifier    string `json:"modifier,omitempty"`
	Type        string `json:"type,omitempty"`
}

// NewStepResult 表示添加工作流步骤的返回结果
type NewStepResult struct {
	SysName  string `json:"sys_name,omitempty"`
	StepName string `json:"step_name,omitempty"`
}
