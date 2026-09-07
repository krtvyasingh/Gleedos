package tlscloak

import "testing"

func TestGetModernCipherSuites(t *testing.T) {
	ciphers := GetModernCipherSuites()
	if len(ciphers) != 3 {
		t.Errorf("expected 3 modern ciphers")
	}
}
