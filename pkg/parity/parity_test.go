package parity

import "testing"

func TestGenerateParity(t *testing.T) {
	data := []byte("0123456789ABCDEF0123456789ABCDEF")
	blocks := GenerateParity(data, 16)
	if len(blocks) != 2 {
		t.Errorf("expected 2 blocks, got %d", len(blocks))
	}
}
