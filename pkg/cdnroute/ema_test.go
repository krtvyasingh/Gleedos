package cdnroute

import "testing"

func TestLatencyEMA(t *testing.T) {
	ema := NewLatencyEMA(0.5)
	v1 := ema.Update(100.0)
	v2 := ema.Update(50.0)
	if v1 != 100.0 || v2 != 75.0 {
		t.Errorf("unexpected EMA: %f, %f", v1, v2)
	}
}
