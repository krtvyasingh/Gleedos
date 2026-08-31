package netutil

import "testing"

func TestDefaultQUICConfig(t *testing.T) {
	cfg := DefaultQUICConfig()
	if !cfg.Enabled || cfg.MaxIdleMs != 30000 {
		t.Errorf("unexpected QUIC config: %+v", cfg)
	}
}
