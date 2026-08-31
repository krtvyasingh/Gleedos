package storage

import "sync"

type LRUCache struct {
	capacity int
	items    map[string]string
	mu       sync.Mutex
}

func NewLRUCache(cap int) *LRUCache {
	if cap <= 0 {
		cap = 100
	}
	return &LRUCache{capacity: cap, items: make(map[string]string)}
}

func (c *LRUCache) Put(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.items) >= c.capacity {
		for key := range c.items {
			delete(c.items, key)
			break
		}
	}
	c.items[k] = v
}
