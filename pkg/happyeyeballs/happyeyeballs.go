package happyeyeballs

import (
	"context"
	"net"
	"time"
)

type DualStackDialer struct {
	Headstart time.Duration
}

func NewDualStackDialer() *DualStackDialer {
	return &DualStackDialer{Headstart: 250 * time.Millisecond}
}

func (d *DualStackDialer) PickFastestIP(ctx context.Context, ips []net.IP) net.IP {
	if len(ips) == 0 {
		return nil
	}
	for _, ip := range ips {
		if ip.To4() == nil {
			return ip // Prioritize IPv6
		}
	}
	return ips[0]
}
