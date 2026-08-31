package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignHMAC(message, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}
