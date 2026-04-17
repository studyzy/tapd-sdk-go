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
}

// CreateModuleRequest 创建模块的请求参数
type CreateModuleRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：模块名称
	Description string // 可选：模块描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateModuleRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	return params
}

// UpdateModuleRequest 更新模块的请求参数
type UpdateModuleRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：模块 ID
	Name        string // 可选：模块名称
	Description string // 可选：模块描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateModuleRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	return params
}

// GetModulesRequest 获取模块列表的请求参数
type GetModulesRequest struct {
	WorkspaceID string // 必填：项目 ID
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetModulesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// CountModulesRequest 获取模块数量的请求参数
type CountModulesRequest struct {
	WorkspaceID string // 必填：项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountModulesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
	}
}

// Version 表示 TAPD 版本
type Version struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

// CreateVersionRequest 创建版本的请求参数
type CreateVersionRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：版本名称
	Description string // 可选：版本描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateVersionRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	return params
}

// UpdateVersionRequest 更新版本的请求参数
type UpdateVersionRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：版本 ID
	Name        string // 可选：版本名称
	Description string // 可选：版本描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateVersionRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	return params
}

// GetVersionsRequest 获取版本列表的请求参数
type GetVersionsRequest struct {
	WorkspaceID string // 必填：项目 ID
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetVersionsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// CountVersionsRequest 获取版本数量的请求参数
type CountVersionsRequest struct {
	WorkspaceID string // 必填：项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountVersionsRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
	}
}

// Baseline 表示 TAPD 基线
type Baseline struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

// CreateBaselineRequest 创建基线的请求参数
type CreateBaselineRequest struct {
	WorkspaceID string // 必填：项目 ID
	Name        string // 必填：基线名称
	Description string // 可选：基线描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CreateBaselineRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"name":         r.Name,
	}
	setOptional(params, "description", r.Description)
	return params
}

// UpdateBaselineRequest 更新基线的请求参数
type UpdateBaselineRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 必填：基线 ID
	Name        string // 可选：基线名称
	Description string // 可选：基线描述
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateBaselineRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	return params
}

// GetBaselinesRequest 获取基线列表的请求参数
type GetBaselinesRequest struct {
	WorkspaceID string // 必填：项目 ID
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetBaselinesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// CountBaselinesRequest 获取基线数量的请求参数
type CountBaselinesRequest struct {
	WorkspaceID string // 必填：项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountBaselinesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
	}
}

// Feature 表示 TAPD 特性
type Feature struct {
	ID          string `json:"id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
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
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UpdateFeatureRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"id":           r.ID,
	}
	setOptional(params, "name", r.Name)
	setOptional(params, "description", r.Description)
	return params
}

// GetFeaturesRequest 获取特性列表的请求参数
type GetFeaturesRequest struct {
	WorkspaceID string // 必填：项目 ID
	Limit int // 可选：返回数量限制
	Page int // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetFeaturesRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}

// CountFeaturesRequest 获取特性数量的请求参数
type CountFeaturesRequest struct {
	WorkspaceID string // 必填：项目 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *CountFeaturesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
	}
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
