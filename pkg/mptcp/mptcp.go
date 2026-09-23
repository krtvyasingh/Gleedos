package mptcp

type MPTCPConfig struct {
	Enabled bool
}

func NewMPTCPConfig(enable bool) *MPTCPConfig {
	return &MPTCPConfig{Enabled: enable}
}

func (m *MPTCPConfig) GetSocketProtocol() int {
	if m.Enabled {
		return 262 // IPPROTO_MPTCP
	}
	return 6 // IPPROTO_TCP
}
