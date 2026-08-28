package tagger

import (
	"bytes"
	"os"
)

func IsFLAC(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	hdr := make([]byte, 4)
	n, _ := f.Read(hdr)
	return n == 4 && bytes.Equal(hdr, []byte("fLaC"))
}
