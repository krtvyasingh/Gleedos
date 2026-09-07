package doh

import "sync"

type ResolverPool struct {
	endpoints []string
	idx       int
	mu        sync.Mutex
}

func NewResolverPool(endpoints []string) *ResolverPool {
	if len(endpoints) == 0 {
		endpoints = []string{"1.1.1.1:853", "8.8.8.8:853", "9.9.9.9:853"}
	}
	return &ResolverPool{endpoints: endpoints}
}

func (r *ResolverPool) NextEndpoint() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	ep := r.endpoints[r.idx%len(r.endpoints)]
	r.idx++
	return ep
}
