package nostr

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ComputeEventID(pubKey, content string, createdAt int64, kind int) string {
	serialized := fmt.Sprintf("[0,\"%s\",%d,%d,[],\"%s\"]", pubKey, createdAt, kind, content)
	h := sha256.Sum256([]byte(serialized))
	return hex.EncodeToString(h[:])
}
