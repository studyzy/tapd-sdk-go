package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetWorkspaceReports 获取项目报告
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/report/get_workspace_reports.html
func (c *Client) GetWorkspaceReports(ctx context.Context, req *model.GetWorkspaceReportsRequest) ([]model.WorkspaceReport, error) {
	data, err := c.doGet(ctx, "/workspace_reports", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.WorkspaceReport](data, "WorkspaceReport")
}
