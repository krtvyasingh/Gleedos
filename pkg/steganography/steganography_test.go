package steganography

import (
	"bytes"
	"testing"
)

func TestSteganography(t *testing.T) {
	original := []byte("ftypmp42...media_stream...")
	secret := []byte("top_secret_auth_token")

	chaffed := HidePayload(original, secret)
	extracted, err := ExtractPayload(chaffed)
	if err != nil || !bytes.Equal(extracted, secret) {
		t.Fatalf("ExtractPayload failed: %v, got %s", err, string(extracted))
	}
}
