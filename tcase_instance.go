package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// AssignTCaseInstance 分配测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/assign_tcase_instance.html
func (c *Client) AssignTCaseInstance(ctx context.Context, req *model.AssignTCaseInstanceRequest) error {
	_, err := c.doPost(ctx, "/tcase_instance/assign", req.ToParams())
	return err
}

// ExecuteTCaseInstance 执行测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/execute_tcase_instance.html
func (c *Client) ExecuteTCaseInstance(ctx context.Context, req *model.ExecuteTCaseInstanceRequest) error {
	_, err := c.doPost(ctx, "/tcase_instance/execute", req.ToParams())
	return err
}

// RemoveTCaseInstance 测试用例移出测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/remove_tcase_instance.html
func (c *Client) RemoveTCaseInstance(ctx context.Context, req *model.RemoveTCaseInstanceRequest) error {
	_, err := c.doPost(ctx, "/tcase_instance/remove_tcase", req.ToParams())
	return err
}

// GetTCaseResult 获取测试用例执行结果
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_result.html
func (c *Client) GetTCaseResult(ctx context.Context, req *model.TCaseResultRequest) (map[string]model.TCaseResultItem, error) {
	data, err := c.doGet(ctx, "/tcase_instance/result", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result map[string]model.TCaseResultItem
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse tcase result: %w", err)
	}
	return result, nil
}

// DeleteTCaseStoryRelation 解除测试用例关联并移出测试计划
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/delete_tcase_story_relation.html
func (c *Client) DeleteTCaseStoryRelation(ctx context.Context, req *model.DeleteTCaseStoryRelationRequest) error {
	_, err := c.doPost(ctx, "/tcase_instance/delete_tcase_story_relation", req.ToParams())
	return err
}
