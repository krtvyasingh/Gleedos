package pqc

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputeKeyFingerprint(pubKey []byte) string {
	h := sha256.Sum256(pubKey)
	return hex.EncodeToString(h[:16])
}
