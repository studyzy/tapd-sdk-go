package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// AddCodeCommitInfo 保存 Commit 提交数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/add_code_commit_info.html
func (c *Client) AddCodeCommitInfo(ctx context.Context, req *model.AddCodeCommitInfoRequest) (json.RawMessage, error) {
	body, err := req.ToJSON()
	if err != nil {
		return nil, err
	}
	return c.doPostJSONBody(ctx, "/code_commit_infos", body)
}

// GetCodeCommitInfos 获取 GIT 关联提交数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/get_code_commit_infos.html
func (c *Client) GetCodeCommitInfos(ctx context.Context, req *model.GetCodeCommitInfosRequest) (json.RawMessage, error) {
	data, err := c.doGet(ctx, "/code_commit_infos", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetCodeCommitObjects 获取指定 commit 关联的 TAPD 业务对象（需求、任务、缺陷）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/get_commit_objects.html
func (c *Client) GetCodeCommitObjects(ctx context.Context, req *model.GetCodeCommitObjectsRequest) (json.RawMessage, error) {
	data, err := c.doGet(ctx, "/code_commit_objects/workitems", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}
