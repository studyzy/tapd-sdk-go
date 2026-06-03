package tapd

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestListWorkspaces(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/user_participant_projects" {
			t.Errorf("unexpected path: %s, want /workspaces/user_participant_projects", r.URL.Path)
		}
		if r.URL.Query().Get("nick") != "testuser" {
			t.Errorf("nick = %q, want %q", r.URL.Query().Get("nick"), "testuser")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Workspace":{"id":"1","name":"Project1","status":"active","category":"project"}},{"Workspace":{"id":"2","name":"Org","status":"active","category":"organization"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	workspaces, err := c.ListWorkspaces(context.Background(), "12345", "testuser")
	if err != nil {
		t.Fatalf("ListWorkspaces() unexpected error: %v", err)
	}
	if len(workspaces) != 2 {
		t.Fatalf("expected 2 workspaces, got %d", len(workspaces))
	}
	if workspaces[0].ID != "1" {
		t.Errorf("workspace ID = %q, want %q", workspaces[0].ID, "1")
	}
	if workspaces[0].Name != "Project1" {
		t.Errorf("workspace Name = %q, want %q", workspaces[0].Name, "Project1")
	}
}

func TestGetWorkspaceInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/get_workspace_info" {
			t.Errorf("unexpected path: %s, want /workspaces/get_workspace_info", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Workspace":{"id":"10","name":"MyProject","status":"active","category":"project","creator":"admin"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ws, err := c.GetWorkspaceInfo(context.Background(), "10")
	if err != nil {
		t.Fatalf("GetWorkspaceInfo() unexpected error: %v", err)
	}
	if ws.ID != "10" {
		t.Errorf("workspace ID = %q, want %q", ws.ID, "10")
	}
	if ws.Name != "MyProject" {
		t.Errorf("workspace Name = %q, want %q", ws.Name, "MyProject")
	}
	if ws.Status != "active" {
		t.Errorf("workspace Status = %q, want %q", ws.Status, "active")
	}
	if ws.Creator != "admin" {
		t.Errorf("workspace Creator = %q, want %q", ws.Creator, "admin")
	}
}

func TestGetWorkspaceInfo_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.GetWorkspaceInfo(context.Background(), "999")
	if err == nil {
		t.Fatal("GetWorkspaceInfo() expected error for empty data, got nil")
	}
	var tapdErr *TAPDError
	if !errors.As(err, &tapdErr) {
		t.Fatalf("expected *TAPDError, got %T", err)
	}
	if tapdErr.ExitCode != 2 {
		t.Errorf("ExitCode = %d, want 2", tapdErr.ExitCode)
	}
}

func TestGetSubWorkspaces(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/sub_workspaces" {
			t.Errorf("unexpected path: %s, want /workspaces/sub_workspaces", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Workspace":{"id":"10104801","name":"TAPD 乌云","pretty_name":"tapd_security","category":"product","status":"normal"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	ws, err := c.GetSubWorkspaces(context.Background(), &model.GetSubWorkspacesRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("GetSubWorkspaces() unexpected error: %v", err)
	}
	if ws.ID != "10104801" {
		t.Errorf("workspace ID = %q, want %q", ws.ID, "10104801")
	}
	if ws.Name != "TAPD 乌云" {
		t.Errorf("workspace Name = %q, want %q", ws.Name, "TAPD 乌云")
	}
	if ws.PrettyName != "tapd_security" {
		t.Errorf("workspace PrettyName = %q, want %q", ws.PrettyName, "tapd_security")
	}
	if ws.Category != "product" {
		t.Errorf("workspace Category = %q, want %q", ws.Category, "product")
	}
}

