// Package model 中的 storage.go 定义了公共存储相关 API 的请求/响应类型
package model

// StorageSaveRequest 保存数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/save.html
type StorageSaveRequest struct {
	Document string // 必填：文档名，若不存在将被创建
	Data     string // 必填：要保存的数据，JSON 格式
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *StorageSaveRequest) ToParams() map[string]string {
	return map[string]string{
		"document": r.Document,
		"data":     r.Data,
	}
}

// StorageQueryRequest 查询数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/query.html
type StorageQueryRequest struct {
	Document  string // 必填：文档名
	Condition string // 可选：查询条件，JSON 格式
	Order     string // 可选：排序规则
	Limit     int    // 可选：条数限制
	Offset    int    // 可选：偏移量
	Fields    string // 可选：返回字段
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *StorageQueryRequest) ToParams() map[string]string {
	params := map[string]string{
		"document": r.Document,
	}
	setOptional(params, "condition", r.Condition)
	setOptional(params, "order", r.Order)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "offset", r.Offset)
	setOptional(params, "fields", r.Fields)
	return params
}

// StorageUpdateRequest 更新数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/update.html
type StorageUpdateRequest struct {
	Document  string // 必填：文档名
	Condition string // 必填：查询条件，JSON 格式
	Data      string // 必填：要更新的数据，JSON 格式
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *StorageUpdateRequest) ToParams() map[string]string {
	return map[string]string{
		"document":  r.Document,
		"condition": r.Condition,
		"data":      r.Data,
	}
}

// StorageDeleteRequest 删除数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/delete.html
type StorageDeleteRequest struct {
	Document  string // 必填：文档名
	Condition string // 必填：查询条件，JSON 格式
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *StorageDeleteRequest) ToParams() map[string]string {
	return map[string]string{
		"document":  r.Document,
		"condition": r.Condition,
	}
}
