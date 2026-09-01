package vault

import (
	"bytes"
	"testing"
)

func TestVaultEncryption(t *testing.T) {
	key := DeriveVaultKey("my-secret-vault-password")
	original := []byte("confidential media video payload")

	encrypted, err := EncryptPayload(original, key)
	if err != nil {
		t.Fatalf("EncryptPayload failed: %v", err)
	}

	decrypted, err := DecryptPayload(encrypted, key)
	if err != nil {
		t.Fatalf("DecryptPayload failed: %v", err)
	}

	if !bytes.Equal(original, decrypted) {
		t.Errorf("decrypted payload mismatch")
	}
}
