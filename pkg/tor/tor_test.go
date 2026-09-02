package tor

import "testing"

func TestAnonymousDialer(t *testing.T) {
	d := NewAnonymousDialer("")
	if d.SOCKS5Proxy != "127.0.0.1:9050" {
		t.Errorf("unexpected default tor proxy: %s", d.SOCKS5Proxy)
	}
}
