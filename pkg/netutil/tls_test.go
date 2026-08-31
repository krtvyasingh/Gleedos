package netutil

import (
	"crypto/tls"
	"testing"
)

func TestSecureTLSConfig(t *testing.T) {
	cfg := SecureTLSConfig()
	if cfg.MinVersion != tls.VersionTLS12 {
		t.Errorf("expected TLS 1.2 minimum version")
	}
}
