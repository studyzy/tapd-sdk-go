package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CountWikiAttachments 获取 Wiki 附件数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_attachments_count.html
func (c *Client) CountWikiAttachments(ctx context.Context, req *model.CountWikiAttachmentsRequest) (int, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_attachments/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}

// GetWikiDrawio 获取 Wiki 内嵌 drawio 数据
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_drawios.html
func (c *Client) GetWikiDrawio(ctx context.Context, req *model.GetWikiDrawioRequest) (*model.WikiDrawio, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_drawios", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseOne[model.WikiDrawio](data, "StaticData")
}

// ListWikiEntityPermissions 获取 Wiki 可访问范围（用户及用户组）
// 注：仅当 Wiki 的 is_private 为 1 才需要调用此接口
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_entity_permissions.html
func (c *Client) ListWikiEntityPermissions(ctx context.Context, req *model.ListWikiEntityPermissionsRequest) ([]model.WikiEntityPermission, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_entity_permissions", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WikiEntityPermission](data, "EntityPermission")
}

// ListWikiFollowers 获取 Wiki 关注人列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_followers.html
func (c *Client) ListWikiFollowers(ctx context.Context, req *model.ListWikiFollowersRequest) ([]model.WikiFollower, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_followers", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WikiFollower](data, "UserFollows")
}

// CountWikiFollowers 获取 Wiki 关注人数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_followers_count.html
func (c *Client) CountWikiFollowers(ctx context.Context, req *model.CountWikiFollowersRequest) (int, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_followers/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}

// ListWikiTags 获取 Wiki 标签信息列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_tags.html
func (c *Client) ListWikiTags(ctx context.Context, req *model.ListWikiTagsRequest) ([]model.WikiTag, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_tags", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WikiTag](data, "Tags")
}

// CountWikiTags 获取 Wiki 标签信息数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_tags_count.html
func (c *Client) CountWikiTags(ctx context.Context, req *model.CountWikiTagsRequest) (int, error) {
	data, err := c.doGet(ctx, "/tapd_wikis_tags/count", req.ToParams())
	if err != nil {
		return 0, err
	}
	return parseCount(data)
}
