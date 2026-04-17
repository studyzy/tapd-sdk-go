package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetWorkspaceReports 获取项目报告
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/report/get_reports.html
func (c *Client) GetWorkspaceReports(ctx context.Context, req *model.WorkspaceIDRequest) (json.RawMessage, error) {
	data, err := c.doGet(ctx, "/reports", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}
