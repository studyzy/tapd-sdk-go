// Package model 定义了 tapd-sdk-go 的所有数据模型结构体
package model

import (
	"encoding/json"
	"fmt"
)

// Workspace 表示 TAPD 项目/工作区
type Workspace struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	PrettyName        string `json:"pretty_name,omitempty"`
	Category          string `json:"category,omitempty"`
	Status            string `json:"status,omitempty"`
	Description       string `json:"description,omitempty"`
	BeginDate         string `json:"begin_date,omitempty"`
	EndDate           string `json:"end_date,omitempty"`
	Closed            string `json:"closed,omitempty"`
	ExternalOn        string `json:"external_on,omitempty"`
	Creator           string `json:"creator,omitempty"`
	Created           string `json:"created,omitempty"`
	ProductType       string `json:"product_type,omitempty"`
	PlatformType      string `json:"platform_type,omitempty"`
	IsSelfDevelopment string `json:"is_self_development,omitempty"`
	Objective         string `json:"objective,omitempty"`
	Secrecy           string `json:"secrecy,omitempty"`
	Schedule          string `json:"schedule,omitempty"`
	Milestone         string `json:"milestone,omitempty"`
	Risk              string `json:"risk,omitempty"`
	CompanyID         string `json:"company_id,omitempty"`
	ParentID          string `json:"parent_id,omitempty"`
	MemberCount       string `json:"member_count,omitempty"`
	CreatorID         string `json:"creator_id,omitempty"`
	TemplateID        string `json:"template_id,omitempty"`
}

// UnmarshalJSON 自定义反序列化，兼容 TAPD API 中 member_count 和 company_id
// 可能返回 JSON number 或 string 两种格式
func (w *Workspace) UnmarshalJSON(data []byte) error {
	type Alias Workspace
	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		// 标准解析失败时，尝试从 raw map 中提取可能是 number 的字段
		var raw map[string]json.RawMessage
		if err2 := json.Unmarshal(data, &raw); err2 != nil {
			return err
		}
		// 移除可能导致类型不匹配的字段后重试
		delete(raw, "member_count")
		delete(raw, "company_id")
		cleaned, _ := json.Marshal(raw)
		if err2 := json.Unmarshal(cleaned, &alias); err2 != nil {
			return err
		}
		// 手动提取 number 类型字段
		if v, ok := raw["member_count"]; ok {
			alias.MemberCount = extractStringFromRaw(v)
		}
		if v, ok := raw["company_id"]; ok {
			alias.CompanyID = extractStringFromRaw(v)
		}
		*w = Workspace(alias)
		return nil
	}
	*w = Workspace(alias)

	// 标准解析成功但 member_count/company_id 可能为空（因为 number 无法解到 string）
	// 如果为空，尝试从 raw 中补充
	if w.MemberCount == "" || w.CompanyID == "" {
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(data, &raw); err == nil {
			if w.MemberCount == "" {
				if v, ok := raw["member_count"]; ok {
					w.MemberCount = extractStringFromRaw(v)
				}
			}
			if w.CompanyID == "" {
				if v, ok := raw["company_id"]; ok {
					w.CompanyID = extractStringFromRaw(v)
				}
			}
		}
	}
	return nil
}

// extractStringFromRaw 从 json.RawMessage 中提取值并转为 string
// 支持 JSON string ("123") 和 JSON number (123) 两种格式
func extractStringFromRaw(raw json.RawMessage) string {
	var s string
	if json.Unmarshal(raw, &s) == nil {
		return s
	}
	var n json.Number
	if json.Unmarshal(raw, &n) == nil {
		return n.String()
	}
	// fallback: 直接用 fmt
	var v interface{}
	if json.Unmarshal(raw, &v) == nil {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// UserWorkspace 表示项目成员信息
type UserWorkspace struct {
	UserID          string   `json:"user_id,omitempty"`
	User            string   `json:"user,omitempty"`
	RoleID          []string `json:"role_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Email           string   `json:"email,omitempty"`
	JoinProjectTime string   `json:"join_project_time,omitempty"`
	RealJoinTime    string   `json:"real_join_time,omitempty"`
	Status          string   `json:"status,omitempty"`
	Allocation      string   `json:"allocation,omitempty"`
}

// ListResponse 表示列表查询的通用响应结构
type ListResponse struct {
	Items   interface{} `json:"items"`
	Total   int         `json:"total,omitempty"`
	Page    int         `json:"page,omitempty"`
	Limit   int         `json:"limit,omitempty"`
	HasMore bool        `json:"has_more,omitempty"`
}

// SuccessResponse 表示创建/更新操作的成功响应
type SuccessResponse struct {
	Success     bool   `json:"success"`
	ID          string `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

// CountResponse 表示计数查询的响应
type CountResponse struct {
	Count int `json:"count"`
}

// TAPDResponse 表示 TAPD API 的统一响应包装格式
type TAPDResponse struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data"`
	Info   string          `json:"info"`
}
