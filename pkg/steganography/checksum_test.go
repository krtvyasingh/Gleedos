package steganography

import "testing"

func TestComputePayloadChecksum(t *testing.T) {
	sum := ComputePayloadChecksum([]byte("secret"))
	if len(sum) != 64 {
		t.Errorf("invalid checksum: %s", sum)
	}
}
