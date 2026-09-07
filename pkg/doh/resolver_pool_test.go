package doh

import "testing"

func TestResolverPool(t *testing.T) {
	pool := NewResolverPool([]string{"ep1", "ep2"})
	e1 := pool.NextEndpoint()
	e2 := pool.NextEndpoint()
	if e1 == e2 || e1 == "" {
		t.Errorf("unexpected endpoint rotation: %s, %s", e1, e2)
	}
}
