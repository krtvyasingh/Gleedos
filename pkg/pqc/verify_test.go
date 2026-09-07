package pqc

import "testing"

func TestVerifyCiphertextLength(t *testing.T) {
	if !VerifyCiphertextLength(make([]byte, 1568)) {
		t.Errorf("expected true for 1568 bytes")
	}
	if VerifyCiphertextLength(make([]byte, 100)) {
		t.Errorf("expected false for 100 bytes")
	}
}
