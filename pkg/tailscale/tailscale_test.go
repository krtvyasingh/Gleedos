package tailscale

import "testing"

func TestFormatNodeAddr(t *testing.T) {
	addr := FormatNodeAddr(MeshNode{TailnetIP: "100.64.0.1"}, 8080)
	if addr != "100.64.0.1" {
		t.Errorf("unexpected tailnet address")
	}
}
