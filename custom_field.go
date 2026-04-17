package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetCustomFields 获取自定义字段配置，返回强类型 []model.CustomFieldConfig
// entityType 取值：stories, tasks, iterations, tcases
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_custom_fields_settings.html
func (c *Client) GetCustomFields(ctx context.Context, req *model.GetCustomFieldsRequest) ([]model.CustomFieldConfig, error) {
	validTypes := map[string]bool{"stories": true, "tasks": true, "iterations": true, "tcases": true}
	if !validTypes[req.EntityType] {
		return nil, fmt.Errorf("invalid entity type: %s, must be one of stories/tasks/iterations/tcases", req.EntityType)
	}
	data, err := c.doGet(ctx, "/"+req.EntityType+"/custom_fields_settings", req.ToParams())
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

// GetStoryFieldsLabel 获取需求所有字段的中英文名，返回 map[string]string（字段英文名→中文名）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_lable.html
func (c *Client) GetStoryFieldsLabel(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/stories/get_fields_lable", req.ToParams())
	if err != nil {
		return nil, err
	}

	var labels map[string]string
	if err := json.Unmarshal(data, &labels); err != nil {
		return nil, fmt.Errorf("failed to parse story fields label: %w", err)
	}
	return labels, nil
}

// GetStoryFieldsInfo 获取需求所有字段及候选值，返回 map[string]model.FieldInfo
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_info.html
func (c *Client) GetStoryFieldsInfo(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]model.FieldInfo, error) {
	data, err := c.doGet(ctx, "/stories/get_fields_info", req.ToParams())
	if err != nil {
		return nil, err
	}

	var fields map[string]model.FieldInfo
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, fmt.Errorf("failed to parse story fields info: %w", err)
	}
	return fields, nil
}

// GetBugFieldsLabel 获取缺陷所有字段的中英文名，返回 map[string]string（字段英文名→中文名）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_bug_fields_lable.html
func (c *Client) GetBugFieldsLabel(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]string, error) {
	data, err := c.doGet(ctx, "/bugs/get_fields_lable", req.ToParams())
	if err != nil {
		return nil, err
	}

	var labels map[string]string
	if err := json.Unmarshal(data, &labels); err != nil {
		return nil, fmt.Errorf("failed to parse bug fields label: %w", err)
	}
	return labels, nil
}

// GetBugFieldsInfo 获取缺陷所有字段及候选值，返回 map[string]model.FieldInfo
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_bug_fields_info.html
func (c *Client) GetBugFieldsInfo(ctx context.Context, req *model.WorkspaceIDRequest) (map[string]model.FieldInfo, error) {
	data, err := c.doGet(ctx, "/bugs/get_fields_info", req.ToParams())
	if err != nil {
		return nil, err
	}

	var fields map[string]model.FieldInfo
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, fmt.Errorf("failed to parse bug fields info: %w", err)
	}
	return fields, nil
}

// GetWorkitemTypes 获取需求类别列表，返回强类型 []model.WorkitemType
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_workitem_types.html
func (c *Client) GetWorkitemTypes(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.WorkitemType, error) {
	data, err := c.doGet(ctx, "/workitem_types", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse workitem types: %w", err)
	}

	results := make([]model.WorkitemType, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["WorkitemType"]; ok {
			var wt model.WorkitemType
			if err := json.Unmarshal(raw, &wt); err == nil {
				results = append(results, wt)
			}
		}
	}
	return results, nil
}
