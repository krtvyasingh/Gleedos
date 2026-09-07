package mpquic

import "sync/atomic"

type RoundRobinBalancer struct {
	counter uint64
}

func NewRoundRobinBalancer() *RoundRobinBalancer {
	return &RoundRobinBalancer{}
}

func (r *RoundRobinBalancer) NextSubflow(subflows []*Subflow) *Subflow {
	if len(subflows) == 0 {
		return nil
	}
	idx := atomic.AddUint64(&r.counter, 1) % uint64(len(subflows))
	return subflows[idx]
}
