package tlscloak

import (
	"crypto/tls"
	"testing"
)

func TestGetCamouflageConfig(t *testing.T) {
	cfg := GetCamouflageConfig(Chrome)
	if cfg.MinVersion != tls.VersionTLS12 || len(cfg.NextProtos) == 0 {
		t.Errorf("unexpected Chrome TLS config: %+v", cfg)
	}
}
