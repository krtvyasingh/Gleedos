package bitset

import "testing"

func TestChunkBitset(t *testing.T) {
	bs := NewChunkBitset(100)
	bs.Set(5)
	if !bs.IsSet(5) {
		t.Errorf("expected bit 5 to be set")
	}
	if bs.IsSet(6) {
		t.Errorf("expected bit 6 to be unset")
	}
}
