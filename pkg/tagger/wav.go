package tagger

import (
	"bytes"
	"os"
)

func IsWAV(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	hdr := make([]byte, 12)
	n, _ := f.Read(hdr)
	return n == 12 && bytes.Equal(hdr[:4], []byte("RIFF")) && bytes.Equal(hdr[8:12], []byte("WAVE"))
}
