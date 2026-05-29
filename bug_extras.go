// Package tapd 中的 bug_extras.go 实现缺陷关联、模板列表/字段、自定义字段配置等扩展接口
package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// LinkBugs 关联缺陷
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/link_bugs.html
func (c *Client) LinkBugs(ctx context.Context, req *model.LinkBugsRequest) error {
	_, err := c.doPost(ctx, "/bugs/link_bugs", req.ToParams())
	return err
}

// DeleteLinkBugs 取消关联缺陷
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/delete_link_bugs.html
func (c *Client) DeleteLinkBugs(ctx context.Context, req *model.DeleteLinkBugsRequest) error {
	_, err := c.doPost(ctx, "/bugs/delete_link_bugs", req.ToParams())
	return err
}

// GetLinkBugs 获取缺陷与其它缺陷的所有关联关系（无分页）
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_link_bugs.html
func (c *Client) GetLinkBugs(ctx context.Context, req *model.GetLinkBugsRequest) ([]model.BugLink, error) {
	data, err := c.doGet(ctx, "/bugs/get_link_bugs", req.ToParams())
	if err != nil {
		return nil, err
	}
	var links []model.BugLink
	if err := json.Unmarshal(data, &links); err != nil {
		return nil, fmt.Errorf("failed to parse link bugs: %w", err)
	}
	return links, nil
}

// GetBugRelatedStories 获取缺陷关联的需求 ID
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_related_stories.html
func (c *Client) GetBugRelatedStories(ctx context.Context, req *model.GetBugRelatedStoriesRequest) ([]model.BugStoryRelation, error) {
	data, err := c.doGet(ctx, "/bugs/get_related_stories", req.ToParams())
	if err != nil {
		return nil, err
	}
	var rels []model.BugStoryRelation
	if err := json.Unmarshal(data, &rels); err != nil {
		return nil, fmt.Errorf("failed to parse bug related stories: %w", err)
	}
	return rels, nil
}

// ListBugTemplates 获取缺陷模板列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_template_list.html
func (c *Client) ListBugTemplates(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.WorkitemTemplate, error) {
	data, err := c.doGet(ctx, "/bugs/template_list", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplate](data, "WorkitemTemplate")
}

// GetDefaultBugTemplate 获取缺陷模板字段
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_default_bug_template.html
func (c *Client) GetDefaultBugTemplate(ctx context.Context, req *model.GetDefaultBugTemplateRequest) ([]model.WorkitemTemplateField, error) {
	data, err := c.doGet(ctx, "/bugs/get_default_bug_template", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkitemTemplateField](data, "WorkitemTemplateField")
}

// GetBugCustomFieldsSettings 获取缺陷自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_custom_fields_settings.html
func (c *Client) GetBugCustomFieldsSettings(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.CustomFieldConfig, error) {
	data, err := c.doGet(ctx, "/bugs/custom_fields_settings", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.CustomFieldConfig](data, "CustomFieldConfig")
}
