package spki

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputeSPKIFingerprint(pubKeyASN1 []byte) string {
	h := sha256.Sum256(pubKeyASN1)
	return hex.EncodeToString(h[:])
}
