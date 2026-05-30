package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestGetWorkflowTransitions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workflows/all_transitions" {
			t.Errorf("unexpected path: %s, want /workflows/all_transitions", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Name":"开始处理","StepPrevious":"open","StepNext":"progressing","Inform":[{"InformType":"owner","InformId":""}],"Appendfield":[{"DBModel":"Story","FieldName":"owner","Notnull":"1","Sort":"0","DefaultValue":[]}]}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkflowRequest{
		WorkspaceID: "1",
		System:      "story",
	}
	transitions, err := c.GetWorkflowTransitions(context.Background(), req)
	if err != nil {
		t.Fatalf("GetWorkflowTransitions() unexpected error: %v", err)
	}
	if len(transitions) != 1 {
		t.Fatalf("expected 1 transition, got %d", len(transitions))
	}
	if transitions[0].Name != "开始处理" {
		t.Errorf("transition name = %q, want %q", transitions[0].Name, "开始处理")
	}
	if transitions[0].StepPrevious != "open" {
		t.Errorf("step previous = %q, want %q", transitions[0].StepPrevious, "open")
	}
	if transitions[0].StepNext != "progressing" {
		t.Errorf("step next = %q, want %q", transitions[0].StepNext, "progressing")
	}
}

func TestGetWorkflowStatusMap(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workflows/status_map" {
			t.Errorf("unexpected path: %s, want /workflows/status_map", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"open":"未开始","progressing":"进行中","done":"已完成"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkflowRequest{
		WorkspaceID: "1",
		System:      "story",
	}
	statusMap, err := c.GetWorkflowStatusMap(context.Background(), req)
	if err != nil {
		t.Fatalf("GetWorkflowStatusMap() unexpected error: %v", err)
	}
	if len(statusMap) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(statusMap))
	}
	if statusMap["open"] != "未开始" {
		t.Errorf("open = %q, want %q", statusMap["open"], "未开始")
	}
	if statusMap["progressing"] != "进行中" {
		t.Errorf("progressing = %q, want %q", statusMap["progressing"], "进行中")
	}
	if statusMap["done"] != "已完成" {
		t.Errorf("done = %q, want %q", statusMap["done"], "已完成")
	}
}

func TestGetWorkflowLastSteps(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workflows/last_steps" {
			t.Errorf("unexpected path: %s, want /workflows/last_steps", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"done":"已完成","rejected":"已拒绝"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkflowRequest{
		WorkspaceID: "1",
		System:      "story",
	}
	lastSteps, err := c.GetWorkflowLastSteps(context.Background(), req)
	if err != nil {
		t.Fatalf("GetWorkflowLastSteps() unexpected error: %v", err)
	}
	if len(lastSteps) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(lastSteps))
	}
	if lastSteps["done"] != "已完成" {
		t.Errorf("done = %q, want %q", lastSteps["done"], "已完成")
	}
	if lastSteps["rejected"] != "已拒绝" {
		t.Errorf("rejected = %q, want %q", lastSteps["rejected"], "已拒绝")
	}
}

func TestGetWorkflowAllLastSteps(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workflows/all_last_steps" {
			t.Errorf("unexpected path: %s, want /workflows/all_last_steps", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"1069990230000187079":{"closed":"已关闭"},"1069990230000131609":{"rejected":"已拒绝"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkflowRequest{
		WorkspaceID: "1",
		System:      "bug",
	}
	allLastSteps, err := c.GetWorkflowAllLastSteps(context.Background(), req)
	if err != nil {
		t.Fatalf("GetWorkflowAllLastSteps() unexpected error: %v", err)
	}
	if len(allLastSteps) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(allLastSteps))
	}
	if allLastSteps["1069990230000187079"]["closed"] != "已关闭" {
		t.Errorf("closed = %q, want %q", allLastSteps["1069990230000187079"]["closed"], "已关闭")
	}
	if allLastSteps["1069990230000131609"]["rejected"] != "已拒绝" {
		t.Errorf("rejected = %q, want %q", allLastSteps["1069990230000131609"]["rejected"], "已拒绝")
	}
}

func TestGetWorkflowFirstStep(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workflows/first_step" {
			t.Errorf("unexpected path: %s, want /workflows/first_step", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"planning":"规划中"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkflowRequest{
		WorkspaceID: "1",
		System:      "story",
	}
	firstStep, err := c.GetWorkflowFirstStep(context.Background(), req)
	if err != nil {
		t.Fatalf("GetWorkflowFirstStep() unexpected error: %v", err)
	}
	if firstStep["planning"] != "规划中" {
		t.Errorf("first step planning = %q, want %q", firstStep["planning"], "规划中")
	}
}

func TestGetWorkflows(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/workflows" {
			t.Errorf("unexpected path: %s, want /workflows", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Workflow":{"id":"1","name":"默认工作流","system":"story"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkflowRequest{
		WorkspaceID: "1",
		System:      "story",
	}
	data, err := c.GetWorkflows(context.Background(), req)
	if err != nil {
		t.Fatalf("GetWorkflows() unexpected error: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty data")
	}
	if data[0].ID != "1" {
		t.Errorf("expected id=1, got %q", data[0].ID)
	}
	if data[0].Name != "默认工作流" {
		t.Errorf("expected name=默认工作流, got %q", data[0].Name)
	}
}