func TestListCompanyProjects(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/projects" {
			t.Errorf("unexpected path: %s, want /workspaces/projects", r.URL.Path)
		}
		if r.URL.Query().Get("company_id") != "12345" {
			t.Errorf("company_id = %q, want %q", r.URL.Query().Get("company_id"), "12345")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Workspace":{"id":"20026861","name":"产品运营2015","pretty_name":"20026861","status":"normal","category":"project","created":"2016-03-10 17:01:45"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	projects, err := c.ListCompanyProjects(context.Background(), &model.ListCompanyProjectsRequest{
		CompanyID: "12345",
	})
	if err != nil {
		t.Fatalf("ListCompanyProjects() unexpected error: %v", err)
	}
	if len(projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(projects))
	}
	if projects[0].ID != "20026861" {
		t.Errorf("project ID = %q, want %q", projects[0].ID, "20026861")
	}
	if projects[0].Name != "产品运营2015" {
		t.Errorf("project Name = %q, want %q", projects[0].Name, "产品运营2015")
	}
	if projects[0].Category != "project" {
		t.Errorf("project Category = %q, want %q", projects[0].Category, "project")
	}
}

func TestGetWorkspaceUsers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/users" {
			t.Errorf("unexpected path: %s, want /workspaces/users", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"UserWorkspace":{"user":"anyechen","role_id":["1000000000000000002","1000000000000000009"],"name":"陈安业","join_project_time":"2024-09-10","status":"1","allocation":"100"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	users, err := c.GetWorkspaceUsers(context.Background(), &model.GetWorkspaceUsersRequest{
		WorkspaceID: "10104801",
	})
	if err != nil {
		t.Fatalf("GetWorkspaceUsers() unexpected error: %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
	if users[0].User != "anyechen" {
		t.Errorf("user = %q, want %q", users[0].User, "anyechen")
	}
	if users[0].Name != "陈安业" {
		t.Errorf("name = %q, want %q", users[0].Name, "陈安业")
	}
	if len(users[0].RoleID) != 2 {
		t.Fatalf("expected 2 role IDs, got %d", len(users[0].RoleID))
	}
	if users[0].RoleID[0] != "1000000000000000002" {
		t.Errorf("role_id[0] = %q, want %q", users[0].RoleID[0], "1000000000000000002")
	}
	if users[0].Allocation != "100" {
		t.Errorf("allocation = %q, want %q", users[0].Allocation, "100")
	}
}

func TestAddWorkspaceMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/add_workspace_member" {
			t.Errorf("unexpected path: %s, want /workspaces/add_workspace_member", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":true},"info":"add member davidning to 10104801 success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.AddWorkspaceMember(context.Background(), &model.AddWorkspaceMemberRequest{
		WorkspaceID: "10104801",
		Nick:        "davidning",
	})
	if err != nil {
		t.Fatalf("AddWorkspaceMember() unexpected error: %v", err)
	}
	if !result.Success {
		t.Errorf("expected success=true, got false")
	}
}

func TestGetRoles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/roles" {
			t.Errorf("unexpected path: %s, want /roles", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "10104801" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "10104801")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"1000000000000000002":"管理员","1000000000000000009":"开发人员","1000000000000000010":"测试人员"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	roles, err := c.GetRoles(context.Background(), "10104801")
	if err != nil {
		t.Fatalf("GetRoles() unexpected error: %v", err)
	}
	if len(roles) != 3 {
		t.Fatalf("expected 3 roles, got %d", len(roles))
	}
	if roles["1000000000000000002"] != "管理员" {
		t.Errorf("role 1000000000000000002 = %q, want %q", roles["1000000000000000002"], "管理员")
	}
	if roles["1000000000000000009"] != "开发人员" {
		t.Errorf("role 1000000000000000009 = %q, want %q", roles["1000000000000000009"], "开发人员")
	}
	if roles["1000000000000000010"] != "测试人员" {
		t.Errorf("role 1000000000000000010 = %q, want %q", roles["1000000000000000010"], "测试人员")
	}
}

func TestCreateMiniProject(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/create_mini_project" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"workspace_id":"12345","workspace_url":"https://www.tapd.cn/12345"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.CreateMiniProject(context.Background(), &model.CreateMiniProjectRequest{
		CompanyID: "1",
		Name:      "Test Project",
		Creator:   "admin",
	})
	if err != nil {
		t.Fatalf("CreateMiniProject() unexpected error: %v", err)
	}
	if result.WorkspaceID != "12345" {
		t.Errorf("workspace_id = %q, want %q", result.WorkspaceID, "12345")
	}
	if result.WorkspaceURL != "https://www.tapd.cn/12345" {
		t.Errorf("workspace_url = %q, want %q", result.WorkspaceURL, "https://www.tapd.cn/12345")
	}
}

