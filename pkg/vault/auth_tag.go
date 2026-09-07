package vault

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GenerateAuthTag(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func VerifyAuthTag(data, key, expectedTag []byte) bool {
	tag := GenerateAuthTag(data, key)
	return hmac.Equal(tag, expectedTag)
}
