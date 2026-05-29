// Package model 中的 attachment.go 定义了 TAPD 附件和图片相关数据模型
package model

import (
	"encoding/json"
)

// Attachment 表示 TAPD 附件
type Attachment struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	EntryID     string `json:"entry_id,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Description string `json:"description,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Created     string `json:"created,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Owner       string `json:"owner,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
}

// ImageInfo 表示 TAPD 图片下载信息
type ImageInfo struct {
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Filename    string `json:"filename,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
}

// UnmarshalJSON 自定义解码器，兼容 TAPD /files/get_image 接口的特殊响应：
// 其它接口的 workspace_id 都是带引号的 JSON 字符串，但这个接口返的是 JSON 数字，
// 默认解码会报 "cannot unmarshal number into Go struct field ... of type string"。
// 此处单独把 workspace_id 取成 json.RawMessage，再按"字符串或数字"宽松转换为 string，
// 其余字段保持标准解码逻辑（通过同结构 alias 避免递归调用本方法）。
func (i *ImageInfo) UnmarshalJSON(data []byte) error {
	type shadow ImageInfo
	aux := struct {
		WorkspaceID json.RawMessage `json:"workspace_id,omitempty"`
		*shadow
	}{shadow: (*shadow)(i)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.WorkspaceID) == 0 || string(aux.WorkspaceID) == "null" {
		return nil
	}
	// 既可能是 "61692131" 也可能是 61692131
	if aux.WorkspaceID[0] == '"' {
		return json.Unmarshal(aux.WorkspaceID, &i.WorkspaceID)
	}
	// 数字形态：用 json.Number 拿到原始字面量，避免精度丢失
	var n json.Number
	if err := json.Unmarshal(aux.WorkspaceID, &n); err != nil {
		return err
	}
	i.WorkspaceID = n.String()
	return nil
}

// GetImageRequest 获取图片的请求参数
type GetImageRequest struct {
	WorkspaceID string // 必填：项目 ID
	ImagePath   string // 必填：图片路径
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetImageRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "image_path", r.ImagePath)
	return params
}

// GetAttachmentsRequest 获取附件列表的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/attachment/get_attachments.html
type GetAttachmentsRequest struct {
	WorkspaceID string // 必填：项目 ID
	ID          string // 可选：附件 ID
	Type        string // 可选：对象类型
	EntryID     string // 可选：条目 ID
	Filename    string // 可选：附件名称
	Owner       string // 可选：上传人
	Limit       int    // 可选：返回数量限制
	Page        int    // 可选：页码
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetAttachmentsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
	}
	setOptional(params, "id", r.ID)
	setOptional(params, "type", r.Type)
	setOptional(params, "entry_id", r.EntryID)
	setOptional(params, "filename", r.Filename)
	setOptional(params, "owner", r.Owner)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	return params
}
