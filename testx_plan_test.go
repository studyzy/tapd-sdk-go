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
