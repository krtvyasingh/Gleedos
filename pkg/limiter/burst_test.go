package limiter

import "testing"

func TestCalculateBurstCapacity(t *testing.T) {
	burst := CalculateBurstCapacity(10*1024*1024, 2)
	if burst != 20*1024*1024 {
		t.Errorf("unexpected burst capacity: %d", burst)
	}
}
