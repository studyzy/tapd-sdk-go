package tapd

import (
	"context"
	"net/http"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestTestxCreatePlanFolder(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/folders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"pf1","Name":"plan folder"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxCreatePlanFolder(context.Background(), &model.TestxCreatePlanFolderRequest{
		Namespace: "ns1",
		ParentUid: "root",
		Name:      "plan folder",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "pf1" {
		t.Errorf("got Uid=%q, want pf1", result.Uid)
	}
}

func TestTestxCreatePlan(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Meta":{"Uid":"plan1","Name":"test plan"}}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxCreatePlan(context.Background(), &model.TestxCreatePlanRequest{
		Namespace: "ns1",
		Plan: &model.TestxPlan{
			Meta: &model.TestxPlanMeta{Name: "test plan"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Meta == nil || result.Meta.Uid != "plan1" {
		t.Errorf("got unexpected plan result: %+v", result)
	}
}

func TestTestxGetPlan(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Meta":{"Uid":"plan1","Name":"my plan","State":"running"}}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxGetPlan(context.Background(), &model.TestxGetPlanRequest{
		Namespace: "ns1",
		Uid:       "plan1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Meta == nil || result.Meta.State != "running" {
		t.Errorf("got unexpected plan state")
	}
}

func TestTestxListPlans(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/folders/folder1/plans-list" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Meta":{"Uid":"p1"}},{"Meta":{"Uid":"p2"}}],"TotalCount":2}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlans(context.Background(), &model.TestxListPlansRequest{
		Namespace: "ns1",
		FolderUid: "folder1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Errorf("got TotalCount=%d, want 2", total)
	}
	if len(result) != 2 {
		t.Errorf("got %d plans, want 2", len(result))
	}
}

func TestTestxListPlanCases(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/cases-search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"CaseUidToDetail":{"c1":{"State":"passed","CaseUid":"c1"}},"PlanCasesInfo":null},"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanCases(context.Background(), &model.TestxListPlanCasesRequest{
		Namespace: "ns1",
		Uid:       "plan1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if result.CaseUidToDetail == nil {
		t.Fatal("CaseUidToDetail is nil")
	}
	if _, ok := result.CaseUidToDetail["c1"]; !ok {
		t.Error("expected case c1 in CaseUidToDetail")
	}
}

func TestTestxPlanStatistics(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/statistics" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"plan1","Statistic":{"TotalCaseCount":10}}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxPlanStatistics(context.Background(), &model.TestxPlanStatisticsRequest{
		Namespace: "ns1",
		Uids:      []string{"plan1"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d items, want 1", len(result))
	}
	if result[0].Uid != "plan1" {
		t.Errorf("got Uid=%q, want plan1", result[0].Uid)
	}
}

func TestTestxListPlanBugs(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/bugs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Id":"bug1","Summary":"a plan bug"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanBugs(context.Background(), &model.TestxListPlanBugsRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d bugs, want 1", len(result))
	}
}

