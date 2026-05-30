package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetLifeTimes 获取状态流转时间
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/measure/get_life_times.html
func (c *Client) GetLifeTimes(ctx context.Context, req *model.GetLifeTimesRequest) ([]model.LifeTime, error) {
	data, err := c.doGet(ctx, "/life_times", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.LifeTime](data, "LifeTime")
}
