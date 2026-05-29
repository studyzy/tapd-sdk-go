// webhook.go 提供 TAPD Webhook 回调载荷的解析助手。
//
// TAPD Webhook 由 TAPD 平台主动 POST 到接入方提供的 URL，因此本文件不发起
// 任何 HTTP 请求；它仅负责把请求体反序列化成 model.WebhookEvent。
package tapd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/studyzy/tapd-sdk-go/model"
)

// Webhook 数据格式常量，对应 TAPD 接入时可选的 "json" 与 "form"。
const (
	WebhookFormatJSON = "json"
	WebhookFormatForm = "form"
)

// ParseWebhookEvent 将 TAPD Webhook 请求体解析为 *model.WebhookEvent。
//
// format 取值 "json" 或 "form"（默认，application/x-www-form-urlencoded）。
// 所有以 old_ 开头的字段会被收集到返回值的 OldFields 中（键去除 "old_" 前缀）。
func ParseWebhookEvent(format string, body []byte) (*model.WebhookEvent, error) {
	switch strings.ToLower(strings.TrimSpace(format)) {
	case WebhookFormatJSON:
		return parseWebhookJSON(body)
	case WebhookFormatForm, "":
		return parseWebhookForm(body)
	default:
		return nil, fmt.Errorf("unsupported webhook format: %q", format)
	}
}

func parseWebhookJSON(body []byte) (*model.WebhookEvent, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("parse webhook json: %w", err)
	}
	evt := &model.WebhookEvent{OldFields: map[string]string{}}
	if err := json.Unmarshal(body, evt); err != nil {
		return nil, fmt.Errorf("parse webhook json: %w", err)
	}
	for k, v := range raw {
		if !strings.HasPrefix(k, "old_") {
			continue
		}
		var s string
		if err := json.Unmarshal(v, &s); err != nil {
			// 非字符串值：保留 JSON 原文作为兜底。
			s = string(v)
		}
		evt.OldFields[strings.TrimPrefix(k, "old_")] = s
	}
	return evt, nil
}

func parseWebhookForm(body []byte) (*model.WebhookEvent, error) {
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, fmt.Errorf("parse webhook form: %w", err)
	}
	evt := &model.WebhookEvent{
		Event:        values.Get("event"),
		EventFrom:    values.Get("event_from"),
		WorkspaceID:  values.Get("workspace_id"),
		CurrentUser:  values.Get("current_user"),
		EventID:      values.Get("event_id"),
		ID:           values.Get("id"),
		Secret:       values.Get("secret"),
		Created:      values.Get("created"),
		ChangeFields: values.Get("change_fields"),
		OldFields:    map[string]string{},
	}
	for k, v := range values {
		if !strings.HasPrefix(k, "old_") || len(v) == 0 {
			continue
		}
		evt.OldFields[strings.TrimPrefix(k, "old_")] = v[0]
	}
	return evt, nil
}