func TestGetMiniProjectList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workspaces/get_mini_project_list_with_permission" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"id":"100","name":"Space1","status":"normal","creator":"admin"},{"id":"101","name":"Space2","status":"normal","creator":"user1"}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	projects, err := c.GetMiniProjectList(context.Background(), &model.GetMiniProjectListRequest{
		User:      "admin",
		CompanyID: "1",
	})
	if err != nil {
		t.Fatalf("GetMiniProjectList() unexpected error: %v", err)
	}
	if len(projects) != 2 {
		t.Fatalf("expected 2 projects, got %d", len(projects))
	}
	if projects[0].ID != "100" {
		t.Errorf("first project id = %q, want %q", projects[0].ID, "100")
	}
	if projects[1].Name != "Space2" {
		t.Errorf("second project name = %q, want %q", projects[1].Name, "Space2")
	}
}

func TestEnableWorkCalendar(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/enable_work_calendar" {
			t.Errorf("unexpected path: %s, want /workspaces/enable_work_calendar", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.EnableWorkCalendar(context.Background(), &model.EnableWorkCalendarRequest{
		WorkspaceID: "48464494",
		Type:        "system",
	})
	if err != nil {
		t.Fatalf("EnableWorkCalendar() unexpected error: %v", err)
	}
	if !result.Success {
		t.Errorf("expected success=true, got false")
	}
}

func TestGetCustomWorkCalendar(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/get_custom_work_calendar" {
			t.Errorf("unexpected path: %s, want /workspaces/get_custom_work_calendar", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "48464494" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "48464494")
		}
		if r.URL.Query().Get("year") != "2025" {
			t.Errorf("year = %q, want %q", r.URL.Query().Get("year"), "2025")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"weekdays":["1","2","3","4","5","6","7"],"holidays":["2025-01-01"],"workdays":["2025-01-02","2025-01-03","2025-01-04"]},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	cal, err := c.GetCustomWorkCalendar(context.Background(), &model.GetCustomWorkCalendarRequest{
		WorkspaceID: "48464494",
		Year:        "2025",
	})
	if err != nil {
		t.Fatalf("GetCustomWorkCalendar() unexpected error: %v", err)
	}
	if len(cal.Weekdays) != 7 {
		t.Fatalf("expected 7 weekdays, got %d", len(cal.Weekdays))
	}
	if cal.Weekdays[0] != "1" {
		t.Errorf("weekdays[0] = %q, want %q", cal.Weekdays[0], "1")
	}
	if len(cal.Holidays) != 1 || cal.Holidays[0] != "2025-01-01" {
		t.Errorf("holidays = %v, want [2025-01-01]", cal.Holidays)
	}
	if len(cal.Workdays) != 3 {
		t.Errorf("expected 3 workdays, got %d", len(cal.Workdays))
	}
}

