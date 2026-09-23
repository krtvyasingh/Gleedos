package iouring

import "testing"

func TestRingSimulator(t *testing.T) {
	r := NewRingSimulator(16)
	cqe := r.Submit(SQE{Opcode: 1, Fd: 3, Buf: []byte("chunk")})
	if cqe.Res != 5 {
		t.Errorf("expected 5 bytes, got %d", cqe.Res)
	}
}
