package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateBoardCard 新建看板工作项
func (c *Client) CreateBoardCard(ctx context.Context, req *model.CreateBoardCardRequest) (*model.BoardCard, error) {
	data, err := c.doPost(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create board card response: %w", err)
	}

	raw, ok := wrapper["BoardCard"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var card model.BoardCard
	if err := json.Unmarshal(raw, &card); err != nil {
		return nil, fmt.Errorf("failed to parse created board card: %w", err)
	}

	return &card, nil
}

// GetBoardCards 获取看板工作项列表
func (c *Client) GetBoardCards(ctx context.Context, req *model.GetBoardCardsRequest) ([]model.BoardCard, error) {
	data, err := c.doGet(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse board card list: %w", err)
	}

	results := make([]model.BoardCard, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["BoardCard"]; ok {
			var card model.BoardCard
			if err := json.Unmarshal(raw, &card); err == nil {
				results = append(results, card)
			}
		}
	}
	return results, nil
}

// UpdateBoardCard 更新看板工作项
func (c *Client) UpdateBoardCard(ctx context.Context, req *model.UpdateBoardCardRequest) (*model.BoardCard, error) {
	data, err := c.doPost(ctx, "/board_cards", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update board card response: %w", err)
	}

	raw, ok := wrapper["BoardCard"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var card model.BoardCard
	if err := json.Unmarshal(raw, &card); err != nil {
		return nil, fmt.Errorf("failed to parse updated board card: %w", err)
	}

	return &card, nil
}

// GetBoardColumns 获取看板板块列表
func (c *Client) GetBoardColumns(ctx context.Context, req *model.GetBoardColumnsRequest) ([]model.BoardColumn, error) {
	data, err := c.doGet(ctx, "/board_columns", req.ToParams())
	if err != nil {
		return nil, err
	}

	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse board column list: %w", err)
	}

	results := make([]model.BoardColumn, 0, len(rawList))
	for _, item := range rawList {
		if raw, ok := item["BoardColumn"]; ok {
			var col model.BoardColumn
			if err := json.Unmarshal(raw, &col); err == nil {
				results = append(results, col)
			}
		}
	}
	return results, nil
}
