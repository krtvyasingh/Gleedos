package netutil

import "time"

type KeepAliveConfig struct {
	IdleInterval time.Duration
	ProbeCount   int
}

func DefaultKeepAliveConfig() KeepAliveConfig {
	return KeepAliveConfig{IdleInterval: 15 * time.Second, ProbeCount: 5}
}
