package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListComments 查询评论列表，返回强类型 Comment 切片
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/get_comments.html
func (c *Client) ListComments(ctx context.Context, req *model.ListCommentsRequest) ([]model.Comment, error) {
	data, err := c.doGet(ctx, "/comments", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Comment](data, "Comment")
}

// AddComment 添加评论，返回新建的评论对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/add_comment.html
func (c *Client) AddComment(ctx context.Context, req *model.AddCommentRequest) (*model.Comment, error) {
	data, err := c.doPost(ctx, "/comments", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Comment](data, "Comment")
}

// UpdateComment 更新评论，返回更新后的评论对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/update_comment.html
func (c *Client) UpdateComment(ctx context.Context, req *model.UpdateCommentRequest) (*model.Comment, error) {
	data, err := c.doPost(ctx, "/comments", req.ToParams())
	if err != nil {
		return nil, err
	}

	// 尝试直接解析（文档示例格式：data 直接是评论对象）
	var comment model.Comment
	if err := json.Unmarshal(data, &comment); err == nil && comment.ID != "" {
		return &comment, nil
	}

	// 兼容：可能有 "Comment" 包裹层
	return parseOne[model.Comment](data, "Comment")
}

// CountComments 查询评论数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/get_comments_count.html
func (c *Client) CountComments(ctx context.Context, req *model.CountCommentsRequest) (int, error) {
	data, err := c.doGet(ctx, "/comments/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}
