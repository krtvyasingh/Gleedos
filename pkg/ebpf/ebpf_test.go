package ebpf

import "testing"

func TestXDPFilter(t *testing.T) {
	f := NewXDPFilter(443)
	if f.ProcessPacket(443) != XDPPass {
		t.Errorf("expected pass for port 443")
	}
}
