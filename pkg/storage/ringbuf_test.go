package storage

import "testing"

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(5)
	n := rb.Write([]byte("12345"))
	if n != 5 {
		t.Errorf("expected 5 bytes written")
	}
}