func TestTestxUpdatePlanFolder(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/folders/pf1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"pf1","Namespace":"ns1","Audit":{"Creator":"xxx","Updater":"xxx","CreatedAt":"2025-07-14T14:10:58+08:00","UpdatedAt":"2025-07-14T14:10:58+08:00","Tenant":""},"ParentUid":"0","Name":"test","Description":"test desc","PlanCount":0,"ArchiveAuto":false,"Folders":[],"Plans":[],"Path":"."}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxUpdatePlanFolder(context.Background(), &model.TestxUpdatePlanFolderRequest{
		Namespace:   "ns1",
		FolderUid:   "pf1",
		Name:        "test",
		Description: "test desc",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "pf1" {
		t.Errorf("got Uid=%q, want pf1", result.Uid)
	}
	if result.Name != "test" {
		t.Errorf("got Name=%q, want test", result.Name)
	}
}

func TestTestxListFolderChildren(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/folders/children" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Folders":[{"Uid":"sub1","Name":"test子目录","PlanCount":1}],"Plans":[{"Uid":"p1","Name":"test计划2"}]}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxListFolderChildren(context.Background(), &model.TestxListFolderChildrenRequest{
		Namespace: "ns1",
		Uid:       "pf1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Folders) != 1 {
		t.Errorf("got %d folders, want 1", len(result.Folders))
	}
	if len(result.Plans) != 1 {
		t.Errorf("got %d plans, want 1", len(result.Plans))
	}
}

func TestTestxUpdatePlan(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Meta":{"Uid":"plan1","Namespace":"ns1","Name":"测试","State":"WAITING","FolderUid":"xx"}}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxUpdatePlan(context.Background(), &model.TestxUpdatePlanRequest{
		Namespace: "ns1",
		Uid:       "plan1",
		Plan:      &model.TestxPlan{Meta: &model.TestxPlanMeta{Name: "测试"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Meta == nil || result.Meta.Uid != "plan1" {
		t.Errorf("got unexpected plan result")
	}
	if result.Meta.Name != "测试" {
		t.Errorf("got Name=%q, want 测试", result.Meta.Name)
	}
}

func TestTestxUpdatePlanTargetScope(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/target-scope" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Meta":{"Uid":"plan1","Namespace":"ns1","Name":"测试","State":"WAITING"}}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxUpdatePlanTargetScope(context.Background(), &model.TestxUpdatePlanTargetScopeRequest{
		Namespace: "ns1",
		Uid:       "plan1",
		Plan:      &model.TestxPlan{Meta: &model.TestxPlanMeta{Name: "测试"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Meta == nil || result.Meta.Uid != "plan1" {
		t.Errorf("got unexpected plan result")
	}
}

func TestTestxBatchUpdatePlanCase(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/cases/batch-update" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"RequestId":"","Error":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.TestxBatchUpdatePlanCase(context.Background(), &model.TestxBatchUpdatePlanCaseRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
		CaseInfos: []model.TestxBatchUpdatePlanCaseInfo{
			{State: "passed", CaseUid: "c1"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxBatchArchivePlan(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/batch-archive" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"RequestId":"","Error":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.TestxBatchArchivePlan(context.Background(), &model.TestxBatchArchivePlanRequest{
		Namespace:   "ns1",
		Uids:        []string{"plan1", "plan2"},
		ArchiveMode: "ARCHIVE",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxListPlanHistories(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/histories" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"h1","Creator":"Creator","CreateAt":"2025-07-16T17:12:10+08:00","PlanUid":"plan1","Fields":[{"Uid":"f1","Name":"测试字段一","PreValue":"test-pre","PostValue":"test"}]}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanHistories(context.Background(), &model.TestxListPlanHistoriesRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d histories, want 1", len(result))
	}
	if result[0].PlanUid != "plan1" {
		t.Errorf("got PlanUid=%q, want plan1", result[0].PlanUid)
	}
}

func TestTestxBatchBindPlanBug(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/cases/bugs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"RequestId":"","Error":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.TestxBatchBindPlanBug(context.Background(), &model.TestxBatchBindPlanBugRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
		CaseUids:  []string{"c1"},
		BindBugs:  []model.TestxIssue{{IssueUid: "bug1", WorkspaceUid: "ws1", Type: "BUG"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxUnbindPlanBug(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/cases/c1/issues/bug1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"RequestId":"","Error":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.TestxUnbindPlanBug(context.Background(), &model.TestxUnbindPlanBugRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
		CaseUid:   "c1",
		IssueUid:  "bug1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxListPlanCaseIssues(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/cases/c1/issues" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"IssueUid":"issue_uid","Namespace":"Namespace","WorkspaceUid":"WorkspaceUid","IssueUrl":"","Type":"BUG","Source":"NONE","Detail":null,"IssueName":"","IsDeleted":false,"Uid":""},{"IssueUid":"issue_uid_2","Namespace":"Namespace","WorkspaceUid":"WorkspaceUid","IssueUrl":"","Type":"BUG","Source":"NONE","Detail":null,"IssueName":"","IsDeleted":false,"Uid":""}],"TotalCount":2}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanCaseIssues(context.Background(), &model.TestxListPlanCaseIssuesRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
		CaseUid:   "c1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Errorf("got TotalCount=%d, want 2", total)
	}
	if len(result) != 2 {
		t.Errorf("got %d issues, want 2", len(result))
	}
	if result[0].IssueUid != "issue_uid" {
		t.Errorf("got IssueUid=%q, want issue_uid", result[0].IssueUid)
	}
}

func TestTestxListPlanCaseEvents(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/cases/c1/events" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Audit":{"Creator":"xxx","Updater":"xxx","CreatedAt":"2025-06-20T11:07:27+08:00","UpdatedAt":"2025-06-20T11:07:27+08:00","Tenant":""},"Type":"UPDATE_RESULT","Detail":"{\"type\":\"update_result\",\"status\":\"fail\"}","Source":"MANUAL","Attachments":[]},{"Audit":{"Creator":"xxx","Updater":"xxx","CreatedAt":"2025-06-20T11:07:25+08:00","UpdatedAt":"2025-06-20T11:07:25+08:00","Tenant":""},"Type":"UPDATE_RESULT","Detail":"{\"type\":\"update_result\",\"status\":\"succeed\"}","Source":"MANUAL","Attachments":[]},{"Audit":{"Creator":"xxx","Updater":"xxx","CreatedAt":"2025-06-20T11:07:17+08:00","UpdatedAt":"2025-06-20T11:07:17+08:00","Tenant":""},"Type":"UPDATE_TESTER","Detail":"{\"type\":\"update_tester\",\"tester\":\"xxx\"}","Source":"MANUAL","Attachments":[]}],"TotalCount":3}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanCaseEvents(context.Background(), &model.TestxListPlanCaseEventsRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
		CaseUid:   "c1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 3 {
		t.Errorf("got TotalCount=%d, want 3", total)
	}
	if len(result) != 3 {
		t.Errorf("got %d events, want 3", len(result))
	}
	if result[0].Type != "UPDATE_RESULT" {
		t.Errorf("got Type=%q, want UPDATE_RESULT", result[0].Type)
	}
}

func TestTestxListPlanBugStatistics(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/bug-statistics" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"PlanUid":"plan_uid","BugCount":0}]}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxListPlanBugStatistics(context.Background(), &model.TestxListPlanBugStatisticsRequest{
		Namespace: "ns1",
		PlanUids:  []string{"plan_uid"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("got %d items, want 1", len(result))
	}
	if result[0].PlanUid != "plan_uid" {
		t.Errorf("got PlanUid=%q, want plan_uid", result[0].PlanUid)
	}
}

func TestTestxListPlanStories(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plans/plan1/stories" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"IssueUid":"issue_uid","Namespace":"","WorkspaceUid":"workspace_uid","IssueUrl":"","Type":"STORY","Source":"NONE","Detail":null,"IssueName":"","IsDeleted":false,"Uid":""},{"IssueUid":"issue_uid_2","Namespace":"","WorkspaceUid":"workspace_uid","IssueUrl":"","Type":"STORY","Source":"NONE","Detail":null,"IssueName":"","IsDeleted":false,"Uid":""}],"TotalCount":2}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanStories(context.Background(), &model.TestxListPlanStoriesRequest{
		Namespace: "ns1",
		PlanUid:   "plan1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Errorf("got TotalCount=%d, want 2", total)
	}
	if len(result) != 2 {
		t.Errorf("got %d stories, want 2", len(result))
	}
	if result[0].Type != "STORY" {
		t.Errorf("got Type=%q, want STORY", result[0].Type)
	}
}

func TestTestxListPlanTemplates(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/plan/v1/namespaces/ns1/plan-templates" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"tpl1","Name":"plan template"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListPlanTemplates(context.Background(), &model.TestxListPlanTemplatesRequest{
		Namespace: "ns1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d templates, want 1", len(result))
	}
	if result[0].Uid != "tpl1" {
		t.Errorf("got Uid=%q, want tpl1", result[0].Uid)
	}
}
