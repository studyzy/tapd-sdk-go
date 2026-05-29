package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// AddThirdRelation 创建 TAPD 业务对象与流水线构建记录关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/pipeline/add_third_relation.html
func (c *Client) AddThirdRelation(ctx context.Context, req *model.AddThirdRelationRequest) (*model.ThirdRelation, error) {
	data, err := c.doPost(ctx, "/third_relations", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.ThirdRelation](data, "ThirdRelations")
}

// GetThirdRelations 获取 TAPD 业务对象与构建记录的关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/pipeline/get_third_relations.html
func (c *Client) GetThirdRelations(ctx context.Context, req *model.GetThirdRelationsRequest) ([]model.ThirdRelation, error) {
	data, err := c.doGet(ctx, "/third_relations", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.ThirdRelation](data, "ThirdRelations")
}

// DeleteThirdRelation 解除指定构建记录与业务对象关联
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/pipeline/delete_third_relations.html
func (c *Client) DeleteThirdRelation(ctx context.Context, req *model.DeleteThirdRelationRequest) (bool, error) {
	data, err := c.doDelete(ctx, "/third_relations", req.ToParams())
	if err != nil {
		return false, err
	}
	var r struct {
		Result bool `json:"result"`
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return false, fmt.Errorf("failed to parse delete third relation response: %w", err)
	}
	return r.Result, nil
}
