package cdnroute

import "testing"

func TestSelectOptimalEdge(t *testing.T) {
	edges := []EdgeMirror{
		{IP: "1.1.1.1", LatencyMs: 50.0, LossRate: 0.1},
		{IP: "2.2.2.2", LatencyMs: 20.0, LossRate: 0.0},
	}
	best := SelectOptimalEdge(edges)
	if best.IP != "2.2.2.2" {
		t.Errorf("expected 2.2.2.2, got %+v", best)
	}
}
