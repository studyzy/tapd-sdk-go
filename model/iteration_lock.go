// Package model 中的 iteration_lock.go 定义了 TAPD 迭代锁定/解锁请求模型
package model

// LockIterationRequest 锁定迭代的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/lock_iteration.html
type LockIterationRequest struct {
	WorkspaceID string // 必填：项目 ID
	IterationID string // 必填：迭代 ID
	LockTypes   string // 可选：锁定对象，取值 __ALL_STORY__ / __ALL_BUG__，多个使用英文逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *LockIterationRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"iteration_id": r.IterationID,
	}
	setOptional(params, "lock_types", r.LockTypes)
	return params
}

// UnlockIterationRequest 解锁迭代的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/iteration/unlock_iteration.html
type UnlockIterationRequest struct {
	WorkspaceID string // 必填：项目 ID
	IterationID string // 必填：迭代 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *UnlockIterationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"iteration_id": r.IterationID,
	}
}
