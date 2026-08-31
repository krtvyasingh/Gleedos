package security

import "testing"

func TestSignHMAC(t *testing.T) {
	sig := SignHMAC([]byte("data"), []byte("secret"))
	if len(sig) != 64 {
		t.Errorf("invalid signature length: %s", sig)
	}
}
