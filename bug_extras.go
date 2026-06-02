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

// BugFilterToQueryToken 将缺陷过滤条件转换成 QueryToken
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/filter_to_query_token.html
func (c *Client) BugFilterToQueryToken(ctx context.Context, req *model.FilterToQueryTokenRequest) (*model.QueryTokenResponse, error) {
	data, err := c.doPost(ctx, "/bugs/filter_to_query_token", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result model.QueryTokenResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse bug filter_to_query_token response: %w", err)
	}
	return &result, nil
}

// CopyBug 复制缺陷
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/copy_bug.html
func (c *Client) CopyBug(ctx context.Context, req *model.CopyBugRequest) (*model.Bug, error) {
	data, err := c.doPost(ctx, "/bugs/copy_bug", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.Bug](data, "Bug")
}

// BatchUpdateBug 批量更新缺陷
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/batch_update_bug.html
func (c *Client) BatchUpdateBug(ctx context.Context, req *model.BatchUpdateBugRequest) error {
	_, err := c.doPost(ctx, "/bugs/batch_update_bug", req.ToParams())
	return err
}

// GetRemovedBugs 获取回收站中的缺陷
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_removed_bugs.html
func (c *Client) GetRemovedBugs(ctx context.Context, req *model.GetRemovedBugsRequest) ([]model.RemovedBug, error) {
	data, err := c.doGet(ctx, "/bugs/get_removed_bugs", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.RemovedBug](data, "RemovedBug")
}

// GetBugsByViewConfID 获取视图对应的缺陷列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/get_bugs_by_view_conf_id.html
func (c *Client) GetBugsByViewConfID(ctx context.Context, req *model.GetBugsByViewConfIDRequest) ([]model.Bug, error) {
	data, err := c.doGet(ctx, "/bugs/get_bugs_by_view_conf_id", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.Bug](data, "Bug")
}

// BugIDsToQueryToken 将缺陷 ID 列表转换成 QueryToken
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/ids_to_query_token.html
func (c *Client) BugIDsToQueryToken(ctx context.Context, req *model.BugIDsToQueryTokenRequest) (*model.QueryTokenResponse, error) {
	data, err := c.doPost(ctx, "/bugs/ids_to_query_token", req.ToParams())
	if err != nil {
		return nil, err
	}
	var result model.QueryTokenResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse bug ids_to_query_token response: %w", err)
	}
	return &result, nil
}

// UpdateBugSystemSelectFieldOptions 更新缺陷系统下拉字段选项
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/bug/update_system_select_field_options.html
func (c *Client) UpdateBugSystemSelectFieldOptions(ctx context.Context, req *model.UpdateBugSystemSelectFieldOptionsRequest) error {
	_, err := c.doPost(ctx, "/bugs/update_system_select_field_options", req.ToParams())
	return err
}
