package doh

import (
	"context"
	"testing"
)

func TestSecureResolver(t *testing.T) {
	res := NewSecureResolver("", true)
	ips, err := res.Resolve(context.Background(), "example.com")
	if err != nil || len(ips) == 0 {
		t.Fatalf("Resolve failed: %v", err)
	}
}
