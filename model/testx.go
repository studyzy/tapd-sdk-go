package model

import "encoding/json"

// TestxResponse testx API 统一响应格式
type TestxResponse struct {
	RequestID  string          `json:"RequestId"`
	Error      json.RawMessage `json:"Error"`
	Data       json.RawMessage `json:"Data"`
	TotalCount int             `json:"TotalCount"`
}

// TestxPageInfo testx 分页参数
type TestxPageInfo struct {
	Offset uint32 `json:"Offset"`
	Limit  uint32 `json:"Limit"`
}

// TestxAudit testx 审计信息
type TestxAudit struct {
	Creator    string `json:"Creator,omitempty"`
	Updater    string `json:"Updater,omitempty"`
	CreatedAt  string `json:"CreatedAt,omitempty"`
	UpdatedAt  string `json:"UpdatedAt,omitempty"`
	Tenant     string `json:"Tenant,omitempty"`
	StartedAt  string `json:"StartedAt,omitempty"`
	EndedAt    string `json:"EndedAt,omitempty"`
	Starter    string `json:"Starter,omitempty"`
	Terminator string `json:"Terminator,omitempty"`
}

// TestxProperty testx 自定义属性
type TestxProperty struct {
	Name  string      `json:"Name"`
	Label string      `json:"Label"`
	Value interface{} `json:"Value"`
	Url   string      `json:"Url,omitempty"`
	Flag  string      `json:"flag,omitempty"`
}

// TestxIssue testx 关联问题
type TestxIssue struct {
	IssueUid     string      `json:"IssueUid,omitempty"`
	Namespace    string      `json:"Namespace,omitempty"`
	WorkspaceUid string      `json:"WorkspaceUid,omitempty"`
	IssueUrl     string      `json:"IssueUrl,omitempty"`
	Type         string      `json:"Type,omitempty"`
	Source       string      `json:"Source,omitempty"`
	Detail       interface{} `json:"Detail,omitempty"`
	IssueName    string      `json:"IssueName,omitempty"`
	IsDeleted    bool        `json:"IsDeleted,omitempty"`
	Uid          string      `json:"Uid,omitempty"`
}

// TestxLabel testx 标签
type TestxLabel struct {
	Name        string `json:"Name,omitempty"`
	Value       string `json:"Value,omitempty"`
	Tag         string `json:"Tag,omitempty"`
	Color       string `json:"Color,omitempty"`
	Uneditable  bool   `json:"Uneditable,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
	Module      string `json:"Module,omitempty"`
	Source      string `json:"Source,omitempty"`
}
