package tapd

import (
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// AddCodeCommitInfo 保存 Commit 提交数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/add_code_commit_info.html
func (c *Client) AddCodeCommitInfo(req *model.AddCodeCommitInfoRequest) (json.RawMessage, error) {
	data, err := c.doPost("/code_commit_infos", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetCodeCommitInfos 获取 GIT 关联提交数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/get_code_commit_infos.html
func (c *Client) GetCodeCommitInfos(req *model.GetCodeCommitInfosRequest) (json.RawMessage, error) {
	data, err := c.doGet("/code_commit_infos", req.ToParams())
	if err != nil {
		return nil, err
	}
	return data, nil
}
