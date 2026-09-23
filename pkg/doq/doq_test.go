package doq

import (
	"context"
	"testing"
)

func TestDOQResolver(t *testing.T) {
	r := NewDOQResolver("")
	ips, err := r.ResolveHost(context.Background(), "example.com")
	if err != nil || len(ips) == 0 {
		t.Fatalf("ResolveHost failed: %v", err)
	}
}
