package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// AddLabel 创建标签
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/label/add_label.html
func (c *Client) AddLabel(ctx context.Context, req *model.AddLabelRequest) (*model.LabelPool, error) {
	data, err := c.doPost(ctx, "/label", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.LabelPool](data, "LabelPool")
}

// UpdateLabel 更新标签（不支持更新标签名称）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/label/update_label.html
func (c *Client) UpdateLabel(ctx context.Context, req *model.UpdateLabelRequest) (*model.LabelPool, error) {
	data, err := c.doPost(ctx, "/label", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.LabelPool](data, "LabelPool")
}

// QueryLabels 查询标签列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/label/query_label.html
func (c *Client) QueryLabels(ctx context.Context, req *model.QueryLabelRequest) ([]model.LabelPool, error) {
	data, err := c.doGet(ctx, "/label", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.LabelPool](data, "LabelPool")
}

// CountLabels 查询标签数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/label/count_label.html
func (c *Client) CountLabels(ctx context.Context, req *model.CountLabelRequest) (int, error) {
	data, err := c.doGet(ctx, "/label/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}
