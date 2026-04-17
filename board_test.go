package tapd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/studyzy/tapd-sdk-go/model"
)

func TestCreateBoardCard(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/board_cards" {
			t.Errorf("unexpected path: %s, want /board_cards", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"BoardCard":{"id":"1001","workspace_id":"11111111","name":"新看板项","board_id":"2001","column_id":"3001","description":"test","owner":"admin","priority":"medium","created":"2026-01-01 10:00:00"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.CreateBoardCardRequest{
		WorkspaceID: "11111111",
		Name:        "新看板项",
		BoardID:     "2001",
		ColumnID:    "3001",
		Description: "test",
		Owner:       "admin",
		Priority:    "medium",
	}
	card, err := c.CreateBoardCard(req)
	if err != nil {
		t.Fatalf("CreateBoardCard() unexpected error: %v", err)
	}
	if card.ID != "1001" {
		t.Errorf("ID = %q, want %q", card.ID, "1001")
	}
	if card.Name != "新看板项" {
		t.Errorf("Name = %q, want %q", card.Name, "新看板项")
	}
	if card.BoardID != "2001" {
		t.Errorf("BoardID = %q, want %q", card.BoardID, "2001")
	}
	if card.Priority != "medium" {
		t.Errorf("Priority = %q, want %q", card.Priority, "medium")
	}
}

func TestGetBoardCards(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/board_cards" {
			t.Errorf("unexpected path: %s, want /board_cards", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"BoardCard":{"id":"1001","workspace_id":"11111111","name":"看板项1","board_id":"2001","column_id":"3001","owner":"admin","created":"2026-01-01 10:00:00"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetBoardCardsRequest{
		WorkspaceID: "11111111",
		BoardID:     "2001",
	}
	cards, err := c.GetBoardCards(req)
	if err != nil {
		t.Fatalf("GetBoardCards() unexpected error: %v", err)
	}
	if len(cards) != 1 {
		t.Fatalf("expected 1 card, got %d", len(cards))
	}
	if cards[0].ID != "1001" {
		t.Errorf("card id = %q, want %q", cards[0].ID, "1001")
	}
	if cards[0].Name != "看板项1" {
		t.Errorf("card name = %q, want %q", cards[0].Name, "看板项1")
	}
	if cards[0].Owner != "admin" {
		t.Errorf("card owner = %q, want %q", cards[0].Owner, "admin")
	}
}

func TestUpdateBoardCard(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/board_cards" {
			t.Errorf("unexpected path: %s, want /board_cards", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":{"BoardCard":{"id":"1001","workspace_id":"11111111","name":"已更新","column_id":"3002"}},"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.UpdateBoardCardRequest{
		WorkspaceID: "11111111",
		ID:          "1001",
		Name:        "已更新",
		ColumnID:    "3002",
	}
	card, err := c.UpdateBoardCard(req)
	if err != nil {
		t.Fatalf("UpdateBoardCard() unexpected error: %v", err)
	}
	if card.ID != "1001" {
		t.Errorf("ID = %q, want %q", card.ID, "1001")
	}
	if card.Name != "已更新" {
		t.Errorf("Name = %q, want %q", card.Name, "已更新")
	}
	if card.ColumnID != "3002" {
		t.Errorf("ColumnID = %q, want %q", card.ColumnID, "3002")
	}
}

func TestGetBoardColumns(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/board_columns" {
			t.Errorf("unexpected path: %s, want /board_columns", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":1,"data":[{"BoardColumn":{"id":"3001","workspace_id":"11111111","board_id":"2001","name":"待处理","sort":"1"}},{"BoardColumn":{"id":"3002","workspace_id":"11111111","board_id":"2001","name":"进行中","sort":"2"}}],"info":"success"}`))
	}))
	defer srv.Close()

	c := NewClientWithBaseURL(srv.URL, "", "test-token", "", "")
	req := &model.GetBoardColumnsRequest{
		WorkspaceID: "11111111",
		BoardID:     "2001",
	}
	cols, err := c.GetBoardColumns(req)
	if err != nil {
		t.Fatalf("GetBoardColumns() unexpected error: %v", err)
	}
	if len(cols) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(cols))
	}
	if cols[0].Name != "待处理" {
		t.Errorf("first column name = %q, want %q", cols[0].Name, "待处理")
	}
	if cols[1].Name != "进行中" {
		t.Errorf("second column name = %q, want %q", cols[1].Name, "进行中")
	}
	if cols[0].Sort != "1" {
		t.Errorf("first column sort = %q, want %q", cols[0].Sort, "1")
	}
}
