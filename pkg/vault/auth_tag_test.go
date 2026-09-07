package vault

import "testing"

func TestAuthTag(t *testing.T) {
	data := []byte("payload")
	key := []byte("secret_key_123456789012345678901")
	tag := GenerateAuthTag(data, key)
	if !VerifyAuthTag(data, key, tag) {
		t.Errorf("tag verification failed")
	}
}
