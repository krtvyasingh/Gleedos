package storage

import "sync"

type MemoryKV struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryKV() *MemoryKV {
	return &MemoryKV{data: make(map[string]string)}
}

func (m *MemoryKV) Set(k, v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}

func (m *MemoryKV) Get(k string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[k]
	return v, ok
}
