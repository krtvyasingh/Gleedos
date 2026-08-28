package hls

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

func TestDecryptAES128(t *testing.T) {
	key := []byte("1234567890123456")
	iv := []byte("abcdefghijklmnop")
	plain := []byte("16_byte_payload_")

	block, _ := aes.NewCipher(key)
	enc := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plain))
	enc.CryptBlocks(cipherText, plain)

	dec, err := DecryptAES128(cipherText, key, iv)
	if err != nil || string(dec) != string(plain) {
		t.Fatalf("DecryptAES128 failed: %v, got %s", err, string(dec))
	}
}
