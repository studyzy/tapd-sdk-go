package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateMiniItem 创建工作项，返回创建后的完整 MiniItem 对象
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/add_mini_item.html
func (c *Client) CreateMiniItem(ctx context.Context, req *model.CreateMiniItemRequest) (*model.MiniItem, error) {
	data, err := c.doPost(ctx, "/mini_items", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create response: %w", err)
	}

	raw, ok := wrapper["MiniItem"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var item model.MiniItem
	if err := json.Unmarshal(raw, &item); err != nil {
		return nil, fmt.Errorf("failed to parse created mini item: %w", err)
	}

	return &item, nil
}

// UpdateMiniItem 更新工作项
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/update_mini_item.html
func (c *Client) UpdateMiniItem(ctx context.Context, req *model.UpdateMiniItemRequest) (*model.MiniItem, error) {
	data, err := c.doPost(ctx, "/mini_items", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update response: %w", err)
	}

	raw, ok := wrapper["MiniItem"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var item model.MiniItem
	if err := json.Unmarshal(raw, &item); err != nil {
		return nil, fmt.Errorf("failed to parse updated mini item: %w", err)
	}

	return &item, nil
}

// ListMiniItems 查询工作项列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_items.html
func (c *Client) ListMiniItems(ctx context.Context, req *model.ListMiniItemsRequest) ([]model.MiniItem, error) {
	data, err := c.doGet(ctx, "/mini_items", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse mini item list: %w", err)
	}

	results := make([]model.MiniItem, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["MiniItem"]; ok {
			var miniItem model.MiniItem
			if err := json.Unmarshal(raw, &miniItem); err == nil {
				results = append(results, miniItem)
			}
		}
	}
	return results, nil
}

// GetMiniItem 获取单个工作项详情
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_items.html
func (c *Client) GetMiniItem(ctx context.Context, workspaceID, id string) (*model.MiniItem, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
		"id":           id,
	}

	data, err := c.doGet(ctx, "/mini_items", params)
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(rawList) == 0 {
		return nil, &TAPDError{ExitCode: 2, Message: fmt.Sprintf("mini item %s not found", id)}
	}

	raw, ok := rawList[0]["MiniItem"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var item model.MiniItem
	if err := json.Unmarshal(raw, &item); err != nil {
		return nil, fmt.Errorf("failed to parse mini item: %w", err)
	}

	return &item, nil
}

// CountMiniItems 查询工作项数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_count.html
func (c *Client) CountMiniItems(ctx context.Context, req *model.CountMiniItemsRequest) (int, error) {
	data, err := c.doGet(ctx, "/mini_items/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}

// CreateMiniItemCategory 创建工作项分组
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/add_mini_item_category.html
func (c *Client) CreateMiniItemCategory(ctx context.Context, req *model.CreateMiniItemCategoryRequest) (*model.Category, error) {
	data, err := c.doPost(ctx, "/mini_item_categories", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create response: %w", err)
	}

	raw, ok := wrapper["Category"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var category model.Category
	if err := json.Unmarshal(raw, &category); err != nil {
		return nil, fmt.Errorf("failed to parse created category: %w", err)
	}

	return &category, nil
}

// UpdateMiniItemCategory 更新工作项分组
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/update_mini_item_category.html
func (c *Client) UpdateMiniItemCategory(ctx context.Context, req *model.UpdateMiniItemCategoryRequest) (*model.Category, error) {
	data, err := c.doPost(ctx, "/mini_item_categories", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update response: %w", err)
	}

	raw, ok := wrapper["Category"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var category model.Category
	if err := json.Unmarshal(raw, &category); err != nil {
		return nil, fmt.Errorf("failed to parse updated category: %w", err)
	}

	return &category, nil
}

// ListMiniItemCategories 查询工作项分组列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_categories.html
func (c *Client) ListMiniItemCategories(ctx context.Context, req *model.ListMiniItemCategoriesRequest) ([]model.Category, error) {
	data, err := c.doGet(ctx, "/mini_item_categories", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse category list: %w", err)
	}

	results := make([]model.Category, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["Category"]; ok {
			var category model.Category
			if err := json.Unmarshal(raw, &category); err == nil {
				results = append(results, category)
			}
		}
	}
	return results, nil
}

// CountMiniItemCategories 查询工作项分组数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_categories_count.html
func (c *Client) CountMiniItemCategories(ctx context.Context, req *model.CountMiniItemCategoriesRequest) (int, error) {
	data, err := c.doGet(ctx, "/mini_item_categories/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}

// GetMiniItemChanges 查询工作项动态列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_changes.html
func (c *Client) GetMiniItemChanges(ctx context.Context, req *model.GetMiniItemChangesRequest) ([]model.WorkitemChange, error) {
	data, err := c.doGet(ctx, "/mini_item_changes", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse mini item changes list: %w", err)
	}

	results := make([]model.WorkitemChange, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["WorkitemChange"]; ok {
			var change model.WorkitemChange
			if err := json.Unmarshal(raw, &change); err == nil {
				results = append(results, change)
			}
		}
	}
	return results, nil
}

// CountMiniItemChanges 查询工作项动态数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_changes_count.html
func (c *Client) CountMiniItemChanges(ctx context.Context, req *model.CountMiniItemChangesRequest) (int, error) {
	data, err := c.doGet(ctx, "/mini_item_changes/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}

	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, nil
}

// GetMiniItemCustomFields 获取工作项自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_custom_fields_settings.html
func (c *Client) GetMiniItemCustomFields(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.CustomFieldConfig, error) {
	data, err := c.doGet(ctx, "/mini_items/custom_fields_settings", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse custom fields: %w", err)
	}

	results := make([]model.CustomFieldConfig, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["CustomFieldConfig"]; ok {
			var cfg model.CustomFieldConfig
			if err := json.Unmarshal(raw, &cfg); err == nil {
				results = append(results, cfg)
			}
		}
	}
	return results, nil
}

// GetMiniItemFieldsLabel 获取工作项所有字段的中英文名
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_fields_label.html
func (c *Client) GetMiniItemFieldsLabel(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/mini_items/get_fields_label", req.ToParams())
	if err != nil {
		return nil, err
	}

	var labels map[string]string
	if err := json.Unmarshal(data, &labels); err != nil {
		return nil, fmt.Errorf("failed to parse mini item fields label: %w", err)
	}
	return labels, nil
}

// CreateMiniItemRelation 添加工作项与其他业务对象的关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/create_mini_item_relation.html
func (c *Client) CreateMiniItemRelation(ctx context.Context, req *model.CreateMiniItemRelationRequest) (bool, error) {
	data, err := c.doPost(ctx, "/mini_items/create_mini_item_relation", req.ToParams())
	if err != nil {
		return false, err
	}

	var result map[string]bool
	if err := json.Unmarshal(data, &result); err != nil {
		return false, fmt.Errorf("failed to parse relation response: %w", err)
	}

	return result["success"], nil
}

// GetMiniItemLinkedStories 获取工作项关联的需求列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_link_stories.html
func (c *Client) GetMiniItemLinkedStories(ctx context.Context, req *model.GetMiniItemLinkedStoriesRequest) ([]model.MiniItemRelation, error) {
	data, err := c.doGet(ctx, "/mini_items/get_link_stories", req.ToParams())
	if err != nil {
		return nil, err
	}

	var relations []model.MiniItemRelation
	if err := json.Unmarshal(data, &relations); err != nil {
		return nil, fmt.Errorf("failed to parse linked stories: %w", err)
	}
	return relations, nil
}

// GetMiniItemRelatedBugs 获取工作项关联的缺陷列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_related_bugs.html
func (c *Client) GetMiniItemRelatedBugs(ctx context.Context, req *model.GetMiniItemRelatedBugsRequest) ([]model.MiniItemRelation, error) {
	data, err := c.doGet(ctx, "/mini_items/get_related_bugs", req.ToParams())
	if err != nil {
		return nil, err
	}

	var relations []model.MiniItemRelation
	if err := json.Unmarshal(data, &relations); err != nil {
		return nil, fmt.Errorf("failed to parse related bugs: %w", err)
	}
	return relations, nil
}

// RemoveMiniItemRelation 解除工作项与其他业务对象的关联关系
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/remove_mini_item_relation.html
func (c *Client) RemoveMiniItemRelation(ctx context.Context, req *model.RemoveMiniItemRelationRequest) (bool, error) {
	data, err := c.doPost(ctx, "/mini_items/remove_mini_item_relation", req.ToParams())
	if err != nil {
		return false, err
	}

	var result map[string]bool
	if err := json.Unmarshal(data, &result); err != nil {
		return false, fmt.Errorf("failed to parse relation response: %w", err)
	}

	return result["success"], nil
}

// GetRemovedMiniItems 获取回收站内的工作项
// API 文档：https://open.tapd.cn/document/api-doc/API文档/mini_api_reference/mini_item/get_removed_mini_items.html
func (c *Client) GetRemovedMiniItems(ctx context.Context, req *model.GetRemovedMiniItemsRequest) ([]model.RemovedMiniItem, error) {
	data, err := c.doGet(ctx, "/mini_items/get_removed_mini_items", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse removed mini items: %w", err)
	}

	results := make([]model.RemovedMiniItem, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["RemovedMiniItem"]; ok {
			var removed model.RemovedMiniItem
			if err := json.Unmarshal(raw, &removed); err == nil {
				results = append(results, removed)
			}
		}
	}
	return results, nil
}
