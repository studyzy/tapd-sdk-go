package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetRelatedBugs 获取需求关联的缺陷列表，返回强类型 []model.StoryBugRelation
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_related_bugs.html
func (c *Client) GetRelatedBugs(ctx context.Context, req *model.GetRelatedBugsRequest) ([]model.StoryBugRelation, error) {
	data, err := c.doGet(ctx, "/stories/get_related_bugs", req.ToParams())
	if err != nil {
		return nil, err
	}

	var relations []model.StoryBugRelation
	if err := json.Unmarshal(data, &relations); err != nil {
		return nil, fmt.Errorf("failed to parse related bugs: %w", err)
	}
	return relations, nil
}

// CreateRelation 创建实体关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/create_story_bug.html
func (c *Client) CreateRelation(ctx context.Context, req *model.CreateRelationRequest) (*model.Relation, error) {
	data, err := c.doPost(ctx, "/relations", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Relation](data, "Relation")
}
