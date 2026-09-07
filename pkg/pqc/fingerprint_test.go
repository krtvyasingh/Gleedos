package pqc

import "testing"

func TestComputeKeyFingerprint(t *testing.T) {
	fp := ComputeKeyFingerprint([]byte("kyber_public_key_bytes"))
	if len(fp) != 32 {
		t.Errorf("invalid fingerprint length: %s", fp)
	}
}
