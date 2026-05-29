package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// StorageSave 保存数据到公共存储
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/save.html
func (c *Client) StorageSave(ctx context.Context, req *model.StorageSaveRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/open_app_storage/save", req.ToParams())
}

// StorageQuery 查询公共存储数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/query.html
func (c *Client) StorageQuery(ctx context.Context, req *model.StorageQueryRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/open_app_storage/query", req.ToParams())
}

// StorageUpdate 更新公共存储数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/update.html
func (c *Client) StorageUpdate(ctx context.Context, req *model.StorageUpdateRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/open_app_storage/update", req.ToParams())
}

// StorageDelete 删除公共存储数据
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/delete.html
func (c *Client) StorageDelete(ctx context.Context, req *model.StorageDeleteRequest) (json.RawMessage, error) {
	return c.doPost(ctx, "/open_app_storage/delete", req.ToParams())
}
