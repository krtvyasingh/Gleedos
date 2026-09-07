package tor

import (
	"bytes"
	"testing"
)

func TestBuildSOCKS5AuthRequest(t *testing.T) {
	req := BuildSOCKS5AuthRequest()
	if !bytes.Equal(req, []byte{0x05, 0x01, 0x00}) {
		t.Errorf("unexpected socks5 request: %v", req)
	}
}
