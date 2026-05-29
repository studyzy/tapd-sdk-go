package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateStoryCategory 创建需求分类
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/add_story_category.html
func (c *Client) CreateStoryCategory(ctx context.Context, req *model.CreateStoryCategoryRequest) (*model.Category, error) {
	data, err := c.doPost(ctx, "/story_categories", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Category](data, "Category")
}

// UpdateStoryCategory 更新需求分类
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/update_story_category.html
func (c *Client) UpdateStoryCategory(ctx context.Context, req *model.UpdateStoryCategoryRequest) (*model.Category, error) {
	data, err := c.doPost(ctx, "/story_categories", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Category](data, "Category")
}

// CountStoryCategories 查询需求分类数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_categories_count.html
func (c *Client) CountStoryCategories(ctx context.Context, req *model.CountStoryCategoriesRequest) (int, error) {
	data, err := c.doGet(ctx, "/story_categories/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}

// AddStoryLinkRelations 创建需求关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/add_story_link_relations.html
func (c *Client) AddStoryLinkRelations(ctx context.Context, req *model.AddStoryLinkRelationsRequest) (bool, error) {
	data, err := c.doPost(ctx, "/stories/add_story_link_relations", req.ToParams())
	if err != nil {
		return false, err
	}
	var r struct {
		Success int `json:"success"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return false, fmt.Errorf("failed to parse add_story_link_relations response: %w", err)
	}
	return r.Success == 1, nil
}

// GetLinkStories 获取需求与其它需求的所有关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_link_stories.html
func (c *Client) GetLinkStories(ctx context.Context, req *model.GetLinkStoriesRequest) ([]model.StoryLinkRelation, error) {
	data, err := c.doGet(ctx, "/stories/get_link_stories", req.ToParams())
	if err != nil {
		return nil, err
	}
	var rel []model.StoryLinkRelation
	if err := json.Unmarshal(data, &rel); err != nil {
		return nil, fmt.Errorf("failed to parse link stories: %w", err)
	}
	return rel, nil
}

// AddStoryTcase 创建需求与测试用例关联关系，返回成功关联的测试用例 ID 列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/create_story_tcase.html
func (c *Client) AddStoryTcase(ctx context.Context, req *model.AddStoryTcaseRequest) ([]string, error) {
	data, err := c.doPost(ctx, "/stories/add_story_tcase", req.ToParams())
	if err != nil {
		return nil, err
	}
	var r struct {
		SuccessID []string `json:"success_id"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("failed to parse add_story_tcase response: %w", err)
	}
	return r.SuccessID, nil
}

// GetStoryTcase 获取需求与测试用例关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_tcase.html
func (c *Client) GetStoryTcase(ctx context.Context, req *model.GetStoryTcaseRequest) ([]model.TestPlanStoryTcaseRelation, error) {
	data, err := c.doGet(ctx, "/stories/get_story_tcase", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.TestPlanStoryTcaseRelation](data, "TestPlanStoryTcaseRelation")
}

// RemoveStoryBugRelations 解除需求与缺陷的关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/remove_story_bug_raletions.html
func (c *Client) RemoveStoryBugRelations(ctx context.Context, req *model.RemoveStoryBugRelationsRequest) (bool, error) {
	data, err := c.doPost(ctx, "/stories/remove_story_bug_raletions", req.ToParams())
	if err != nil {
		return false, err
	}
	var r struct {
		Success bool `json:"success"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return false, fmt.Errorf("failed to parse remove_story_bug_raletions response: %w", err)
	}
	return r.Success, nil
}

// GetTimeRelativeStories 获取需求前后置关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_time_relative_stories.html
func (c *Client) GetTimeRelativeStories(ctx context.Context, req *model.GetTimeRelativeStoriesRequest) ([]model.WorkitemTimeRelation, error) {
	data, err := c.doGet(ctx, "/stories/get_time_relative_stories", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTimeRelation](data, "WorkitemTimeRelation")
}

// SaveTimeRelations 批量新增或修改需求前后置关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/save_time_relations.html
func (c *Client) SaveTimeRelations(ctx context.Context, req *model.SaveTimeRelationsRequest) (bool, error) {
	data, err := c.doPost(ctx, "/stories/save_time_relations", req.ToParams())
	if err != nil {
		return false, err
	}
	var r struct {
		Result bool `json:"result"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return false, fmt.Errorf("failed to parse save_time_relations response: %w", err)
	}
	return r.Result, nil
}

// DeleteTimeRelations 批量删除需求前后置关系，返回实际删除条数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/delete_time_relations.html
func (c *Client) DeleteTimeRelations(ctx context.Context, req *model.DeleteTimeRelationsRequest) (int, error) {
	data, err := c.doPost(ctx, "/stories/delete_time_relations", req.ToParams())
	if err != nil {
		return 0, err
	}
	var r struct {
		Num int `json:"num"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return 0, fmt.Errorf("failed to parse delete_time_relations response: %w", err)
	}
	return r.Num, nil
}

// GetStoryTemplateList 获取需求模板列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_template_list.html
func (c *Client) GetStoryTemplateList(ctx context.Context, req *model.GetStoryTemplateListRequest) ([]model.WorkitemTemplate, error) {
	data, err := c.doGet(ctx, "/stories/template_list", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplate](data, "WorkitemTemplate")
}

// GetDefaultStoryTemplate 获取需求模板字段
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_default_story_template.html
func (c *Client) GetDefaultStoryTemplate(ctx context.Context, req *model.GetDefaultStoryTemplateRequest) ([]model.WorkitemTemplateField, error) {
	data, err := c.doGet(ctx, "/stories/get_default_story_template", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplateField](data, "WorkitemTemplateField")
}

// GetStorySteps 使用并行工作流的需求，获取其节点信息
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_steps.html
func (c *Client) GetStorySteps(ctx context.Context, req *model.GetStoryStepsRequest) ([]model.WorkitemStepInfo, error) {
	data, err := c.doGet(ctx, "/stories/get_story_step_list", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemStepInfo](data, "WorkitemStepInfo")
}

// GetRemovedStories 获取回收站中的需求
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_removed_stories.html
func (c *Client) GetRemovedStories(ctx context.Context, req *model.GetRemovedStoriesRequest) ([]model.RemovedStory, error) {
	data, err := c.doGet(ctx, "/stories/get_removed_stories", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.RemovedStory](data, "RemovedStory")
}

// CopyStory 复制需求
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/copy_story.html
func (c *Client) CopyStory(ctx context.Context, req *model.CopyStoryRequest) (*model.Story, error) {
	data, err := c.doPost(ctx, "/stories/copy_story", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Story](data, "Story")
}

// UpdateStoryParent 更新父需求
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/update_story_parent.html
func (c *Client) UpdateStoryParent(ctx context.Context, req *model.UpdateStoryParentRequest) (*model.Story, error) {
	data, err := c.doPost(ctx, "/stories/update_story_parent", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Story](data, "Story")
}

// ChangeWorkitemType 更新需求的需求类别。
// 该接口返回 data 为非 Story 包裹的扁平对象，且字段非常多（含 markdown_description 等），
// 因此返回 json.RawMessage 由调用方按需解析。
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/change_workitem_type.html
func (c *Client) ChangeWorkitemType(ctx context.Context, req *model.ChangeWorkitemTypeRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/stories/change_workitem_type", req.ToParams())
}

// GetStoriesByViewConfID 获取视图对应的需求列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_stories_by_view_conf_id.html
func (c *Client) GetStoriesByViewConfID(ctx context.Context, req *model.GetStoriesByViewConfIDRequest) ([]model.Story, error) {
	data, err := c.doGet(ctx, "/stories/get_stories_by_view_conf_id", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.Story](data, "Story")
}

// BatchUpdateStory 批量更新需求，返回 API 提示信息（如 "batch update success"）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/batch_update_story.html
func (c *Client) BatchUpdateStory(ctx context.Context, req *model.BatchUpdateStoryRequest) (string, error) {
	data, err := c.doPost(ctx, "/stories/batch_update_story", req.ToParams())
	if err != nil {
		return "", err
	}
	var r struct {
		Msg string `json:"msg"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return "", fmt.Errorf("failed to parse batch_update_story response: %w", err)
	}
	return r.Msg, nil
}

// GetSecretInfo 获取需求保密信息
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_secret_info.html
func (c *Client) GetSecretInfo(ctx context.Context, req *model.GetSecretInfoRequest) (*model.SecretInfo, error) {
	data, err := c.doGet(ctx, "/stories/get_secret_info", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result model.SecretInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse secret info: %w", err)
	}
	return &result, nil
}

// GetSecretStories 获取保密需求 ID 列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_secret_stories.html
func (c *Client) GetSecretStories(ctx context.Context, req *model.GetSecretStoriesRequest) (*model.SecretStoriesList, error) {
	data, err := c.doGet(ctx, "/secret_stories", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result model.SecretStoriesList
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse secret stories: %w", err)
	}
	return &result, nil
}

// CountSecretStories 获取保密需求数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_secret_stories_count.html
func (c *Client) CountSecretStories(ctx context.Context, workspaceID string) (int, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
	}
	data, err := c.doGet(ctx, "/secret_stories/count", params)
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}

// BatchUpdateSecretInfo 批量修改需求的保密信息
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/batch_update_secret_info.html
func (c *Client) BatchUpdateSecretInfo(ctx context.Context, req *model.BatchUpdateSecretInfoRequest) (string, error) {
	data, err := c.doPost(ctx, "/stories/batch_update_secret_info", req.ToParams())
	if err != nil {
		return "", err
	}
	var r struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return "", fmt.Errorf("failed to parse batch_update_secret_info response: %w", err)
	}
	return r.Msg, nil
}

