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

func TestGetUserViewList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user_oauth/get_user_view_list" {
			t.Errorf("unexpected path: %s, want /user_oauth/get_user_view_list", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s, want GET", r.Method)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want 10104801", r.URL.Query().Get("workspace_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":1,"data":{"1000000000000000016":{"id":"1010104801016212067","title":"所有的","enable":"1","type":"system","default_show":"true","view_id":"1000000000000000016","sort":"1"},"1000000000000000017":{"id":"1010104801016212069","title":"我负责的","enable":"1","type":"system","default_show":"false","view_id":"1000000000000000017","sort":"2"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	views, err := c.GetUserViewList(context.Background(), &model.GetUserViewListRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("GetUserViewList() unexpected error: %v", err)
	}
	if len(views) != 2 {
		t.Fatalf("expected 2 views, got %d", len(views))
	}
	v, ok := views["1000000000000000016"]
	if !ok {
		t.Fatal("expected key '1000000000000000016' in views")
	}
	if v.ID != "1010104801016212067" {
		t.Errorf("id = %q, want %q", v.ID, "1010104801016212067")
	}
	if v.Title != "所有的" {
		t.Errorf("title = %q, want %q", v.Title, "所有的")
	}
	if v.Type != "system" {
		t.Errorf("type = %q, want %q", v.Type, "system")
	}
	if v.DefaultShow != "true" {
		t.Errorf("default_show = %q, want %q", v.DefaultShow, "true")
	}
}

func TestGetUsersInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/info" {
			t.Errorf("unexpected path: %s, want /users/info", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s, want GET", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":1,"data":{"id":"6081","nick":"robertyang","name":"杨晓俊","avatar":"http://tiger.oa.com/0/users/avatar/6081/jpg/0/large","enabled":"1","status_id":"1","status_name":"在职"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	info, err := c.GetUsersInfo(context.Background())
	if err != nil {
		t.Fatalf("GetUsersInfo() unexpected error: %v", err)
	}
	if info.ID != "6081" {
		t.Errorf("id = %q, want %q", info.ID, "6081")
	}
	if info.Nick != "robertyang" {
		t.Errorf("nick = %q, want %q", info.Nick, "robertyang")
	}
	if info.Name != "杨晓俊" {
		t.Errorf("name = %q, want %q", info.Name, "杨晓俊")
	}
	if info.Enabled != "1" {
		t.Errorf("enabled = %q, want %q", info.Enabled, "1")
	}
	if info.StatusName != "在职" {
		t.Errorf("status_name = %q, want %q", info.StatusName, "在职")
	}
}
