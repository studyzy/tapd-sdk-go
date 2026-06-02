package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListCategories 查询需求分类列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_categories.html
func (c *Client) ListCategories(ctx context.Context, req *model.ListStoryCategoriesRequest) ([]model.Category, error) {
	data, err := c.doGet(ctx, "/story_categories", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Category](data, "Category")
}
