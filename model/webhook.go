// Package model 中的 webhook.go 定义了 TAPD Webhook 事件回调的数据模型。
//
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/webhook/webhook_document.html
//
// TAPD Webhook 由 TAPD 平台主动 POST 到接入方提供的 URL，数据格式可选 json
// 或 form（application/x-www-form-urlencoded，默认）。本 SDK 仅提供反序列化
// 用的结构体与解析助手，不主动调用任何 TAPD 接口。
package model

// Webhook 事件名常量。文档列出了需求/缺陷/任务/发布评审的创建、更新、状态
// 变更、删除以及前后置对象绑定/解绑事件。
const (
	WebhookEventStoryCreate       = "story::create"
	WebhookEventStoryUpdate       = "story::update"
	WebhookEventStoryStatusChange = "story::status_change"
	WebhookEventStoryDelete       = "story::delete"
	WebhookEventBugCreate         = "bug::create"
	WebhookEventBugUpdate         = "bug::update"
	WebhookEventBugStatusChange   = "bug::status_change"
	WebhookEventBugDelete         = "bug::delete"
	WebhookEventTaskCreate        = "task::create"
	WebhookEventTaskUpdate        = "task::update"
	WebhookEventTaskStatusChange  = "task::status_change"
	WebhookEventTaskDelete        = "task::delete"
	WebhookEventLaunchformCreate  = "launchform::create"
	WebhookEventLaunchformUpdate  = "launchform::update"
	WebhookEventRelationBinding   = "relation::binding"
	WebhookEventRelationUnbinding = "relation::unbinding"
)

// WebhookEvent 表示一次 TAPD Webhook 回调的载荷。
//
// 同一结构体覆盖文档中描述的「新建 / 更新 / 删除」三类事件；状态变更
// （*::status_change）以及前后置对象绑定/解绑事件复用相同字段。所有 old_*
// 字段（仅出现在更新与状态变更事件中）通过 OldFields 暴露，键已去除
// "old_" 前缀。
type WebhookEvent struct {
	// Event 事件名，例如 story::create。
	Event string `json:"event,omitempty"`
	// EventFrom 事件触发位置，可能取值：web、api。
	EventFrom string `json:"event_from,omitempty"`
	// WorkspaceID 项目 ID。
	WorkspaceID string `json:"workspace_id,omitempty"`
	// CurrentUser 操作人昵称。
	CurrentUser string `json:"current_user,omitempty"`
	// EventID 事件 ID。
	EventID string `json:"event_id,omitempty"`
	// ID 事件涉及对象的 ID（需求/缺陷/任务/发布评审为 19 位长 ID）。
	ID string `json:"id,omitempty"`
	// Secret 接入时双方约定的验证密码。
	Secret string `json:"secret,omitempty"`
	// Created 事件触发时间，格式 Y-m-d H:i:s，如 2017-04-12 09:04:29。
	Created string `json:"created,omitempty"`

	// ChangeFields 仅出现在更新/状态变更事件，逗号分隔的变更字段名列表。
	ChangeFields string `json:"change_fields,omitempty"`

	// OldFields 保存所有 old_* 字段的变更前值，键已去除 "old_" 前缀。
	// 对于 form 编码请求，值为字符串原文；对于 json 请求，非字符串值会
	// 以其 JSON 文本形式保存。
	OldFields map[string]string `json:"-"`
}
