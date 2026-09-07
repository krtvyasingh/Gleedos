package torrentv2

import (
	"crypto/sha256"
	"bytes"
)

func VerifyLeafNode(leafHash []byte, parentHash []byte) bool {
	combined := append(leafHash, leafHash...)
	expected := sha256.Sum256(combined)
	return bytes.Equal(expected[:], parentHash)
}
