package torrentv2

import "testing"

func TestComputeMerkleRoot(t *testing.T) {
	h1 := []byte("hash1")
	h2 := []byte("hash2")
	root := ComputeMerkleRoot([][]byte{h1, h2})
	if len(root) != 64 {
		t.Errorf("invalid merkle root: %s", root)
	}
}
