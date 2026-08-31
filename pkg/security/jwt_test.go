package security

import "testing"

func TestGenerateSecureToken(t *testing.T) {
	tok, err := GenerateSecureToken(16)
	if err != nil || len(tok) != 32 {
		t.Fatalf("unexpected token: %s, %v", tok, err)
	}
}
