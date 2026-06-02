package tapd

import (
	"context"
	"net/http"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestTestxListReports(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/report/v1/namespaces/ns1/reports" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"rpt1","Title":"report1"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListReports(context.Background(), &model.TestxListReportsRequest{
		Namespace: "ns1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d reports, want 1", len(result))
	}
	if result[0].Uid != "rpt1" {
		t.Errorf("got Uid=%q, want rpt1", result[0].Uid)
	}
}

func TestTestxGetReport(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/report/v1/namespaces/ns1/reports/rpt1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"rpt1","Title":"my report","Summary":"summary"}]}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxGetReport(context.Background(), &model.TestxGetReportRequest{
		Namespace: "ns1",
		Uid:       "rpt1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Title != "my report" {
		t.Errorf("got Title=%q, want my report", result.Title)
	}
}

func TestTestxGetReportData(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/report/v1/namespaces/ns1/reports/rpt1/templates/tpl1/data" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":{"key":"value"}}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxGetReportData(context.Background(), &model.TestxGetReportDataRequest{
		Namespace:   "ns1",
		ReportUid:   "rpt1",
		TemplateUid: "tpl1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if len(result.Raw) == 0 {
		t.Error("expected non-empty Raw data")
	}
}

func TestTestxListReportTemplates(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/report/v1/namespaces/ns1/templates" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"Uid":"tpl1","Title":"template1"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListReportTemplates(context.Background(), &model.TestxListReportTemplatesRequest{
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
