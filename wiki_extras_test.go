package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestCountWikiAttachments(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_attachments/count" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":7},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountWikiAttachments(context.Background(), &model.CountWikiAttachmentsRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("CountWikiAttachments() unexpected error: %v", err)
	}
	if count != 7 {
		t.Errorf("count = %d, want 7", count)
	}
}

func TestGetWikiDrawio(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_drawios" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("id") != "1100000000000001102" {
			t.Errorf("missing id param: %q", r.URL.Query().Get("id"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"StaticData":{"id":"1100000000000001102","values":"<mxGraphModel/>"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	d, err := c.GetWikiDrawio(context.Background(), &model.GetWikiDrawioRequest{
		WorkspaceID: "10104801",
		ID:          "1100000000000001102",
	})
	if err != nil {
		t.Fatalf("GetWikiDrawio() unexpected error: %v", err)
	}
	if d.ID != "1100000000000001102" {
		t.Errorf("ID = %q", d.ID)
	}
	if d.Values != "<mxGraphModel/>" {
		t.Errorf("Values = %q", d.Values)
	}
}

func TestListWikiEntityPermissions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_entity_permissions" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("wiki_id") != "1210104801001897607" {
			t.Errorf("missing wiki_id: %q", r.URL.Query().Get("wiki_id"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"EntityPermission":{"id":"1","workspace_id":"10158241","entry_type":"wiki","target_type":"role_id","target_id":"1000000000000000002","wiki_id":"1210158241000048769"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	perms, err := c.ListWikiEntityPermissions(context.Background(), &model.ListWikiEntityPermissionsRequest{
		WorkspaceID: "10104801",
		WikiID:      "1210104801001897607",
	})
	if err != nil {
		t.Fatalf("ListWikiEntityPermissions() unexpected error: %v", err)
	}
	if len(perms) != 1 {
		t.Fatalf("expected 1 permission, got %d", len(perms))
	}
	if perms[0].TargetType != "role_id" {
		t.Errorf("TargetType = %q", perms[0].TargetType)
	}
}

func TestListWikiFollowers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_followers" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"UserFollows":{"id":"1220358001000000595","workspace_id":"20358001","created":"2020-10-26 14:47:15","user":"aidenxiao","wiki_id":"1220358001000044887"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	followers, err := c.ListWikiFollowers(context.Background(), &model.ListWikiFollowersRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("ListWikiFollowers() unexpected error: %v", err)
	}
	if len(followers) != 1 {
		t.Fatalf("expected 1 follower, got %d", len(followers))
	}
	if followers[0].User != "aidenxiao" {
		t.Errorf("User = %q", followers[0].User)
	}
}

func TestCountWikiFollowers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_followers/count" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":23},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountWikiFollowers(context.Background(), &model.CountWikiFollowersRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("CountWikiFollowers() unexpected error: %v", err)
	}
	if count != 23 {
		t.Errorf("count = %d, want 23", count)
	}
}

func TestListWikiTags(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_tags" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Tags":{"creator":"huanjinxie","created":"2021-01-07 20:40:05","wiki_id":"1220358527000044697","tag":"首页"}},{"Tags":{"creator":"huanjinxie","created":"2021-01-07 20:40:05","wiki_id":"1220358527000044697","tag":"测试"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	tags, err := c.ListWikiTags(context.Background(), &model.ListWikiTagsRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("ListWikiTags() unexpected error: %v", err)
	}
	if len(tags) != 2 {
		t.Fatalf("expected 2 tags, got %d", len(tags))
	}
	if tags[0].Tag != "首页" {
		t.Errorf("Tag[0] = %q", tags[0].Tag)
	}
}

func TestCountWikiTags(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tapd_wikis_tags/count" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":2},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	count, err := c.CountWikiTags(context.Background(), &model.CountWikiTagsRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("CountWikiTags() unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}

func TestCountWikis_AllParams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("creator") != "alice" {
			t.Errorf("missing creator param")
		}
		if r.URL.Query().Get("name") != "技术" {
			t.Errorf("missing name param")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":3},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.CountWikis(context.Background(), &model.CountWikisRequest{
		WorkspaceID: "10104801",
		Creator:     "alice",
		Name:        "技术",
	})
	if err != nil {
		t.Fatalf("CountWikis() unexpected error: %v", err)
	}
}
