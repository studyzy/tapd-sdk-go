package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// AssignTCaseInstance 分配测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/assign_tcase_instance.html
func (c *Client) AssignTCaseInstance(ctx context.Context, req *model.AssignTCaseInstanceRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/tcase_instance/assign", req.ToParams())
}

// ExecuteTCaseInstance 执行测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/execute_tcase_instance.html
func (c *Client) ExecuteTCaseInstance(ctx context.Context, req *model.ExecuteTCaseInstanceRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/tcase_instance/execute", req.ToParams())
}

// RemoveTCaseInstance 测试用例移出测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/remove_tcase_instance.html
func (c *Client) RemoveTCaseInstance(ctx context.Context, req *model.RemoveTCaseInstanceRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/tcase_instance/remove_tcase", req.ToParams())
}

// GetTCaseResult 获取测试用例执行结果
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_result.html
func (c *Client) GetTCaseResult(ctx context.Context, req *model.TCaseResultRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/tcase_instance/result", req.ToParams())
}

// DeleteTCaseStoryRelation 解除测试用例关联并移出测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/delete_tcase_story_relation.html
func (c *Client) DeleteTCaseStoryRelation(ctx context.Context, req *model.DeleteTCaseStoryRelationRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/tcase_instance/delete_tcase_story_relation", req.ToParams())
}
