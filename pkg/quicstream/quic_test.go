package quicstream

import (
	"context"
	"testing"
)

func TestQUICManager(t *testing.T) {
	qm := NewQUICManager()
	ok, err := qm.DialQUIC(context.Background(), "example.com:443")
	if err != nil || !ok {
		t.Errorf("DialQUIC failed: %v", err)
	}
}
