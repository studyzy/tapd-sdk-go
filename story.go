package tapd

import (
	"context"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListStories 查询需求列表，返回强类型 Story 切片，自定义字段通过 CustomFields map 保留
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories.html
func (c *Client) ListStories(ctx context.Context, req *model.ListStoriesRequest) ([]model.Story, error) {
	data, err := c.doGet(ctx, "/stories", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Story](data, "Story")
}

// GetStory 获取单个需求详情，description 字段保留原始 HTML
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories.html
func (c *Client) GetStory(ctx context.Context, workspaceID, id string) (*model.Story, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
		"id":           id,
	}

	data, err := c.doGet(ctx, "/stories", params)
	if err != nil {
		return nil, err
	}

	stories, err := parseList[model.Story](data, "Story")
	if err != nil {
		return nil, err
	}

	if len(stories) == 0 {
		return nil, &TAPDError{ExitCode: 2, Message: fmt.Sprintf("story %s not found", id)}
	}

	story := stories[0]
	story.URL = fmt.Sprintf("%s/%s/prong/stories/view/%s", c.webURL, workspaceID, id)

	return &story, nil
}

// CreateStory 创建需求，返回创建后的完整 Story 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/add_story.html
func (c *Client) CreateStory(ctx context.Context, req *model.CreateStoryRequest) (*model.Story, error) {
	data, err := c.doPost(ctx, "/stories", req.ToParams())
	if err != nil {
		return nil, err
	}

	story, err := parseOne[model.Story](data, "Story")
	if err != nil {
		return nil, err
	}

	story.URL = fmt.Sprintf("%s/%s/prong/stories/view/%s", c.webURL, req.WorkspaceID, story.ID)

	return story, nil
}

// UpdateStory 更新需求
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/update_story.html
func (c *Client) UpdateStory(ctx context.Context, req *model.UpdateStoryRequest) (*model.Story, error) {
	data, err := c.doPost(ctx, "/stories", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Story](data, "Story")
}

// CountStories 查询需求数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories_count.html
func (c *Client) CountStories(ctx context.Context, req *model.CountStoriesRequest) (int, error) {
	data, err := c.doGet(ctx, "/stories/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}