// StoryFilterToQueryToken 将需求过滤条件转换成 QueryToken
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/filter_to_query_token.html
func (c *Client) StoryFilterToQueryToken(ctx context.Context, req *model.FilterToQueryTokenRequest) (*model.QueryTokenResponse, error) {
	data, err := c.doPost(ctx, "/stories/filter_to_query_token", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result model.QueryTokenResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse filter_to_query_token response: %w", err)
	}
	return &result, nil
}

// StoryIDsToQueryToken 将需求 ID 列表转换成 QueryToken
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/story_ids_to_query_token.html
func (c *Client) StoryIDsToQueryToken(ctx context.Context, req *model.IDsToQueryTokenRequest) (*model.QueryTokenResponse, error) {
	data, err := c.doPost(ctx, "/stories/ids_to_query_token", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result model.QueryTokenResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse ids_to_query_token response: %w", err)
	}
	return &result, nil
}

// RemoveStoryLinkRelation 解除需求关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/remove_story_link_relation.html
func (c *Client) RemoveStoryLinkRelation(ctx context.Context, req *model.RemoveStoryLinkRelationRequest) (bool, error) {
	data, err := c.doPost(ctx, "/stories/remove_story_link_relation", req.ToParams())
	if err != nil {
		return false, err
	}
	var r struct {
		Success int `json:"success"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return false, fmt.Errorf("failed to parse remove_story_link_relation response: %w", err)
	}
	return r.Success == 1, nil
}

// ResetWorkitemSteps 并行工作流流程重置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/reset_workitem_steps.html
func (c *Client) ResetWorkitemSteps(ctx context.Context, req *model.ResetWorkitemStepsRequest) (*model.Story, error) {
	data, err := c.doPost(ctx, "/stories/reset_workitem_steps", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Story](data, "Story")
}

// UpdateStoryStepStatus 完成进行中的节点（并行工作流）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/update_story_step_status.html
func (c *Client) UpdateStoryStepStatus(ctx context.Context, req *model.UpdateStoryStepStatusRequest) (*model.Story, error) {
	data, err := c.doPost(ctx, "/stories/update_story_step_status", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Story](data, "Story")
}

// CountStoriesByCategories 获取指定分类下需求数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/count_by_categories.html
func (c *Client) CountStoriesByCategories(ctx context.Context, req *model.CountStoriesByCategoriesRequest) (map[string]int, error) {
	data, err := c.doGet(ctx, "/stories/count_by_categories", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse count_by_categories response: %w", err)
	}
	return result, nil
}
