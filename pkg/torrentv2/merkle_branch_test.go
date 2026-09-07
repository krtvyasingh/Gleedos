package torrentv2

import (
	"crypto/sha256"
	"testing"
)

func TestVerifyLeafNode(t *testing.T) {
	leaf := []byte("leaf_hash_data_32_bytes_long_123")
	comb := append(leaf, leaf...)
	parent := sha256.Sum256(comb)
	if !VerifyLeafNode(leaf, parent[:]) {
		t.Errorf("merkle validation failed")
	}
}