func TestSetCustomWorkCalendar(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/set_custom_work_calendar" {
			t.Errorf("unexpected path: %s, want /workspaces/set_custom_work_calendar", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.SetCustomWorkCalendar(context.Background(), &model.SetCustomWorkCalendarRequest{
		WorkspaceID: "48464494",
		Year:        "2025",
		Weekdays:    []string{"1", "2", "3", "4", "5"},
		Holidays:    []string{"2025-01-01"},
		Workdays:    []string{"2025-01-04"},
	})
	if err != nil {
		t.Fatalf("SetCustomWorkCalendar() unexpected error: %v", err)
	}
	if !result.Success {
		t.Errorf("expected success=true, got false")
	}
}

func TestGetWorkCalendarSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/get_work_calendar_settings" {
			t.Errorf("unexpected path: %s, want /workspaces/get_work_calendar_settings", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "48464494" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "48464494")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"name":"中国大陆法定工作日","type":"system","enable":true},{"name":"自定义工作日","type":"custom","enable":false}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	settings, err := c.GetWorkCalendarSettings(context.Background(), "48464494")
	if err != nil {
		t.Fatalf("GetWorkCalendarSettings() unexpected error: %v", err)
	}
	if len(settings) != 2 {
		t.Fatalf("expected 2 settings, got %d", len(settings))
	}
	if settings[0].Name != "中国大陆法定工作日" {
		t.Errorf("settings[0].Name = %q, want %q", settings[0].Name, "中国大陆法定工作日")
	}
	if settings[0].Type != "system" {
		t.Errorf("settings[0].Type = %q, want %q", settings[0].Type, "system")
	}
	if !settings[0].Enable {
		t.Errorf("settings[0].Enable = false, want true")
	}
	if settings[1].Enable {
		t.Errorf("settings[1].Enable = true, want false")
	}
}

func TestGetWorkitemsLongIDByShortIDs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/get_workitems_long_id_by_short_ids" {
			t.Errorf("unexpected path: %s, want /workspaces/get_workitems_long_id_by_short_ids", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "48464494" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "48464494")
		}
		if r.URL.Query().Get("entity_type") != "story" {
			t.Errorf("entity_type = %q, want %q", r.URL.Query().Get("entity_type"), "story")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"valid_id_map":[{"short_id":"1000276","long_id":"1148464494001000276","entity_type":"story","workspace_id":"48464494","company_id":"39418254"},{"short_id":"1000277","long_id":"1148464494001000277","entity_type":"story","workspace_id":"48464494","company_id":"39418254"}],"invalid_long_ids":["1000104"],"invalid_short_ids":[]},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.GetWorkitemsLongIDByShortIDs(context.Background(), &model.GetWorkitemsLongIDByShortIDsRequest{
		WorkspaceID: "48464494",
		EntityType:  "story",
		ShortIDs:    "1000276;1000277;1000104",
	})
	if err != nil {
		t.Fatalf("GetWorkitemsLongIDByShortIDs() unexpected error: %v", err)
	}
	if len(result.ValidIDMap) != 2 {
		t.Fatalf("expected 2 valid id maps, got %d", len(result.ValidIDMap))
	}
	if result.ValidIDMap[0].ShortID != "1000276" {
		t.Errorf("ValidIDMap[0].ShortID = %q, want %q", result.ValidIDMap[0].ShortID, "1000276")
	}
	if result.ValidIDMap[0].LongID != "1148464494001000276" {
		t.Errorf("ValidIDMap[0].LongID = %q, want %q", result.ValidIDMap[0].LongID, "1148464494001000276")
	}
	if result.ValidIDMap[1].CompanyID != "39418254" {
		t.Errorf("ValidIDMap[1].CompanyID = %q, want %q", result.ValidIDMap[1].CompanyID, "39418254")
	}
	if len(result.InvalidLongIDs) != 1 || result.InvalidLongIDs[0] != "1000104" {
		t.Errorf("InvalidLongIDs = %v, want [1000104]", result.InvalidLongIDs)
	}
}

