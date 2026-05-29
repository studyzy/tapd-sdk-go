package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestGetPersonalSetting(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/get_personal_setting" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "48464494" {
			t.Errorf("workspace_id = %q", r.URL.Query().Get("workspace_id"))
		}
		if r.URL.Query().Get("nick") != "ocenhu" {
			t.Errorf("nick = %q", r.URL.Query().Get("nick"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":1,"data":{"language":"zh_CN","message_setting":[{"type":"LETTER_STORY","disable":["LETTER_STORY_FOLLOW_UPDATE"],"enable":["LETTER_STORY_FOLLOW_DELETE"]}]},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	got, err := c.GetPersonalSetting(context.Background(), &model.GetPersonalSettingRequest{
		WorkspaceID: "48464494",
		Nick:        "ocenhu",
	})
	if err != nil {
		t.Fatalf("GetPersonalSetting() unexpected error: %v", err)
	}
	if got.Language != "zh_CN" {
		t.Errorf("language = %q, want zh_CN", got.Language)
	}
	if len(got.MessageSetting) != 1 {
		t.Fatalf("expected 1 message_setting, got %d", len(got.MessageSetting))
	}
	if got.MessageSetting[0].Type != "LETTER_STORY" {
		t.Errorf("type = %q", got.MessageSetting[0].Type)
	}
	if len(got.MessageSetting[0].Disable) != 1 || got.MessageSetting[0].Disable[0] != "LETTER_STORY_FOLLOW_UPDATE" {
		t.Errorf("disable = %v", got.MessageSetting[0].Disable)
	}
	if len(got.MessageSetting[0].Enable) != 1 || got.MessageSetting[0].Enable[0] != "LETTER_STORY_FOLLOW_DELETE" {
		t.Errorf("enable = %v", got.MessageSetting[0].Enable)
	}
}

func TestGetThirdUserMapping(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/get_third_user_mapping" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "48464494" {
			t.Errorf("workspace_id = %q", r.URL.Query().Get("workspace_id"))
		}
		if r.URL.Query().Get("user_id") != "1223" {
			t.Errorf("user_id = %q", r.URL.Query().Get("user_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":1,"data":{"third_partys":[{"third_party_id":"123","third_party_type":"qywx"}]},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	got, err := c.GetThirdUserMapping(context.Background(), &model.GetThirdUserMappingRequest{
		WorkspaceID: "48464494",
		UserID:      "1223",
	})
	if err != nil {
		t.Fatalf("GetThirdUserMapping() unexpected error: %v", err)
	}
	if len(got.ThirdPartys) != 1 {
		t.Fatalf("expected 1 third_party, got %d", len(got.ThirdPartys))
	}
	if got.ThirdPartys[0].ThirdPartyID != "123" {
		t.Errorf("third_party_id = %q", got.ThirdPartys[0].ThirdPartyID)
	}
	if got.ThirdPartys[0].ThirdPartyType != "qywx" {
		t.Errorf("third_party_type = %q", got.ThirdPartys[0].ThirdPartyType)
	}
}
