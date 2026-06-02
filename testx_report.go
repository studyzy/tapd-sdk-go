package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/studyzy/tapd-sdk-go/model"
)

// TestxListReports 获取报告列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/report/list_report.html
func (c *Client) TestxListReports(ctx context.Context, req *model.TestxListReportsRequest) ([]model.TestxReport, int, error) {
	path := fmt.Sprintf("/api/testx/report/v1/namespaces/%s/reports", req.Namespace)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	if req.Search != "" {
		q.Set("Search", req.Search)
	}
	if req.StartAt != "" {
		q.Set("StartAt", req.StartAt)
	}
	if req.EndAt != "" {
		q.Set("EndAt", req.EndAt)
	}
	if len(req.Creators) > 0 {
		q.Set("Creators", strings.Join(req.Creators, ","))
	}
	if len(req.PlanUids) > 0 {
		q.Set("PlanUids", strings.Join(req.PlanUids, ","))
	}
	if req.WithAssociated {
		q.Set("WithAssociated", "true")
	}
	if req.TemplateUid != "" {
		q.Set("TemplateUid", req.TemplateUid)
	}
	if req.Source != "" {
		q.Set("Source", req.Source)
	}
	if len(req.Sources) > 0 {
		q.Set("Sources", strings.Join(req.Sources, ","))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxReport
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx reports: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxGetReport 获取报告详情
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/report/get_report.html
func (c *Client) TestxGetReport(ctx context.Context, req *model.TestxGetReportRequest) (*model.TestxReport, error) {
	path := fmt.Sprintf("/api/testx/report/v1/namespaces/%s/reports/%s", req.Namespace, req.Uid)
	resp, err := c.doTestxGet(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	var result []model.TestxReport
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx report: %w", err)
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

// TestxGetReportData 获取报告详情数据
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/report/get_report_data.html
func (c *Client) TestxGetReportData(ctx context.Context, req *model.TestxGetReportDataRequest) (*model.TestxReportData, error) {
	path := fmt.Sprintf("/api/testx/report/v1/namespaces/%s/reports/%s/templates/%s/data",
		req.Namespace, req.ReportUid, req.TemplateUid)
	resp, err := c.doTestxGet(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	return &model.TestxReportData{Raw: resp.Data}, nil
}

// TestxListReportTemplates 获取报告模板列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/report/list_template.html
func (c *Client) TestxListReportTemplates(ctx context.Context, req *model.TestxListReportTemplatesRequest) ([]model.TestxReportTemplate, int, error) {
	path := fmt.Sprintf("/api/testx/report/v1/namespaces/%s/templates", req.Namespace)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxReportTemplate
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx report templates: %w", err)
	}
	return result, resp.TotalCount, nil
}
