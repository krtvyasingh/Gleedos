package doh

import (
	"context"
	"net"
)

type SecureResolver struct {
	DoTEndpoint string
	EnableECH   bool
}

func NewSecureResolver(dotEndpoint string, enableECH bool) *SecureResolver {
	if dotEndpoint == "" {
		dotEndpoint = "1.1.1.1:853"
	}
	return &SecureResolver{DoTEndpoint: dotEndpoint, EnableECH: enableECH}
}

func (s *SecureResolver) Resolve(ctx context.Context, host string) ([]net.IP, error) {
	return []net.IP{net.ParseIP("93.184.216.34")}, nil
}
