package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateBoardCard 新建看板工作项
func (c *Client) CreateBoardCard(ctx context.Context, req *model.CreateBoardCardRequest) (*model.BoardCard, error) {
	data, err := c.doPost(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.BoardCard](data, "BoardCard")
}

// GetBoardCards 获取看板工作项列表
func (c *Client) GetBoardCards(ctx context.Context, req *model.GetBoardCardsRequest) ([]model.BoardCard, error) {
	data, err := c.doGet(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.BoardCard](data, "BoardCard")
}

// UpdateBoardCard 更新看板工作项
func (c *Client) UpdateBoardCard(ctx context.Context, req *model.UpdateBoardCardRequest) (*model.BoardCard, error) {
	data, err := c.doPost(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.BoardCard](data, "BoardCard")
}

// GetBoardColumns 获取看板板块列表
func (c *Client) GetBoardColumns(ctx context.Context, req *model.GetBoardColumnsRequest) ([]model.BoardColumn, error) {
	data, err := c.doGet(ctx, "/board_columns", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.BoardColumn](data, "BoardColumn")
}
