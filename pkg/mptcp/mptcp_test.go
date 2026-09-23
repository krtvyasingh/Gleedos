package mptcp

import "testing"

func TestMPTCPConfig(t *testing.T) {
	cfg := NewMPTCPConfig(true)
	if cfg.GetSocketProtocol() != 262 {
		t.Errorf("expected MPTCP protocol 262")
	}
}
