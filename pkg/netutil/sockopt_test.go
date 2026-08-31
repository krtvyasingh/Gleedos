package netutil

import "testing"

func TestDefaultSocketOptions(t *testing.T) {
	opts := DefaultSocketOptions()
	if !opts.TCPNoDelay || opts.ReadBufferSize != 64*1024 {
		t.Errorf("unexpected socket options: %+v", opts)
	}
}
