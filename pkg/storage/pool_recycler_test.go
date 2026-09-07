package storage

import "testing"

func TestBufferPool(t *testing.T) {
	bp := NewBufferPool(4096)
	buf := bp.Get()
	if len(buf) != 4096 {
		t.Errorf("unexpected buffer size: %d", len(buf))
	}
	bp.Put(buf)
}
