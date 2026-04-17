package tapd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestListMiniItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items" {
			t.Errorf("unexpected path: %s, want /mini_items", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"MiniItem":{"id":"100","name":"Test Item","status":"open","owner":"user1","creator":"admin","workspace_id":"1"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListMiniItemsRequest{
		WorkspaceID: "1",
	}
	items, err := c.ListMiniItems(context.Background(), req)
	if err != nil {
		t.Fatalf("ListMiniItems() unexpected error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].ID != "100" {
		t.Errorf("id = %q, want %q", items[0].ID, "100")
	}
	if items[0].Name != "Test Item" {
		t.Errorf("name = %q, want %q", items[0].Name, "Test Item")
	}
	if items[0].Status != "open" {
		t.Errorf("status = %q, want %q", items[0].Status, "open")
	}
	if items[0].Owner != "user1" {
		t.Errorf("owner = %q, want %q", items[0].Owner, "user1")
	}
	if items[0].Creator != "admin" {
		t.Errorf("creator = %q, want %q", items[0].Creator, "admin")
	}
}

func TestGetMiniItem(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items" {
			t.Errorf("unexpected path: %s, want /mini_items", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"MiniItem":{"id":"100","name":"Test Item","description":"<p>Hello</p>","creator":"admin","created":"2026-01-01 10:00:00"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	item, err := c.GetMiniItem(context.Background(), "1", "100")
	if err != nil {
		t.Fatalf("GetMiniItem() unexpected error: %v", err)
	}
	if item.ID != "100" {
		t.Errorf("id = %q, want %q", item.ID, "100")
	}
	if item.Description != "<p>Hello</p>" {
		t.Errorf("description = %q, want %q", item.Description, "<p>Hello</p>")
	}
	if item.Creator != "admin" {
		t.Errorf("creator = %q, want %q", item.Creator, "admin")
	}
	if item.Created != "2026-01-01 10:00:00" {
		t.Errorf("created = %q, want %q", item.Created, "2026-01-01 10:00:00")
	}
}

func TestGetMiniItem_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	_, err := c.GetMiniItem(context.Background(), "1", "999")
	if err == nil {
		t.Fatal("expected error for not found, got nil")
	}
}

func TestCreateMiniItem(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/mini_items" {
			t.Errorf("unexpected path: %s, want /mini_items", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"MiniItem":{"id":"200","name":"New Item","workspace_id":"1","status":"open"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateMiniItemRequest{
		WorkspaceID: "1",
		Name:        "New Item",
	}
	item, err := c.CreateMiniItem(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateMiniItem() unexpected error: %v", err)
	}
	if item.ID != "200" {
		t.Errorf("id = %q, want %q", item.ID, "200")
	}
	if item.Name != "New Item" {
		t.Errorf("name = %q, want %q", item.Name, "New Item")
	}
}

func TestUpdateMiniItem(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/mini_items" {
			t.Errorf("unexpected path: %s, want /mini_items", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"MiniItem":{"id":"100","name":"Updated","status":"done"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.UpdateMiniItemRequest{
		WorkspaceID: "1",
		ID:          "100",
		Name:        "Updated",
		Status:      "done",
	}
	item, err := c.UpdateMiniItem(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateMiniItem() unexpected error: %v", err)
	}
	if item.ID != "100" {
		t.Errorf("id = %q, want %q", item.ID, "100")
	}
	if item.Name != "Updated" {
		t.Errorf("name = %q, want %q", item.Name, "Updated")
	}
	if item.Status != "done" {
		t.Errorf("status = %q, want %q", item.Status, "done")
	}
}

func TestCountMiniItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items/count" {
			t.Errorf("unexpected path: %s, want /mini_items/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":42},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountMiniItemsRequest{
		WorkspaceID: "1",
	}
	count, err := c.CountMiniItems(context.Background(), req)
	if err != nil {
		t.Fatalf("CountMiniItems() unexpected error: %v", err)
	}
	if count != 42 {
		t.Errorf("count = %d, want 42", count)
	}
}

func TestListMiniItems_CustomFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"MiniItem":{"id":"100","name":"Test","custom_field_one":"cf1","custom_field_9":"cf9"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	items, err := c.ListMiniItems(context.Background(), &model.ListMiniItemsRequest{WorkspaceID: "1"})
	if err != nil {
		t.Fatalf("ListMiniItems() unexpected error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].CustomFields["custom_field_one"] != "cf1" {
		t.Errorf("custom_field_one = %q, want %q", items[0].CustomFields["custom_field_one"], "cf1")
	}
	if items[0].CustomFields["custom_field_9"] != "cf9" {
		t.Errorf("custom_field_9 = %q, want %q", items[0].CustomFields["custom_field_9"], "cf9")
	}
}

func TestCreateMiniItemCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/mini_item_categories" {
			t.Errorf("unexpected path: %s, want /mini_item_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Category":{"id":"300","name":"Test Category","workspace_id":"1","creator":"admin","created":"2026-01-01 10:00:00"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateMiniItemCategoryRequest{
		WorkspaceID: "1",
		Name:        "Test Category",
	}
	cat, err := c.CreateMiniItemCategory(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateMiniItemCategory() unexpected error: %v", err)
	}
	if cat.ID != "300" {
		t.Errorf("id = %q, want %q", cat.ID, "300")
	}
	if cat.Name != "Test Category" {
		t.Errorf("name = %q, want %q", cat.Name, "Test Category")
	}
}

func TestListMiniItemCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_item_categories" {
			t.Errorf("unexpected path: %s, want /mini_item_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"Category":{"id":"300","name":"Cat1"}},{"Category":{"id":"301","name":"Cat2"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.ListMiniItemCategoriesRequest{
		WorkspaceID: "1",
	}
	cats, err := c.ListMiniItemCategories(context.Background(), req)
	if err != nil {
		t.Fatalf("ListMiniItemCategories() unexpected error: %v", err)
	}
	if len(cats) != 2 {
		t.Fatalf("expected 2 categories, got %d", len(cats))
	}
	if cats[0].ID != "300" {
		t.Errorf("first category id = %q, want %q", cats[0].ID, "300")
	}
	if cats[1].Name != "Cat2" {
		t.Errorf("second category name = %q, want %q", cats[1].Name, "Cat2")
	}
}

func TestCountMiniItemCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_item_categories/count" {
			t.Errorf("unexpected path: %s, want /mini_item_categories/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":9},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountMiniItemCategoriesRequest{
		WorkspaceID: "1",
	}
	count, err := c.CountMiniItemCategories(context.Background(), req)
	if err != nil {
		t.Fatalf("CountMiniItemCategories() unexpected error: %v", err)
	}
	if count != 9 {
		t.Errorf("count = %d, want 9", count)
	}
}

func TestGetMiniItemChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_item_changes" {
			t.Errorf("unexpected path: %s, want /mini_item_changes", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"WorkitemChange":{"id":"500","workspace_id":"1","creator":"admin","change_type":"manual_update","mini_item_id":"100"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetMiniItemChangesRequest{
		WorkspaceID: "1",
		MiniItemID:  "100",
	}
	changes, err := c.GetMiniItemChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("GetMiniItemChanges() unexpected error: %v", err)
	}
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].ID != "500" {
		t.Errorf("id = %q, want %q", changes[0].ID, "500")
	}
	if changes[0].ChangeType != "manual_update" {
		t.Errorf("change_type = %q, want %q", changes[0].ChangeType, "manual_update")
	}
	if changes[0].MiniItemID != "100" {
		t.Errorf("mini_item_id = %q, want %q", changes[0].MiniItemID, "100")
	}
}

func TestCountMiniItemChanges(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_item_changes/count" {
			t.Errorf("unexpected path: %s, want /mini_item_changes/count", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"count":21},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CountMiniItemChangesRequest{
		WorkspaceID: "1",
		MiniItemID:  "100",
	}
	count, err := c.CountMiniItemChanges(context.Background(), req)
	if err != nil {
		t.Fatalf("CountMiniItemChanges() unexpected error: %v", err)
	}
	if count != 21 {
		t.Errorf("count = %d, want 21", count)
	}
}

func TestCreateMiniItemRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/mini_items/create_mini_item_relation" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateMiniItemRelationRequest{
		WorkspaceID: "1",
		TargetType:  "story",
		SourceID:    "100",
		TargetID:    "200",
	}
	success, err := c.CreateMiniItemRelation(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateMiniItemRelation() unexpected error: %v", err)
	}
	if !success {
		t.Error("expected success = true")
	}
}

func TestRemoveMiniItemRelation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/mini_items/remove_mini_item_relation" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"success":true},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.RemoveMiniItemRelationRequest{
		WorkspaceID: "1",
		MiniItemID:  "100",
		TargetType:  "bug",
		TargetID:    "200",
	}
	success, err := c.RemoveMiniItemRelation(context.Background(), req)
	if err != nil {
		t.Fatalf("RemoveMiniItemRelation() unexpected error: %v", err)
	}
	if !success {
		t.Error("expected success = true")
	}
}

func TestGetRemovedMiniItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items/get_removed_mini_items" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"RemovedMiniItem":{"id":"100","name":"Deleted Item","creator":"admin","operation_user":"admin","deleted":"2026-01-01 12:00:00"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetRemovedMiniItemsRequest{
		WorkspaceID: "1",
	}
	items, err := c.GetRemovedMiniItems(context.Background(), req)
	if err != nil {
		t.Fatalf("GetRemovedMiniItems() unexpected error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].ID != "100" {
		t.Errorf("id = %q, want %q", items[0].ID, "100")
	}
	if items[0].Name != "Deleted Item" {
		t.Errorf("name = %q, want %q", items[0].Name, "Deleted Item")
	}
	if items[0].OperationUser != "admin" {
		t.Errorf("operation_user = %q, want %q", items[0].OperationUser, "admin")
	}
}

