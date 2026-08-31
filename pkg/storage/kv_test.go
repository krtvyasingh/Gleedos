package storage

import "testing"

func TestMemoryKV(t *testing.T) {
	kv := NewMemoryKV()
	kv.Set("session_1", "active")
	val, ok := kv.Get("session_1")
	if !ok || val != "active" {
		t.Errorf("unexpected value: %s", val)
	}
}
