package tapd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestLinkBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/link_bugs" {
			t.Errorf("unexpected path: %s, want /bugs/link_bugs", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":true,"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.LinkBugs(context.Background(), &model.LinkBugsRequest{
		WorkspaceID: "10104801",
		BugID:       "1001",
		RelateBugs:  "1002,1003",
	})
	if err != nil {
		t.Fatalf("LinkBugs() unexpected error: %v", err)
	}
}

func TestDeleteLinkBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/delete_link_bugs" {
			t.Errorf("unexpected path: %s, want /bugs/delete_link_bugs", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":true,"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.DeleteLinkBugs(context.Background(), &model.DeleteLinkBugsRequest{
		WorkspaceID: "10104801",
		BugID:       "1001",
		LinkIDs:     "2001,2002",
	})
	if err != nil {
		t.Fatalf("DeleteLinkBugs() unexpected error: %v", err)
	}
}

func TestGetLinkBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/get_link_bugs" {
			t.Errorf("unexpected path: %s, want /bugs/get_link_bugs", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"type":"repeat","id":"1010104801085894269","workspace_id":"10104801","actas":"target","linked_workspace_id":10104801,"link_id":"1162187798001000534"}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	links, err := c.GetLinkBugs(context.Background(), &model.GetLinkBugsRequest{
		WorkspaceID: "10104801",
		BugID:       "1001",
	})
	if err != nil {
		t.Fatalf("GetLinkBugs() unexpected error: %v", err)
	}
	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}
	if links[0].Type != "repeat" {
		t.Errorf("type = %q, want %q", links[0].Type, "repeat")
	}
	if links[0].ID != "1010104801085894269" {
		t.Errorf("id = %q, want %q", links[0].ID, "1010104801085894269")
	}
	if links[0].WorkspaceID != "10104801" {
		t.Errorf("workspace_id = %q, want %q", links[0].WorkspaceID, "10104801")
	}
	if links[0].Actas != "target" {
		t.Errorf("actas = %q, want %q", links[0].Actas, "target")
	}
	if links[0].LinkedWorkspaceID != 10104801 {
		t.Errorf("linked_workspace_id = %d, want %d", links[0].LinkedWorkspaceID, 10104801)
	}
	if links[0].LinkID != "1162187798001000534" {
		t.Errorf("link_id = %q, want %q", links[0].LinkID, "1162187798001000534")
	}
}

func TestGetBugRelatedStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/get_related_stories" {
			t.Errorf("unexpected path: %s, want /bugs/get_related_stories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"workspace_id":"10104801","bug_id":"1010104801083691309","story_id":"1010104801866181263"}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	rels, err := c.GetBugRelatedStories(context.Background(), &model.GetBugRelatedStoriesRequest{
		WorkspaceID: "10104801",
		BugID:       "1010104801083691309",
	})
	if err != nil {
		t.Fatalf("GetBugRelatedStories() unexpected error: %v", err)
	}
	if len(rels) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(rels))
	}
	if rels[0].WorkspaceID != "10104801" {
		t.Errorf("workspace_id = %q, want %q", rels[0].WorkspaceID, "10104801")
	}
	if rels[0].BugID != "1010104801083691309" {
		t.Errorf("bug_id = %q, want %q", rels[0].BugID, "1010104801083691309")
	}
	if rels[0].StoryID != "1010104801866181263" {
		t.Errorf("story_id = %q, want %q", rels[0].StoryID, "1010104801866181263")
	}
}

