package tailscale

import "testing"

func TestIsTailnetIP(t *testing.T) {
	if !IsTailnetIP("100.64.1.5") {
		t.Errorf("expected valid tailnet IP")
	}
	if IsTailnetIP("192.168.1.1") {
		t.Errorf("expected non-tailnet IP")
	}
}
