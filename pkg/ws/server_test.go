package ws

import (
	"net/http/httptest"
	"testing"
)

func TestHubBroadcast(t *testing.T) {
	hub := NewHub()
	hub.Broadcast(map[string]string{"status": "downloading"})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ws", nil)
	hub.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}
