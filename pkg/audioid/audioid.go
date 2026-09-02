package audioid

import (
	"crypto/sha256"
	"encoding/hex"
)

func FingerprintSamples(samples []float64) string {
	h := sha256.New()
	for _, s := range samples {
		h.Write([]byte{byte(s * 127)})
	}
	return hex.EncodeToString(h.Sum(nil))
}
