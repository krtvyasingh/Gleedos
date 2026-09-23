package shm

import "testing"

func TestSHMBuffer(t *testing.T) {
	shm := NewSHMBuffer("shm_media_1", 1024*1024)
	if shm.Size != 1024*1024 {
		t.Errorf("unexpected size: %d", shm.Size)
	}
}
