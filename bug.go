package tapd

import (
	"context"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListBugs 查询缺陷列表，返回强类型 Bug 切片，自定义字段通过 CustomFields map 保留
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs.html
func (c *Client) ListBugs(ctx context.Context, req *model.ListBugsRequest) ([]model.Bug, error) {
	data, err := c.doGet(ctx, "/bugs", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Bug](data, "Bug")
}

// GetBug 获取单个缺陷详情，description 字段保留原始 HTML
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs.html
func (c *Client) GetBug(ctx context.Context, workspaceID, id string) (*model.Bug, error) {
	params := map[string]string{
		"workspace_id": workspaceID,
		"id":           id,
	}

	data, err := c.doGet(ctx, "/bugs", params)
	if err != nil {
		return nil, err
	}

	bugs, err := parseList[model.Bug](data, "Bug")
	if err != nil {
		return nil, err
	}

	if len(bugs) == 0 {
		return nil, &TAPDError{ExitCode: 2, Message: fmt.Sprintf("bug %s not found", id)}
	}

	bug := bugs[0]
	bug.URL = fmt.Sprintf("%s/%s/bugtrace/bugs/view/%s", c.webURL, workspaceID, id)

	return &bug, nil
}

// CreateBug 创建缺陷，返回创建后的完整 Bug 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/add_bug.html
func (c *Client) CreateBug(ctx context.Context, req *model.CreateBugRequest) (*model.Bug, error) {
	data, err := c.doPost(ctx, "/bugs", req.ToParams())
	if err != nil {
		return nil, err
	}

	bug, err := parseOne[model.Bug](data, "Bug")
	if err != nil {
		return nil, err
	}

	bug.URL = fmt.Sprintf("%s/%s/bugtrace/bugs/view/%s", c.webURL, req.WorkspaceID, bug.ID)

	return bug, nil
}

// UpdateBug 更新缺陷
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/update_bug.html
func (c *Client) UpdateBug(ctx context.Context, req *model.UpdateBugRequest) (*model.Bug, error) {
	data, err := c.doPost(ctx, "/bugs", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Bug](data, "Bug")
}

// CountBugs 查询缺陷数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs_count.html
func (c *Client) CountBugs(ctx context.Context, req *model.CountBugsRequest) (int, error) {
	data, err := c.doGet(ctx, "/bugs/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}
