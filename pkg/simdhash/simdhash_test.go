package simdhash

import "testing"

func TestChecksumFast(t *testing.T) {
	sum := ChecksumFast([]byte("simd_stream"))
	if sum == 0 {
		t.Errorf("invalid checksum")
	}
}
