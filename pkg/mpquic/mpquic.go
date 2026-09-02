package mpquic

import "sync"

type Subflow struct {
	InterfaceName string
	BandwidthBps  int64
	RTTMs         int
}

type MultipathScheduler struct {
	subflows []*Subflow
	mu       sync.RWMutex
}

func NewMultipathScheduler() *MultipathScheduler {
	return &MultipathScheduler{}
}

func (m *MultipathScheduler) AddSubflow(iface string, bw int64, rtt int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.subflows = append(m.subflows, &Subflow{InterfaceName: iface, BandwidthBps: bw, RTTMs: rtt})
}

func (m *MultipathScheduler) GetTotalBandwidth() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var total int64
	for _, s := range m.subflows {
		total += s.BandwidthBps
	}
	return total
}
