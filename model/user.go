// Package model 中的 user.go 定义了用户相关 API 的请求/响应类型
package model

// GetPersonalSettingRequest 获取用户个人配置的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/user/get_personal_setting.html
type GetPersonalSettingRequest struct {
	WorkspaceID string // 必填：公司 ID
	Nick        string // 必填：用户唯一标识
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetPersonalSettingRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"nick":         r.Nick,
	}
}

// MessageSetting 用户消息通知配置项
type MessageSetting struct {
	Type    string   `json:"type,omitempty"`    // 类别，例如 LETTER_STORY、MAIL_BUG 等
	Enable  []string `json:"enable,omitempty"`  // 启用的功能项
	Disable []string `json:"disable,omitempty"` // 禁用的功能项
}

// PersonalSetting 用户个人配置
type PersonalSetting struct {
	Language       string           `json:"language,omitempty"`        // 当前用户设置的系统语种
	MessageSetting []MessageSetting `json:"message_setting,omitempty"` // 消息通知配置
}

// GetThirdUserMappingRequest 获取用户在三方系统映射的 userId 请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/user/get_third_user_mapping.html
type GetThirdUserMappingRequest struct {
	WorkspaceID string // 必填：公司 ID
	UserID      string // 必填：用户 ID
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetThirdUserMappingRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"user_id":      r.UserID,
	}
}

// ThirdPartyMapping 用户在第三方系统的映射条目
type ThirdPartyMapping struct {
	ThirdPartyID   string `json:"third_party_id,omitempty"`   // 第三方系统人员 ID
	ThirdPartyType string `json:"third_party_type,omitempty"` // 系统类别
}

// ThirdUserMapping 用户关联的第三方系统映射信息
type ThirdUserMapping struct {
	ThirdPartys []ThirdPartyMapping `json:"third_partys,omitempty"`
}
