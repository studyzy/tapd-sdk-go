package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetLifeTimes 获取状态流转时间
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/measure/get_life_times.html
func (c *Client) GetLifeTimes(ctx context.Context, req *model.GetLifeTimesRequest) (json.RawMessage, error) {
	data, err := c.doGet(ctx, "/status_flows/life_times", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}
