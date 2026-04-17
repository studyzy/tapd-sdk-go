package tapd

import (
	"encoding/json"
	"testing"
)

type testItem struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func TestParseList(t *testing.T) {
	data := json.RawMessage(`[{"Item":{"id":"1","name":"A"}},{"Item":{"id":"2","name":"B"}}]`)
	items, err := parseList[testItem](data, "Item")
	if err != nil {
		t.Fatalf("parseList() unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].ID != "1" || items[0].Name != "A" {
		t.Errorf("first item = %+v, want {ID:1 Name:A}", items[0])
	}
	if items[1].ID != "2" || items[1].Name != "B" {
		t.Errorf("second item = %+v, want {ID:2 Name:B}", items[1])
	}
}

func TestParseList_Empty(t *testing.T) {
	data := json.RawMessage(`[]`)
	items, err := parseList[testItem](data, "Item")
	if err != nil {
		t.Fatalf("parseList() unexpected error: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 items, got %d", len(items))
	}
}

func TestParseList_WrongKey(t *testing.T) {
	data := json.RawMessage(`[{"Other":{"id":"1"}}]`)
	items, err := parseList[testItem](data, "Item")
	if err != nil {
		t.Fatalf("parseList() unexpected error: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 items for wrong key, got %d", len(items))
	}
}

func TestParseList_InvalidJSON(t *testing.T) {
	data := json.RawMessage(`not json`)
	_, err := parseList[testItem](data, "Item")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestParseList_ItemUnmarshalError(t *testing.T) {
	// key 匹配但值不是合法的 testItem JSON，应返回错误而非静默跳过
	data := json.RawMessage(`[{"Item":{"id":"1","name":"A"}},{"Item":"bad"}]`)
	_, err := parseList[testItem](data, "Item")
	if err == nil {
		t.Fatal("expected error for invalid item, got nil")
	}
}

func TestParseList_MixedKeys(t *testing.T) {
	// 不含目标 key 的条目应被正常跳过
	data := json.RawMessage(`[{"Item":{"id":"1","name":"A"}},{"Other":{"id":"2"}}]`)
	items, err := parseList[testItem](data, "Item")
	if err != nil {
		t.Fatalf("parseList() unexpected error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].ID != "1" {
		t.Errorf("item.ID = %q, want %q", items[0].ID, "1")
	}
}

func TestParseOne(t *testing.T) {
	data := json.RawMessage(`{"Item":{"id":"1","name":"A"}}`)
	item, err := parseOne[testItem](data, "Item")
	if err != nil {
		t.Fatalf("parseOne() unexpected error: %v", err)
	}
	if item.ID != "1" || item.Name != "A" {
		t.Errorf("item = %+v, want {ID:1 Name:A}", item)
	}
}

func TestParseOne_MissingKey(t *testing.T) {
	data := json.RawMessage(`{"Other":{"id":"1"}}`)
	_, err := parseOne[testItem](data, "Item")
	if err == nil {
		t.Fatal("expected error for missing key")
	}
}

func TestParseOne_InvalidJSON(t *testing.T) {
	data := json.RawMessage(`not json`)
	_, err := parseOne[testItem](data, "Item")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestParseCount(t *testing.T) {
	data := json.RawMessage(`{"count":42}`)
	count, err := parseCount(data)
	if err != nil {
		t.Fatalf("parseCount() unexpected error: %v", err)
	}
	if count != 42 {
		t.Errorf("count = %d, want 42", count)
	}
}

func TestParseCount_Zero(t *testing.T) {
	data := json.RawMessage(`{"count":0}`)
	count, err := parseCount(data)
	if err != nil {
		t.Fatalf("parseCount() unexpected error: %v", err)
	}
	if count != 0 {
		t.Errorf("count = %d, want 0", count)
	}
}

func TestParseCount_MissingKey(t *testing.T) {
	data := json.RawMessage(`{}`)
	_, err := parseCount(data)
	if err == nil {
		t.Fatalf("parseCount() expected error, got nil")
	}
}
