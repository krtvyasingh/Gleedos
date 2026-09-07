package vault

import "testing"

func TestDeriveKeyWithSalt(t *testing.T) {
	k1 := DeriveKeyWithSalt("password", "salt1", 100)
	k2 := DeriveKeyWithSalt("password", "salt2", 100)
	if k1 == k2 || len(k1) != 64 {
		t.Errorf("invalid key derivation: %s, %s", k1, k2)
	}
}
