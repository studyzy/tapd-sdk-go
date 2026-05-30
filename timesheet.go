package tapd

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListTimesheets 查询花费工时列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/get_timesheets.html
func (c *Client) ListTimesheets(ctx context.Context, req *model.ListTimesheetsRequest) ([]model.Timesheet, error) {
	data, err := c.doGet(ctx, "/timesheets", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Timesheet](data, "Timesheet")
}

// AddTimesheet 填写花费工时
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/add_timesheet.html
func (c *Client) AddTimesheet(ctx context.Context, req *model.AddTimesheetRequest) (*model.Timesheet, error) {
	data, err := c.doPost(ctx, "/timesheets", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Timesheet](data, "Timesheet")
}

// UpdateTimesheet 更新花费工时
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/update_timesheet.html
func (c *Client) UpdateTimesheet(ctx context.Context, req *model.UpdateTimesheetRequest) (*model.Timesheet, error) {
	data, err := c.doPost(ctx, "/timesheets", req.ToParams())
	if err != nil {
		return nil, err
	}

	// 尝试直接解析
	var ts model.Timesheet
	if err := json.Unmarshal(data, &ts); err == nil && ts.ID != "" {
		return &ts, nil
	}

	// 兼容 "Timesheet" 包裹
	return parseOne[model.Timesheet](data, "Timesheet")
}

// CountTimesheets 获取工时花费数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/timesheet/get_timesheets_count.html
func (c *Client) CountTimesheets(ctx context.Context, req *model.CountTimesheetsRequest) (int, error) {
	data, err := c.doGet(ctx, "/timesheets/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// DeleteTimesheets 删除工时花费
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/timesheet/delete_timesheets.html
func (c *Client) DeleteTimesheets(ctx context.Context, req *model.DeleteTimesheetsRequest) (json.RawMessage, error) {
	form := url.Values{}
	form.Set("workspace_id", req.WorkspaceID)
	form.Set("entity_type", req.EntityType)
	form.Set("entity_id", req.EntityID)
	for _, id := range req.CostIDs {
		form.Add("cost_ids[]", id)
	}
	return c.doPostForm(ctx, "/timesheets/delete_timesheets", form)
}
