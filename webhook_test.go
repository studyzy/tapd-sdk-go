package tapd

import (
	"testing"
)

func TestParseWebhookEvent_JSON(t *testing.T) {
	body := []byte(`{"event":"story::create","event_from":"web","workspace_id":"123","current_user":"admin","event_id":"evt001","id":"100001","secret":"mysecret","created":"2024-01-01 10:00:00","change_fields":"status,owner","old_status":"planning","old_owner":"alice"}`)

	evt, err := ParseWebhookEvent(WebhookFormatJSON, body)
	if err != nil {
		t.Fatalf("ParseWebhookEvent JSON: unexpected error: %v", err)
	}
	if evt.Event != "story::create" {
		t.Errorf("event = %q, want %q", evt.Event, "story::create")
	}
	if evt.WorkspaceID != "123" {
		t.Errorf("workspace_id = %q, want %q", evt.WorkspaceID, "123")
	}
	if evt.CurrentUser != "admin" {
		t.Errorf("current_user = %q, want %q", evt.CurrentUser, "admin")
	}
	if evt.ID != "100001" {
		t.Errorf("id = %q, want %q", evt.ID, "100001")
	}
	if evt.Secret != "mysecret" {
		t.Errorf("secret = %q, want %q", evt.Secret, "mysecret")
	}
	if evt.ChangeFields != "status,owner" {
		t.Errorf("change_fields = %q, want %q", evt.ChangeFields, "status,owner")
	}
	// old_* 字段应去掉前缀后收入 OldFields
	if got := evt.OldFields["status"]; got != "planning" {
		t.Errorf("OldFields[status] = %q, want %q", got, "planning")
	}
	if got := evt.OldFields["owner"]; got != "alice" {
		t.Errorf("OldFields[owner] = %q, want %q", got, "alice")
	}
}

func TestParseWebhookEvent_Form(t *testing.T) {
	body := []byte("event=bug%3A%3Aupdate&event_from=api&workspace_id=456&current_user=bob&event_id=evt002&id=200002&secret=s3cr3t&created=2024-06-01+09%3A30%3A00&change_fields=severity&old_severity=high")

	evt, err := ParseWebhookEvent(WebhookFormatForm, body)
	if err != nil {
		t.Fatalf("ParseWebhookEvent form: unexpected error: %v", err)
	}
	if evt.Event != "bug::update" {
		t.Errorf("event = %q, want %q", evt.Event, "bug::update")
	}
	if evt.WorkspaceID != "456" {
		t.Errorf("workspace_id = %q, want %q", evt.WorkspaceID, "456")
	}
	if evt.CurrentUser != "bob" {
		t.Errorf("current_user = %q, want %q", evt.CurrentUser, "bob")
	}
	if evt.EventFrom != "api" {
		t.Errorf("event_from = %q, want %q", evt.EventFrom, "api")
	}
	if evt.ID != "200002" {
		t.Errorf("id = %q, want %q", evt.ID, "200002")
	}
	if evt.ChangeFields != "severity" {
		t.Errorf("change_fields = %q, want %q", evt.ChangeFields, "severity")
	}
	if got := evt.OldFields["severity"]; got != "high" {
		t.Errorf("OldFields[severity] = %q, want %q", got, "high")
	}
}

func TestParseWebhookEvent_EmptyBody(t *testing.T) {
	// JSON 格式：空 body 应返回 error
	_, err := ParseWebhookEvent(WebhookFormatJSON, []byte{})
	if err == nil {
		t.Error("ParseWebhookEvent JSON empty body: expected error, got nil")
	}

	// Form 格式：空 body 应返回空事件，不报错（url.ParseQuery("") 返回空 map，合法）
	evt, err := ParseWebhookEvent(WebhookFormatForm, []byte{})
	if err != nil {
		t.Errorf("ParseWebhookEvent form empty body: unexpected error: %v", err)
	}
	if evt == nil {
		t.Fatal("ParseWebhookEvent form empty body: expected non-nil event")
	}
	if evt.Event != "" {
		t.Errorf("event = %q, want empty string", evt.Event)
	}
}

func TestParseWebhookEvent_MalformedJSON(t *testing.T) {
	body := []byte(`{"event":"story::create","workspace_id":`) // 截断的 JSON

	_, err := ParseWebhookEvent(WebhookFormatJSON, body)
	if err == nil {
		t.Error("ParseWebhookEvent malformed JSON: expected error, got nil")
	}
}

func TestParseWebhookEvent_UnsupportedFormat(t *testing.T) {
	_, err := ParseWebhookEvent("xml", []byte("<event/>"))
	if err == nil {
		t.Error("ParseWebhookEvent unsupported format: expected error, got nil")
	}
}