func TestGetMiniItemLinkedStories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items/get_link_stories" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"type":"direct_relate","id":"200","workspace_id":"1","linked_workspace_id":1}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetMiniItemLinkedStoriesRequest{
		WorkspaceID: "1",
		MiniItemID:  "100",
	}
	relations, err := c.GetMiniItemLinkedStories(context.Background(), req)
	if err != nil {
		t.Fatalf("GetMiniItemLinkedStories() unexpected error: %v", err)
	}
	if len(relations) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(relations))
	}
	if relations[0].Type != "direct_relate" {
		t.Errorf("type = %q, want %q", relations[0].Type, "direct_relate")
	}
	if relations[0].ID != "200" {
		t.Errorf("id = %q, want %q", relations[0].ID, "200")
	}
}

func TestGetMiniItemRelatedBugs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items/get_related_bugs" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"type":"direct_relate","id":"300","workspace_id":"1","linked_workspace_id":2}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetMiniItemRelatedBugsRequest{
		WorkspaceID: "1",
		MiniItemID:  "100",
	}
	relations, err := c.GetMiniItemRelatedBugs(context.Background(), req)
	if err != nil {
		t.Fatalf("GetMiniItemRelatedBugs() unexpected error: %v", err)
	}
	if len(relations) != 1 {
		t.Fatalf("expected 1 relation, got %d", len(relations))
	}
	if relations[0].ID != "300" {
		t.Errorf("id = %q, want %q", relations[0].ID, "300")
	}
}

