// Package model 中的 tcase_instance.go 定义 TAPD 测试用例执行（tcase_instance）相关请求参数
package model

// AssignTCaseInstanceRequest 分配测试用例的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/assign_tcase_instance.html
type AssignTCaseInstanceRequest struct {
	WorkspaceID string // 必填：项目 ID
	TestPlanID  string // 必填：测试计划 ID
	TCaseID     string // 可选：用例 ID（与 CategoryID 不能同时为空），多个用 , 隔开
	CategoryID  string // 可选：用例目录 ID（与 TCaseID 不能同时为空）
	Executor    string // 可选：执行人
	Assignee    string // 可选：负责人
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *AssignTCaseInstanceRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"test_plan_id": r.TestPlanID,
	}
	setOptional(params, "tcase_id", r.TCaseID)
	setOptional(params, "category_id", r.CategoryID)
	setOptional(params, "executor", r.Executor)
	setOptional(params, "assignee", r.Assignee)
	return params
}

// ExecuteTCaseInstanceRequest 执行测试用例的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/execute_tcase_instance.html
type ExecuteTCaseInstanceRequest struct {
	WorkspaceID  string // 必填：项目 ID
	TestPlanID   string // 必填：测试计划 ID
	TCaseID      string // 必填：用例 ID，多个用 , 隔开
	ResultStatus string // 必填：执行结果（pass/no_pass/block）
	LastExecutor string // 必填：执行人
	ResultRemark string // 可选：实际执行结果
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *ExecuteTCaseInstanceRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id":  r.WorkspaceID,
		"test_plan_id":  r.TestPlanID,
		"tcase_id":      r.TCaseID,
		"result_status": r.ResultStatus,
		"last_executor": r.LastExecutor,
	}
	setOptional(params, "result_remark", r.ResultRemark)
	return params
}

// RemoveTCaseInstanceRequest 测试用例移出测试计划的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/remove_tcase_instance.html
type RemoveTCaseInstanceRequest struct {
	WorkspaceID string // 必填：项目 ID
	TestPlanID  string // 必填：测试计划 ID
	TCaseID     string // 必填：测试用例 ID，多个用 , 分割
	StoryID     string // 可选：需求 ID（仅当用例关联了需求时需要传入）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *RemoveTCaseInstanceRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"test_plan_id": r.TestPlanID,
		"tcase_id":     r.TCaseID,
	}
	setOptional(params, "story_id", r.StoryID)
	return params
}

// TCaseResultRequest 获取测试用例执行结果的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/get_tcase_result.html
type TCaseResultRequest struct {
	WorkspaceID string // 必填：项目 ID
	TestPlanID  string // 必填：测试计划 ID
	TCaseID     string // 必填：用例 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *TCaseResultRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"test_plan_id": r.TestPlanID,
		"tcase_id":     r.TCaseID,
	}
}

// DeleteTCaseStoryRelationRequest 解除测试用例关联并移出测试计划的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/tcase/delete_tcase_story_relation.html
type DeleteTCaseStoryRelationRequest struct {
	WorkspaceID string // 必填：项目 ID
	StoryID     string // 必填：需求 ID
	TCaseID     string // 必填：测试用例 ID
	TestPlanID  string // 必填：测试计划 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *DeleteTCaseStoryRelationRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"story_id":     r.StoryID,
		"tcase_id":     r.TCaseID,
		"test_plan_id": r.TestPlanID,
	}
}
