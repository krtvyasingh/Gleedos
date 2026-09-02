package torrentv2

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputeMerkleRoot(chunkHashes [][]byte) string {
	if len(chunkHashes) == 0 {
		return ""
	}
	h := sha256.New()
	for _, ch := range chunkHashes {
		h.Write(ch)
	}
	return hex.EncodeToString(h.Sum(nil))
}
