package tapd

import (
	"context"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateBoardCard 新建看板工作项
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/board/add_board_card.html
func (c *Client) CreateBoardCard(ctx context.Context, req *model.CreateBoardCardRequest) (*model.BoardCard, error) {
	data, err := c.doPost(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	list, err := parseList[model.BoardCard](data, "BoardCard")
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("empty response from create board card")
	}
	return &list[0], nil
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
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/board/update_board_card.html
func (c *Client) UpdateBoardCard(ctx context.Context, req *model.UpdateBoardCardRequest) (*model.BoardCard, error) {
	data, err := c.doPost(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	list, err := parseList[model.BoardCard](data, "BoardCard")
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("empty response from update board card")
	}
	return &list[0], nil
}

// GetBoardColumns 获取看板板块列表
func (c *Client) GetBoardColumns(ctx context.Context, req *model.GetBoardColumnsRequest) ([]model.BoardColumn, error) {
	data, err := c.doGet(ctx, "/board_columns", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.BoardColumn](data, "Column")
}
