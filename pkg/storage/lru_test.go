package storage

import "testing"

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put("a", "1")
	lru.Put("b", "2")
	lru.Put("c", "3")
	if len(lru.items) > 2 {
		t.Errorf("capacity exceeded")
	}
}
