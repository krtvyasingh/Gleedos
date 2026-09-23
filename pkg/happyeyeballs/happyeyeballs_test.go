package happyeyeballs

import (
	"context"
	"net"
	"testing"
)

func TestDualStackDialer(t *testing.T) {
	d := NewDualStackDialer()
	v4 := net.ParseIP("192.168.1.1")
	v6 := net.ParseIP("2001:db8::1")
	best := d.PickFastestIP(context.Background(), []net.IP{v4, v6})
	if best.String() != v6.String() {
		t.Errorf("expected IPv6 to be prioritized")
	}
}
