package vvc

import "bytes"

func IsAV1(header []byte) bool {
	return len(header) >= 4 && bytes.Contains(header, []byte("av01"))
}

func IsVVC(header []byte) bool {
	return len(header) >= 4 && bytes.Contains(header, []byte("vvc1"))
}
