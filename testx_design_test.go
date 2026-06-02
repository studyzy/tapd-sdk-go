package tapd

import (
	"context"
	"net/http"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestTestxSearchDesigns(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/design/v2/namespaces/ns1/designs/search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"meta":{"uid":"d1","name":"design1"}}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxSearchDesigns(context.Background(), &model.TestxSearchDesignsRequest{
		Namespace: "ns1",
		Filter:    &model.TestxSearchDesignsFilter{Search: "design1"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d designs, want 1", len(result))
	}
	if result[0].Meta == nil || result[0].Meta.Uid != "d1" {
		t.Error("unexpected design result")
	}
}

func TestTestxListDesignStats(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/design/v2/namespaces/ns1/stat-list" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"design_uid":"d1"}]}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.TestxListDesignStats(context.Background(), &model.TestxListDesignStatsRequest{
		Namespace:  "ns1",
		DesignUids: []string{"d1"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("got %d stats, want 1", len(result))
	}
	if result[0].DesignUid != "d1" {
		t.Errorf("got DesignUid=%q, want d1", result[0].DesignUid)
	}
}

func TestTestxListDesignLabels(t *testing.T) {
	srv := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/testx/design/v2/namespaces/ns1/labels" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Query().Get("design_uid") != "d1" {
			t.Errorf("unexpected design_uid param: %s", r.URL.Query().Get("design_uid"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Error":null,"Data":[{"name":"label1","value":"val1"}],"TotalCount":1}`))
	})
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, total, err := c.TestxListDesignLabels(context.Background(), &model.TestxListDesignLabelsRequest{
		Namespace: "ns1",
		DesignUid: "d1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Errorf("got TotalCount=%d, want 1", total)
	}
	if len(result) != 1 {
		t.Errorf("got %d labels, want 1", len(result))
	}
	if result[0].Name != "label1" {
		t.Errorf("got Name=%q, want label1", result[0].Name)
	}
}
