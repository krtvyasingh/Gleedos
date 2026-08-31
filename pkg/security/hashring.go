package security

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputeChunkHash(chunk []byte) string {
	h := sha256.Sum256(chunk)
	return hex.EncodeToString(h[:])
}
