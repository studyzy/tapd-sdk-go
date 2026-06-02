package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/studyzy/tapd-sdk-go/model"
)

// TestxSearchDesigns 查询测试设计列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/design/search_design.html
func (c *Client) TestxSearchDesigns(ctx context.Context, req *model.TestxSearchDesignsRequest) ([]model.TestxDesign, int, error) {
	path := fmt.Sprintf("/api/testx/design/v2/namespaces/%s/designs/search", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxDesign
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx designs: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxListDesignStats 查询测试设计列表统计信息
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/design/list_stat.html
func (c *Client) TestxListDesignStats(ctx context.Context, req *model.TestxListDesignStatsRequest) ([]model.TestxDesignStat, error) {
	path := fmt.Sprintf("/api/testx/design/v2/namespaces/%s/stat-list", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result []model.TestxDesignStat
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx design stats: %w", err)
	}
	return result, nil
}

// TestxListDesignLabels 查询测试设计标签列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/design/list_label.html
func (c *Client) TestxListDesignLabels(ctx context.Context, req *model.TestxListDesignLabelsRequest) ([]model.TestxDesignLabel, int, error) {
	path := fmt.Sprintf("/api/testx/design/v2/namespaces/%s/labels", req.Namespace)
	q := url.Values{}
	if req.DesignUid != "" {
		q.Set("design_uid", req.DesignUid)
	}
	if req.Kind != "" {
		q.Set("kind", req.Kind)
	}
	if req.Name != "" {
		q.Set("name", req.Name)
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxDesignLabel
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx design labels: %w", err)
	}
	return result, resp.TotalCount, nil
}
