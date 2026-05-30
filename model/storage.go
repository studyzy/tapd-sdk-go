// Package model 中的 storage.go 定义了公共存储相关 API 的请求/响应类型
package model

import "encoding/json"

// StorageSaveRequest 保存数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/save.html
type StorageSaveRequest struct {
	Document string      // 必填：文档名，若不存在将被创建
	Data     interface{} // 必填：要保存的数据，支持单条对象或批量数组
}

// ToJSON 将请求结构体序列化为 JSON 请求体
func (r *StorageSaveRequest) ToJSON() ([]byte, error) {
	body := map[string]interface{}{
		"document": r.Document,
		"data":     r.Data,
	}
	return json.Marshal(body)
}

// StorageQueryRequest 查询数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/query.html
type StorageQueryRequest struct {
	Document  string      // 必填：文档名
	Condition interface{} // 可选：查询条件，参考条件语法文档
	Order     string      // 可选：排序规则
	Limit     int         // 可选：条数限制
	Offset    int         // 可选：偏移量
	Fields    string      // 可选：返回字段
}

// ToJSON 将请求结构体序列化为 JSON 请求体
func (r *StorageQueryRequest) ToJSON() ([]byte, error) {
	body := map[string]interface{}{
		"document": r.Document,
	}
	if r.Condition != nil {
		body["condition"] = r.Condition
	}
	if r.Order != "" {
		body["order"] = r.Order
	}
	if r.Limit > 0 {
		body["limit"] = r.Limit
	}
	if r.Offset > 0 {
		body["offset"] = r.Offset
	}
	if r.Fields != "" {
		body["fields"] = r.Fields
	}
	return json.Marshal(body)
}

// StorageUpdateRequest 更新数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/update.html
type StorageUpdateRequest struct {
	Document  string      // 必填：文档名
	Condition interface{} // 必填：查询条件，参考条件语法文档
	Data      interface{} // 必填：要更新的数据，键值对
}

// ToJSON 将请求结构体序列化为 JSON 请求体
func (r *StorageUpdateRequest) ToJSON() ([]byte, error) {
	body := map[string]interface{}{
		"document":  r.Document,
		"condition": r.Condition,
		"data":      r.Data,
	}
	return json.Marshal(body)
}

// StorageDeleteRequest 删除数据的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/storage/delete.html
type StorageDeleteRequest struct {
	Document  string      // 必填：文档名
	Condition interface{} // 必填：查询条件，参考条件语法文档
}

// ToJSON 将请求结构体序列化为 JSON 请求体
func (r *StorageDeleteRequest) ToJSON() ([]byte, error) {
	body := map[string]interface{}{
		"document":  r.Document,
		"condition": r.Condition,
	}
	return json.Marshal(body)
}
