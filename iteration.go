package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListIterations 查询迭代列表，返回强类型 Iteration 切片
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations.html
func (c *Client) ListIterations(ctx context.Context, req *model.ListIterationsRequest) ([]model.Iteration, error) {
	data, err := c.doGet(ctx, "/iterations", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Iteration](data, "Iteration")
}

// CreateIteration 创建迭代，返回创建后的完整 Iteration 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/add_iteration.html
func (c *Client) CreateIteration(ctx context.Context, req *model.CreateIterationRequest) (*model.Iteration, error) {
	data, err := c.doPost(ctx, "/iterations", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Iteration](data, "Iteration")
}

// UpdateIteration 更新迭代，返回更新后的完整 Iteration 对象
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/update_iteration.html
func (c *Client) UpdateIteration(ctx context.Context, req *model.UpdateIterationRequest) (*model.Iteration, error) {
	data, err := c.doPost(ctx, "/iterations", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.Iteration](data, "Iteration")
}

// CountIterations 查询迭代数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations_count.html
func (c *Client) CountIterations(ctx context.Context, req *model.CountIterationsRequest) (int, error) {
	data, err := c.doGet(ctx, "/iterations/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// LockIteration 锁定迭代
func (c *Client) LockIteration(ctx context.Context, req *model.IterationLockRequest) error {
	_, err := c.doPost(ctx, "/iterations/lock", req.ToParams())
	return err
}

// UnlockIteration 解锁迭代
func (c *Client) UnlockIteration(ctx context.Context, req *model.IterationLockRequest) error {
	_, err := c.doPost(ctx, "/iterations/unlock", req.ToParams())
	return err
}
