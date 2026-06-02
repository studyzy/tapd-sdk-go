package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestStorageSave(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open_app_storage/save" {
			t.Errorf("unexpected path: %s, want /open_app_storage/save", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		body, _ := io.ReadAll(r.Body)
		var m map[string]interface{}
		if err := json.Unmarshal(body, &m); err != nil {
			t.Fatalf("body is not valid JSON: %v", err)
		}
		if m["document"] != "test_doc" {
			t.Errorf("document = %v, want test_doc", m["document"])
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.StorageSave(context.Background(), &model.StorageSaveRequest{
		Document: "test_doc",
		Data:     map[string]interface{}{"field1": "value1", "field2": "value2"},
	})
	if err != nil {
		t.Fatalf("StorageSave() unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStorageQuery(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open_app_storage/query" {
			t.Errorf("unexpected path: %s, want /open_app_storage/query", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		body, _ := io.ReadAll(r.Body)
		var m map[string]interface{}
		if err := json.Unmarshal(body, &m); err != nil {
			t.Fatalf("body is not valid JSON: %v", err)
		}
		if m["document"] != "test_doc" {
			t.Errorf("document = %v, want test_doc", m["document"])
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[{"id":"1","context_data":"{\"field1\": \"value1\"}","field1":"value1"},{"id":"2","context_data":"{\"field1\": \"value1\"}","field1":"value1"}],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.StorageQuery(context.Background(), &model.StorageQueryRequest{
		Document: "test_doc",
		Limit:    10,
		Offset:   0,
	})
	if err != nil {
		t.Fatalf("StorageQuery() unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	// 验证返回数据包含预期的 JSON 数组
	var items []map[string]interface{}
	if err := json.Unmarshal(result, &items); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0]["id"] != "1" {
		t.Errorf("first item id = %v, want 1", items[0]["id"])
	}
}

func TestStorageUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open_app_storage/update" {
			t.Errorf("unexpected path: %s, want /open_app_storage/update", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		body, _ := io.ReadAll(r.Body)
		var m map[string]interface{}
		if err := json.Unmarshal(body, &m); err != nil {
			t.Fatalf("body is not valid JSON: %v", err)
		}
		if m["document"] != "test_doc" {
			t.Errorf("document = %v, want test_doc", m["document"])
		}
		if m["condition"] == nil {
			t.Error("expected condition in body")
		}
		if m["data"] == nil {
			t.Error("expected data in body")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.StorageUpdate(context.Background(), &model.StorageUpdateRequest{
		Document:  "test_doc",
		Condition: map[string]interface{}{"id": "1"},
		Data:      map[string]interface{}{"field1": "updated_value"},
	})
	if err != nil {
		t.Fatalf("StorageUpdate() unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStorageDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open_app_storage/delete" {
			t.Errorf("unexpected path: %s, want /open_app_storage/delete", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		body, _ := io.ReadAll(r.Body)
		var m map[string]interface{}
		if err := json.Unmarshal(body, &m); err != nil {
			t.Fatalf("body is not valid JSON: %v", err)
		}
		if m["document"] != "test_doc" {
			t.Errorf("document = %v, want test_doc", m["document"])
		}
		if m["condition"] == nil {
			t.Error("expected condition in body")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":1,"data":[],"info":"success"}`)
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	result, err := c.StorageDelete(context.Background(), &model.StorageDeleteRequest{
		Document:  "test_doc",
		Condition: map[string]interface{}{"id": "1"},
	})
	if err != nil {
		t.Fatalf("StorageDelete() unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
