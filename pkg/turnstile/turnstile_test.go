package turnstile

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHarvester(t *testing.T) {
	h := NewHarvester()
	body := `{"token":"tok_123","session_id":"sess_1"}`
	req := httptest.NewRequest("POST", "/solve", strings.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}
