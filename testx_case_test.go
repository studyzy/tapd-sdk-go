package tapd

import (
	"context"
	"net/http"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestTestxCreateCaseRepo(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"123","Name":"test-repo","Type":"DEFALT"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxCreateCaseRepo(context.Background(), &model.TestxCreateCaseRepoRequest{
		Namespace: "ns1",
		Name:      "test-repo",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "123" {
		t.Errorf("got Uid=%q, want 123", result.Uid)
	}
}

func TestTestxGetCaseRepo(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"repo1","Name":"my-repo"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxGetCaseRepo(context.Background(), &model.TestxGetCaseRepoRequest{
		Namespace: "ns1",
		RepoUid:   "repo1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "my-repo" {
		t.Errorf("got Name=%q, want my-repo", result.Name)
	}
}

func TestTestxListCaseRepos(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Query().Get("Search") != "test" {
			t.Errorf("unexpected Search param: %s", r.URL.Query().Get("Search"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"r1","Name":"repo1"},{"Uid":"r2","Name":"repo2"}],"TotalCount":2}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListCaseRepos(context.Background(), &model.TestxListCaseReposRequest{
		Namespace: "ns1",
		Search:    "test",
		PageInfo:  &model.TestxPageInfo{Offset: 0, Limit: 10},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Errorf("got TotalCount=%d, want 2", total)
	}
	if len(result) != 2 {
		t.Errorf("got %d repos, want 2", len(result))
	}
}

func TestTestxCreateCaseFolder(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/folders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"f1","Name":"folder1"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxCreateCaseFolder(context.Background(), &model.TestxCreateCaseFolderRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		Folder: &model.TestxCreateCaseFolderRequestFolder{
			Name: "folder1",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "f1" {
		t.Errorf("got Uid=%q, want f1", result.Uid)
	}
}

func TestTestxCreateCase(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"case1","Name":"test case"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxCreateCase(context.Background(), &model.TestxCreateCaseRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		Case: &model.TestxCreateCaseRequestCase{
			Name:     "test case",
			Priority: "P1",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "case1" {
		t.Errorf("got Uid=%q, want case1", result.Uid)
	}
}

func TestTestxUpdateCase(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"case1","Name":"updated case"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxUpdateCase(context.Background(), &model.TestxUpdateCaseRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
		Case: &model.TestxCreateCaseRequestCase{
			Name:     "updated case",
			Priority: "P2",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "updated case" {
		t.Errorf("got Name=%q, want updated case", result.Name)
	}
}

func TestTestxSearchCases(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Cases":[{"Uid":"c1","Name":"case1"}],"Folders":[]},"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxSearchCases(context.Background(), &model.TestxSearchCasesRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		Filter:         &model.TestxSearchCasesFilter{Name: "case1"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result.Cases) != 1 {
		t.Errorf("got %d cases, want 1", len(result.Cases))
	}
}

func TestTestxBatchUpdateCases(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/batch-update" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.TestxBatchUpdateCases(context.Background(), &model.TestxBatchUpdateCasesRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUids:       []string{"c1", "c2"},
		UpdateInfos: []model.TestxBatchUpdateCaseInfo{
			{FieldName: "Priority", FieldValue: "P1", Action: "set"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxListCaseHistorys(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1/history" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"h1","ChangeType":"update"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListCaseHistorys(context.Background(), &model.TestxListCaseHistorysRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
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
}

func TestTestxListCaseBugs(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1/bugs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Bug":{"Id":"bug1","Summary":"a bug"}}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListCaseBugs(context.Background(), &model.TestxListCaseBugsRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
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

func TestTestxBatchBindCaseBug(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1/bugs/batch-bind" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":null}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	err := c.TestxBatchBindCaseBug(context.Background(), &model.TestxBatchBindCaseBugRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
		BindBugs: []model.TestxBindBug{
			{IssueUid: "bug1", WorkspaceUid: "ws1"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxUpdateCaseRepo(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Audit":{"Creator":"xx","Updater":"xx","CreatedAt":"","UpdatedAt":"","Tenant":"xxx"},"Namespace":"xxx","Uid":"17114","Name":"xxx","Description":"x","Type":"DEFALT"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxUpdateCaseRepo(context.Background(), &model.TestxUpdateCaseRepoRequest{
		Namespace:   "ns1",
		RepoUid:     "repo1",
		Name:        "xxx",
		Description: "x",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "17114" {
		t.Errorf("got Uid=%q, want 17114", result.Uid)
	}
	if result.Name != "xxx" {
		t.Errorf("got Name=%q, want xxx", result.Name)
	}
}

func TestTestxUpdateCaseFolder(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/folders/f1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"12572731","RepoUid":"","RepoVersionUid":"","FolderUid":"","FullPath":"","Name":"test","Owners":["xxxx"],"Description":"test","CaseCount":0,"UUID":"40340dec-cea7-4fe6-9a96-ee07eaef2bde","Path":".","Audit":{"Creator":"xx","Updater":"xx","CreatedAt":"2025-07-12T10:26:55+08:00","UpdatedAt":"2025-07-12T10:26:55+08:00","Tenant":"xx"}}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxUpdateCaseFolder(context.Background(), &model.TestxUpdateCaseFolderRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		FolderUid:      "f1",
		Folder: &model.TestxCreateCaseFolderRequestFolder{
			Name:        "test",
			Description: "test",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "12572731" {
		t.Errorf("got Uid=%q, want 12572731", result.Uid)
	}
	if result.Name != "test" {
		t.Errorf("got Name=%q, want test", result.Name)
	}
}

func TestTestxBatchCreateCases(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/batch-create" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"Uid":"12572733","RepoUid":"17099","RepoVersionUid":"18167","FolderUid":"12571827","FullPath":"/test/","UUID":"2f275e19-c053-4af9-b682-07d8c5ed13a6","Name":"测试","Description":"test desc","Priority":"P1","PreConditions":"test pre cond","Type":"44907","StepType":"STEP","Source":"TESTX","IsManualRelation":false,"Owners":["734242230"],"ManHourEstimated":"1","Path":".12571827.","RunTimes":"0","IsFolder":false}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxBatchCreateCases(context.Background(), &model.TestxBatchCreateCasesRequest{
		Namespace:       "ns1",
		RepoUid:         "repo1",
		RepoVersionUid:  "v1",
		TargetFolderUid: "12571827",
		Cases: []model.TestxCreateCaseRequestCase{
			{Name: "测试", Priority: "P1"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Uid != "12572733" {
		t.Errorf("got Uid=%q, want 12572733", result.Uid)
	}
	if result.Name != "测试" {
		t.Errorf("got Name=%q, want 测试", result.Name)
	}
}

func TestTestxListCaseExecutions(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1/executions" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"SourceName":"测试","SourceUid":"case-xx","Executor":"xx","StartExecuteTime":"","EndExecuteTime":"2025-07-01 17:02:59","ExecuteState":"FAIL","Message":"","BugCount":"0","LinkData":{"DesignUid":"xxxx"},"Source":"DESIGN"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListCaseExecutions(context.Background(), &model.TestxListCaseExecutionsRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d executions, want 1", len(result))
	}
	if result[0].ExecuteState != "FAIL" {
		t.Errorf("got ExecuteState=%q, want FAIL", result[0].ExecuteState)
	}
}

func TestTestxListCaseReviews(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1/reviews" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"SourceName":"测试复现","SourceUid":"xx.xxx","Reviewer":"734242230","ReviewTime":"2025-07-12 16:28:06","ReviewState":"REVIEW_RESULT_AGREE","Message":"11","LinkData":{"DesignUid":"xxx"},"Source":"DESIGN","Uid":"xxx","MainUid":"xxx","SourceKind":"CASE","Total":0,"CaseUid":"xxx"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListCaseReviews(context.Background(), &model.TestxListCaseReviewsRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d reviews, want 1", len(result))
	}
	if result[0].ReviewState != "REVIEW_RESULT_AGREE" {
		t.Errorf("got ReviewState=%q, want REVIEW_RESULT_AGREE", result[0].ReviewState)
	}
}

func TestTestxBatchUnbindCaseBug(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/repos/repo1/versions/v1/cases/case1/bugs/batch-unbind" {
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
	err := c.TestxBatchUnbindCaseBug(context.Background(), &model.TestxBatchUnbindCaseBugRequest{
		Namespace:      "ns1",
		RepoUid:        "repo1",
		RepoVersionUid: "v1",
		CaseUid:        "case1",
		BugUids:        []string{"bug1", "bug2"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTestxListCaseTemplates(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/case/v1/namespaces/ns1/case-templates" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"tpl1","Name":"template1"}]}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxListCaseTemplates(context.Background(), &model.TestxListCaseTemplatesRequest{
		Namespace: "ns1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("got %d templates, want 1", len(result))
	}
	if result[0].Uid != "tpl1" {
		t.Errorf("got Uid=%q, want tpl1", result[0].Uid)
	}
}
