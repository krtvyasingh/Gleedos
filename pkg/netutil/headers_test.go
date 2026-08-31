package netutil

import (
	"net/http"
	"testing"
)

func TestApplyStandardHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	ApplyStandardHeaders(req)
	if req.Header.Get("Accept") != "*/*" {
		t.Errorf("missing Accept header")
	}
}
