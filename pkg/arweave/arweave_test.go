package arweave

import (
	"strings"
	"testing"
)

func TestFormatArweaveManifest(t *testing.T) {
	m := FormatArweaveManifest("tx_123", "video/mp4")
	if !strings.Contains(m, "tx_123") {
		t.Errorf("unexpected manifest: %s", m)
	}
}
