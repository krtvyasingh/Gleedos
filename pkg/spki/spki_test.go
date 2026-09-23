package spki

import "testing"

func TestComputeSPKIFingerprint(t *testing.T) {
	fp := ComputeSPKIFingerprint([]byte("asn1_pub_key"))
	if len(fp) != 64 {
		t.Errorf("invalid fingerprint: %s", fp)
	}
}
