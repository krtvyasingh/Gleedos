package steganography

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputePayloadChecksum(payload []byte) string {
	h := sha256.Sum256(payload)
	return hex.EncodeToString(h[:])
}
