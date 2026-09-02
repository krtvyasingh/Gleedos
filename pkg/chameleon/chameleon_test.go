package chameleon

import (
	"net/http"
	"testing"
)

func TestApplyMobileWireProfile(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	ApplyMobileWireProfile(req, "youtube_ios")
	if req.Header.Get("X-YouTube-Client-Name") != "5" {
		t.Errorf("missing mobile headers")
	}
}
