package netutil

import "testing"

func TestDoHResolver(t *testing.T) {
	r := NewDoHResolver("")
	if r.Endpoint != "https://cloudflare-dns.com/dns-query" {
		t.Errorf("unexpected default DoH endpoint: %s", r.Endpoint)
	}
}
