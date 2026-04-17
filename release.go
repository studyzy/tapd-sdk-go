package tapd

import (
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// CreateRelease 创建发布计划
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/add_release.html
func (c *Client) CreateRelease(req *model.CreateReleaseRequest) (*model.Release, error) {
	data, err := c.doPost("/releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse create release response: %w", err)
	}

	raw, ok := wrapper["Release"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var release model.Release
	if err := json.Unmarshal(raw, &release); err != nil {
		return nil, fmt.Errorf("failed to parse created release: %w", err)
	}
	return &release, nil
}

// UpdateRelease 更新发布计划
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/update_release.html
func (c *Client) UpdateRelease(req *model.UpdateReleaseRequest) (*model.Release, error) {
	data, err := c.doPost("/releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse update release response: %w", err)
	}

	raw, ok := wrapper["Release"]
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var release model.Release
	if err := json.Unmarshal(raw, &release); err != nil {
		return nil, fmt.Errorf("failed to parse updated release: %w", err)
	}
	return &release, nil
}

// CountReleases 获取发布计划数量
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_releases_count.html
func (c *Client) CountReleases(req *model.WorkspaceIDRequest) (int, error) {
	data, err := c.doGet("/releases/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	var result struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse release count: %w", err)
	}
	return result.Count, nil
}
