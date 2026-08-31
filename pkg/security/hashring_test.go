package security

import "testing"

func TestComputeChunkHash(t *testing.T) {
	h := ComputeChunkHash([]byte("chunk_data"))
	if len(h) != 64 {
		t.Errorf("invalid hash length: %s", h)
	}
}
