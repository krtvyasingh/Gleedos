package vault

import (
	"crypto/sha256"
	"encoding/hex"
)

func DeriveKeyWithSalt(passphrase, salt string, iterations int) string {
	h := sha256.New()
	h.Write([]byte(passphrase + salt))
	res := h.Sum(nil)
	for i := 1; i < iterations; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