func TestGetWorkspaceDocuments(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/documents/get_workspace_documents" {
			t.Errorf("unexpected path: %s, want /documents/get_workspace_documents", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "47043561" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "47043561")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Document":{"id":"1147043561001001330","workspace_id":"47043561","name":"熟悉思维导图","type":"mindmap","folder_id":"1147043561001000694","creator":"TAPD","modifier":"TAPD","created":"2021-09-09 16:08:52","modified":"2021-09-09 16:08:52"}},{"Document":{"id":"1147043561001001329","workspace_id":"47043561","name":"文档功能使用秘籍","type":"word","folder_id":"1147043561001000694","creator":"TAPD","modifier":"TAPD","created":"2021-09-09 16:08:51","modified":"2021-09-09 16:08:51"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	docs, err := c.GetWorkspaceDocuments(context.Background(), &model.GetWorkspaceDocumentsRequest{
		WorkspaceID: "47043561",
	})
	if err != nil {
		t.Fatalf("GetWorkspaceDocuments() unexpected error: %v", err)
	}
	if len(docs) != 2 {
		t.Fatalf("expected 2 documents, got %d", len(docs))
	}
	if docs[0].ID != "1147043561001001330" {
		t.Errorf("docs[0].ID = %q, want %q", docs[0].ID, "1147043561001001330")
	}
	if docs[0].Name != "熟悉思维导图" {
		t.Errorf("docs[0].Name = %q, want %q", docs[0].Name, "熟悉思维导图")
	}
	if docs[0].Type != "mindmap" {
		t.Errorf("docs[0].Type = %q, want %q", docs[0].Type, "mindmap")
	}
	if docs[1].Name != "文档功能使用秘籍" {
		t.Errorf("docs[1].Name = %q, want %q", docs[1].Name, "文档功能使用秘籍")
	}
	if docs[1].Creator != "TAPD" {
		t.Errorf("docs[1].Creator = %q, want %q", docs[1].Creator, "TAPD")
	}
}

func TestUpdateWorkspaceInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/update_workspace_info" {
			t.Errorf("unexpected path: %s, want /workspaces/update_workspace_info", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":"update workspace success","info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	msg, err := c.UpdateWorkspaceInfo(context.Background(), &model.UpdateWorkspaceInfoRequest{
		WorkspaceID: "69999237",
		Field:       "end_date",
		Value:       "2025-03-03",
	})
	if err != nil {
		t.Fatalf("UpdateWorkspaceInfo() unexpected error: %v", err)
	}
	if msg != "update workspace success" {
		t.Errorf("message = %q, want %q", msg, "update workspace success")
	}
}

func TestGetWorkspaceCustomFieldSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/workspaces/workspace_custom_field_settings" {
			t.Errorf("unexpected path: %s, want /workspaces/workspace_custom_field_settings", r.URL.Path)
		}
		if r.URL.Query().Get("workspace_id") != "20001871" {
			t.Errorf("workspace_id = %q, want %q", r.URL.Query().Get("workspace_id"), "20001871")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"CustomFieldConfig":{"id":"1010158231215016293","workspace_id":"20001871","entry_type":"story","custom_field":"custom_field_three","type":"text","name":"自定义字段1","options":"","enabled":"1"}},{"CustomFieldConfig":{"id":"1010158231215016295","workspace_id":"20001871","entry_type":"bug","custom_field":"custom_field_four","type":"select","name":"自定义字段2","options":"{\"1\":\"选项A\",\"2\":\"选项B\"}","enabled":"1"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	configs, err := c.GetWorkspaceCustomFieldSettings(context.Background(), "20001871")
	if err != nil {
		t.Fatalf("GetWorkspaceCustomFieldSettings() unexpected error: %v", err)
	}
	if len(configs) != 2 {
		t.Fatalf("expected 2 configs, got %d", len(configs))
	}
	if configs[0].ID != "1010158231215016293" {
		t.Errorf("configs[0].ID = %q, want %q", configs[0].ID, "1010158231215016293")
	}
	if configs[0].EntryType != "story" {
		t.Errorf("configs[0].EntryType = %q, want %q", configs[0].EntryType, "story")
	}
	if configs[0].Name != "自定义字段1" {
		t.Errorf("configs[0].Name = %q, want %q", configs[0].Name, "自定义字段1")
	}
	if configs[1].Type != "select" {
		t.Errorf("configs[1].Type = %q, want %q", configs[1].Type, "select")
	}
	if configs[1].Enabled != "1" {
		t.Errorf("configs[1].Enabled = %q, want %q", configs[1].Enabled, "1")
	}
}
