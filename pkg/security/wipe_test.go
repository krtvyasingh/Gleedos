package security

import "testing"

func TestSecureWipe(t *testing.T) {
	buf := []byte("sensitive_auth_key")
	SecureWipe(buf)
	for _, b := range buf {
		if b != 0 {
			t.Fatalf("buffer was not zeroed")
		}
	}
}
