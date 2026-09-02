package pqc

import "testing"

func TestKyberPQC(t *testing.T) {
	kp, err := GenerateKyberKeypair()
	if err != nil || len(kp.PublicKey) != 1568 {
		t.Fatalf("GenerateKyberKeypair failed: %v", err)
	}
	ct, ss, err := Encapsulate(kp.PublicKey)
	if err != nil || len(ct) == 0 || len(ss) != 32 {
		t.Fatalf("Encapsulate failed: %v", err)
	}
}
