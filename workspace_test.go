package tapd

import (
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Workspace":{"id":"1","name":"Project1","status":"active","category":"project"}},{"Workspace":{"id":"2","name":"Org","status":"active","category":"organization"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	workspaces, err := c.ListWorkspaces()
	if err != nil {
		t.Fatalf("ListWorkspaces() unexpected error: %v", err)
	}
	if len(workspaces) != 1 {
		t.Fatalf("expected 1 workspace after filtering organization, got %d", len(workspaces))
	}
	if workspaces[0].ID != "1" {
		t.Errorf("workspace ID = %q, want %q", workspaces[0].ID, "1")
	}
	if workspaces[0].Name != "Project1" {
		t.Errorf("workspace Name = %q, want %q", workspaces[0].Name, "Project1")
	}
	if workspaces[0].Category != "project" {
		t.Errorf("workspace Category = %q, want %q", workspaces[0].Category, "project")
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
	ws, err := c.GetWorkspaceInfo("10")
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
	_, err := c.GetWorkspaceInfo("999")
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
	ws, err := c.GetSubWorkspaces(&model.GetSubWorkspacesRequest{
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
	projects, err := c.ListCompanyProjects(&model.ListCompanyProjectsRequest{
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
	users, err := c.GetWorkspaceUsers(&model.GetWorkspaceUsersRequest{
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
	result, err := c.AddWorkspaceMember(&model.AddWorkspaceMemberRequest{
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
	roles, err := c.GetRoles("10104801")
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
