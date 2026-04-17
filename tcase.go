package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListTCases 查询测试用例列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases.html
func (c *Client) ListTCases(ctx context.Context, req *model.ListTCasesRequest) ([]model.TCase, error) {
	data, err := c.doGet(ctx, "/tcases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.TCase](data, "Tcase")
}

// CountTCases 查询测试用例数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases_count.html
func (c *Client) CountTCases(ctx context.Context, req *model.CountTCasesRequest) (int, error) {
	data, err := c.doGet(ctx, "/tcases/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// CreateTCase 创建测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase.html
func (c *Client) CreateTCase(ctx context.Context, req *model.CreateTCaseRequest) (*model.TCase, error) {
	data, err := c.doPost(ctx, "/tcases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.TCase](data, "Tcase")
}

// BatchCreateTCases 批量创建测试用例
func (c *Client) BatchCreateTCases(ctx context.Context, req *model.BatchCreateTCasesRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/tcases/batch_save", req.ToParams())
}
