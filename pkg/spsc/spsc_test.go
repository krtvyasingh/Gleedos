package spsc

import "testing"

func TestSPSCRingBuffer(t *testing.T) {
	rb := NewRingBuffer(4) // 16 elements
	if !rb.Push(0x42) {
		t.Fatalf("push failed")
	}
	b, ok := rb.Pop()
	if !ok || b != 0x42 {
		t.Fatalf("pop mismatch: %v, %v", b, ok)
	}
}