func TestCreateMiniItem_WithCustomFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.FormValue("custom_field_one") != "value1" {
			t.Errorf("custom_field_one = %q, want %q", r.FormValue("custom_field_one"), "value1")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"MiniItem":{"id":"200","name":"New","custom_field_one":"value1"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	item, err := c.CreateMiniItem(context.Background(), &model.CreateMiniItemRequest{
		WorkspaceID:  "1",
		Name:         "New",
		CustomFields: map[string]string{"custom_field_one": "value1"},
	})
	if err != nil {
		t.Fatalf("CreateMiniItem() unexpected error: %v", err)
	}
	if item.CustomFields["custom_field_one"] != "value1" {
		t.Errorf("custom_field_one = %q, want %q", item.CustomFields["custom_field_one"], "value1")
	}
}

func TestUpdateMiniItemCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/mini_item_categories" {
			t.Errorf("unexpected path: %s, want /mini_item_categories", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"Category":{"id":"300","name":"Updated Cat","workspace_id":"1"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.UpdateMiniItemCategoryRequest{
		WorkspaceID: "1",
		ID:          "300",
		Name:        "Updated Cat",
	}
	cat, err := c.UpdateMiniItemCategory(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateMiniItemCategory() unexpected error: %v", err)
	}
	if cat.ID != "300" {
		t.Errorf("id = %q, want %q", cat.ID, "300")
	}
	if cat.Name != "Updated Cat" {
		t.Errorf("name = %q, want %q", cat.Name, "Updated Cat")
	}
}

func TestGetMiniItemCustomFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items/custom_fields_settings" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"CustomFieldConfig":{"id":"1","workspace_id":"1","custom_field":"custom_field_one","type":"file","name":"图片与文件","enabled":"1"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkspaceIDRequest{WorkspaceID: "1"}
	configs, err := c.GetMiniItemCustomFields(context.Background(), req)
	if err != nil {
		t.Fatalf("GetMiniItemCustomFields() unexpected error: %v", err)
	}
	if len(configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(configs))
	}
	if configs[0].CustomField != "custom_field_one" {
		t.Errorf("custom_field = %q, want %q", configs[0].CustomField, "custom_field_one")
	}
	if configs[0].Name != "图片与文件" {
		t.Errorf("name = %q, want %q", configs[0].Name, "图片与文件")
	}
}

func TestGetMiniItemFieldsLabel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mini_items/get_fields_label" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"标题":"name","状态":"status","处理人":"owner"},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.WorkspaceIDRequest{WorkspaceID: "1"}
	labels, err := c.GetMiniItemFieldsLabel(context.Background(), req)
	if err != nil {
		t.Fatalf("GetMiniItemFieldsLabel() unexpected error: %v", err)
	}
	if labels["标题"] != "name" {
		t.Errorf("标题 = %q, want %q", labels["标题"], "name")
	}
	if labels["状态"] != "status" {
		t.Errorf("状态 = %q, want %q", labels["状态"], "status")
	}
}
