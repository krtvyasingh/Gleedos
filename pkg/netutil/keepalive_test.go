package netutil

import (
	"testing"
	"time"
)

func TestDefaultKeepAliveConfig(t *testing.T) {
	cfg := DefaultKeepAliveConfig()
	if cfg.IdleInterval != 15*time.Second || cfg.ProbeCount != 5 {
		t.Errorf("unexpected keepalive config: %+v", cfg)
	}
}
