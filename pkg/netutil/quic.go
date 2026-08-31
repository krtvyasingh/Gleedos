package netutil

type QUICConfig struct {
	Enabled   bool
	MaxIdleMs int
}

func DefaultQUICConfig() QUICConfig {
	return QUICConfig{Enabled: true, MaxIdleMs: 30000}
}
