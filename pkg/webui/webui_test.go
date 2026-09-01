package webui

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	Handler().ServeHTTP(rec, req)
	if !strings.Contains(rec.Body.String(), "Gleedos Media Engine") {
		t.Errorf("unexpected HTML output")
	}
}