func TestListBugTemplates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/template_list" {
			t.Errorf("unexpected path: %s, want /bugs/template_list", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"WorkitemTemplate":{"id":"1010104801000068639","name":"创建模板","description":"AA","sort":"1","default":"0","creator":"v_xuanfang","editor_type":"1"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	templates, err := c.ListBugTemplates(context.Background(), &model.WorkspaceIDRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("ListBugTemplates() unexpected error: %v", err)
	}
	if len(templates) != 1 {
		t.Fatalf("expected 1 template, got %d", len(templates))
	}
	if templates[0].ID != "1010104801000068639" {
		t.Errorf("id = %q, want %q", templates[0].ID, "1010104801000068639")
	}
	if templates[0].Name != "创建模板" {
		t.Errorf("name = %q, want %q", templates[0].Name, "创建模板")
	}
}

func TestGetDefaultBugTemplate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/get_default_bug_template" {
			t.Errorf("unexpected path: %s, want /bugs/get_default_bug_template", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"WorkitemTemplateField":{"id":"1010104801000778831","workspace_id":"10104801","type":"bug","template_id":"1010104801000068639","field":"title","value":"","required":"1","sort":"0"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	fields, err := c.GetDefaultBugTemplate(context.Background(), &model.GetDefaultBugTemplateRequest{
		WorkspaceID: "10104801",
		TemplateID:  "1010104801000068639",
	})
	if err != nil {
		t.Fatalf("GetDefaultBugTemplate() unexpected error: %v", err)
	}
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	if fields[0].ID != "1010104801000778831" {
		t.Errorf("id = %q, want %q", fields[0].ID, "1010104801000778831")
	}
	if fields[0].Field != "title" {
		t.Errorf("field = %q, want %q", fields[0].Field, "title")
	}
	if fields[0].Required != "1" {
		t.Errorf("required = %q, want %q", fields[0].Required, "1")
	}
}

func TestGetBugCustomFieldsSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/custom_fields_settings" {
			t.Errorf("unexpected path: %s, want /bugs/custom_fields_settings", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"CustomFieldConfig":{"id":"1010158231077902981","workspace_id":"10158231","entry_type":"bug","custom_field":"custom_field_one","type":"radio","name":"安全漏洞类型","options":"XSS注入|SQL注入|越权","enabled":"1","sort":"1"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	configs, err := c.GetBugCustomFieldsSettings(context.Background(), &model.WorkspaceIDRequest{WorkspaceID: "10158231"})
	if err != nil {
		t.Fatalf("GetBugCustomFieldsSettings() unexpected error: %v", err)
	}
	if len(configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(configs))
	}
	if configs[0].ID != "1010158231077902981" {
		t.Errorf("id = %q, want %q", configs[0].ID, "1010158231077902981")
	}
	if configs[0].CustomField != "custom_field_one" {
		t.Errorf("custom_field = %q, want %q", configs[0].CustomField, "custom_field_one")
	}
	if configs[0].Name != "安全漏洞类型" {
		t.Errorf("name = %q, want %q", configs[0].Name, "安全漏洞类型")
	}
	if configs[0].EntryType != "bug" {
		t.Errorf("entry_type = %q, want %q", configs[0].EntryType, "bug")
	}
}

func TestBugFilterToQueryToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/filter_to_query_token" {
			t.Errorf("unexpected path: %s, want /bugs/filter_to_query_token", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"queryToken":"71ab88eeb45d084d8fbc85686a0d2399","href":"http://www.tapd.cn/tapd_fe/10104801/bug/list?queryToken=71ab88eeb45d084d8fbc85686a0d2399"},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.BugFilterToQueryToken(context.Background(), &model.FilterToQueryTokenRequest{
		WorkspaceID: "10104801",
		Filters:     map[string]string{"status": "new"},
	})
	if err != nil {
		t.Fatalf("BugFilterToQueryToken() unexpected error: %v", err)
	}
	if result.QueryToken != "71ab88eeb45d084d8fbc85686a0d2399" {
		t.Errorf("queryToken = %q, want %q", result.QueryToken, "71ab88eeb45d084d8fbc85686a0d2399")
	}
	if result.Href == "" {
		t.Error("href should not be empty")
	}
}

func TestCopyBug(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/copy_bug" {
			t.Errorf("unexpected path: %s, want /bugs/copy_bug", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"Bug":{"id":"1010104801500664637","title":"同步复制-源","workspace_id":"10104801","status":"rejected","reporter":"anyechen","created":"2020-11-26 15:22:34","current_owner":"anyechen"}},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	bug, err := c.CopyBug(context.Background(), &model.CopyBugRequest{
		WorkspaceID:    "10104801",
		SrcBugID:       "1001",
		DstWorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("CopyBug() unexpected error: %v", err)
	}
	if bug.ID != "1010104801500664637" {
		t.Errorf("id = %q, want %q", bug.ID, "1010104801500664637")
	}
	if bug.Title != "同步复制-源" {
		t.Errorf("title = %q, want %q", bug.Title, "同步复制-源")
	}
	if bug.Status != "rejected" {
		t.Errorf("status = %q, want %q", bug.Status, "rejected")
	}
	if bug.Reporter != "anyechen" {
		t.Errorf("reporter = %q, want %q", bug.Reporter, "anyechen")
	}
}

func TestBatchUpdateBug(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/batch_update_bug" {
			t.Errorf("unexpected path: %s, want /bugs/batch_update_bug", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"msg":"batch update success"},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.BatchUpdateBug(context.Background(), &model.BatchUpdateBugRequest{
		WorkspaceID: "10104801",
		Workitems: []model.BatchUpdateBugItem{
			{ID: "1001", Status: "resolved"},
			{ID: "1002", Status: "closed"},
		},
	})
	if err != nil {
		t.Fatalf("BatchUpdateBug() unexpected error: %v", err)
	}
}

func TestGetRemovedBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/get_removed_bugs" {
			t.Errorf("unexpected path: %s, want /bugs/get_removed_bugs", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"RemovedBug":{"id":"1100000755500695186","title":"标题呀","reporter":"gobichen","created":"2021-04-22 21:29:41","operation_user":"v_tingtdong","modified":"2021-04-23 11:04:59","removed_comment":"{\"action\":\"delete\"}","type":"delete","new_bug_url":""}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	bugs, err := c.GetRemovedBugs(context.Background(), &model.GetRemovedBugsRequest{WorkspaceID: "10104801"})
	if err != nil {
		t.Fatalf("GetRemovedBugs() unexpected error: %v", err)
	}
	if len(bugs) != 1 {
		t.Fatalf("expected 1 removed bug, got %d", len(bugs))
	}
	if bugs[0].ID != "1100000755500695186" {
		t.Errorf("id = %q, want %q", bugs[0].ID, "1100000755500695186")
	}
	if bugs[0].Title != "标题呀" {
		t.Errorf("title = %q, want %q", bugs[0].Title, "标题呀")
	}
	if bugs[0].Reporter != "gobichen" {
		t.Errorf("reporter = %q, want %q", bugs[0].Reporter, "gobichen")
	}
	if bugs[0].Type != "delete" {
		t.Errorf("type = %q, want %q", bugs[0].Type, "delete")
	}
	if bugs[0].OperationUser != "v_tingtdong" {
		t.Errorf("operation_user = %q, want %q", bugs[0].OperationUser, "v_tingtdong")
	}
}

func TestGetBugsByViewConfID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/get_bugs_by_view_conf_id" {
			t.Errorf("unexpected path: %s, want /bugs/get_bugs_by_view_conf_id", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"Bug":{"id":"1010104801084955735","title":"test","status":"new","current_owner":"","created":"2021-01-21 11:02:09","reporter":"v_xuanfang"}}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	bugs, err := c.GetBugsByViewConfID(context.Background(), &model.GetBugsByViewConfIDRequest{
		WorkspaceID: "10104801",
		ViewConfID:  "view001",
	})
	if err != nil {
		t.Fatalf("GetBugsByViewConfID() unexpected error: %v", err)
	}
	if len(bugs) != 1 {
		t.Fatalf("expected 1 bug, got %d", len(bugs))
	}
	if bugs[0].ID != "1010104801084955735" {
		t.Errorf("id = %q, want %q", bugs[0].ID, "1010104801084955735")
	}
	if bugs[0].Title != "test" {
		t.Errorf("title = %q, want %q", bugs[0].Title, "test")
	}
	if bugs[0].Status != "new" {
		t.Errorf("status = %q, want %q", bugs[0].Status, "new")
	}
	if bugs[0].Reporter != "v_xuanfang" {
		t.Errorf("reporter = %q, want %q", bugs[0].Reporter, "v_xuanfang")
	}
}

func TestBugIDsToQueryToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/ids_to_query_token" {
			t.Errorf("unexpected path: %s, want /bugs/ids_to_query_token", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":{"queryToken":"71ab88eeb45d084d8fbc85686a0d2399","href":"http://www.tapd.cn/tapd_fe/10104801/bug/list?queryToken=71ab88eeb45d084d8fbc85686a0d2399"},"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.BugIDsToQueryToken(context.Background(), &model.BugIDsToQueryTokenRequest{
		WorkspaceID: "10104801",
		IDs:         "1001,1002,1003",
	})
	if err != nil {
		t.Fatalf("BugIDsToQueryToken() unexpected error: %v", err)
	}
	if result.QueryToken != "71ab88eeb45d084d8fbc85686a0d2399" {
		t.Errorf("queryToken = %q, want %q", result.QueryToken, "71ab88eeb45d084d8fbc85686a0d2399")
	}
	if result.Href == "" {
		t.Error("href should not be empty")
	}
}

func TestUpdateBugSystemSelectFieldOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/bugs/update_system_select_field_options" {
			t.Errorf("unexpected path: %s, want /bugs/update_system_select_field_options", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":true,"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.UpdateBugSystemSelectFieldOptions(context.Background(), &model.UpdateBugSystemSelectFieldOptionsRequest{
		WorkspaceID: "10104801",
		Field:       "severity",
		Options:     `["critical","major","minor"]`,
		Value:       "critical",
	})
	if err != nil {
		t.Fatalf("UpdateBugSystemSelectFieldOptions() unexpected error: %v", err)
	}
}
