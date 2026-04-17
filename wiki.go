package tapd

import (
	"context"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListWikis 查询 Wiki 文档列表，返回强类型 []model.Wiki
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis.html
func (c *Client) ListWikis(ctx context.Context, req *model.ListWikisRequest) ([]model.Wiki, error) {
	data, err := c.doGet(ctx, "/tapd_wikis", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Wiki](data, "Wiki")
}

// GetWiki 获取单个 Wiki 文档详情，description 字段保留原始 HTML
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis.html
func (c *Client) GetWiki(ctx context.Context, workspaceID, id string) (*model.Wiki, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
		"id":           id,
	}

	data, err := c.doGet(ctx, "/tapd_wikis", params)
	if err != nil {
		return nil, err
	}

	wikis, err := parseList[model.Wiki](data, "Wiki")
	if err != nil {
		return nil, err
	}

	if len(wikis) == 0 {
		return nil, &TAPDError{ExitCode: 2, Message: fmt.Sprintf("wiki %s not found", id)}
	}

	wiki := wikis[0]
	wiki.URL = fmt.Sprintf("%s/%s/markdown_wikis/view/%s", c.webURL, workspaceID, id)

	return &wiki, nil
}

// CreateWiki 创建 Wiki 文档，返回创建后的完整 Wiki 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/add_tapd_wiki.html
func (c *Client) CreateWiki(ctx context.Context, req *model.CreateWikiRequest) (*model.Wiki, error) {
	data, err := c.doPost(ctx, "/tapd_wikis", req.ToParams())
	if err != nil {
		return nil, err
	}

	wiki, err := parseOne[model.Wiki](data, "Wiki")
	if err != nil {
		return nil, err
	}

	wiki.URL = fmt.Sprintf("%s/%s/markdown_wikis/view/%s", c.webURL, req.WorkspaceID, wiki.ID)

	return wiki, nil
}

// CountWikis 获取 Wiki 数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_count.html
func (c *Client) CountWikis(ctx context.Context, req *model.CountWikisRequest) (int, error) {
	data, err := c.doGet(ctx, "/tapd_wikis/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// UpdateWiki 更新 Wiki 文档
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/update_tapd_wiki.html
func (c *Client) UpdateWiki(ctx context.Context, req *model.UpdateWikiRequest) (*model.Wiki, error) {
	data, err := c.doPost(ctx, "/tapd_wikis", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Wiki](data, "Wiki")
}
