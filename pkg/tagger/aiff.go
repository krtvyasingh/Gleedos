package tagger

import "bytes"

func IsAIFF(hdr []byte) bool {
	return len(hdr) >= 12 && bytes.Equal(hdr[:4], []byte("FORM")) && bytes.Equal(hdr[8:12], []byte("AIFF"))
}
