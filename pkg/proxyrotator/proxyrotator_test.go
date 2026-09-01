package proxyrotator

import "testing"

func TestProxyPool(t *testing.T) {
	pool := NewProxyPool([]string{"http://proxy1:8080", "http://proxy2:8080"})
	p1 := pool.Next()
	p2 := pool.Next()
	if p1 == p2 || p1 == "" {
		t.Errorf("unexpected proxy rotation: %s, %s", p1, p2)
	}
}
