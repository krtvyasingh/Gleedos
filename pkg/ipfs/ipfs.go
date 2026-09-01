package ipfs

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateMockCID(data []byte) string {
	h := sha256.Sum256(data)
	return "Qm" + hex.EncodeToString(h[:])[:44]
}
