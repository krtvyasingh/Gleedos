package tor

import "testing"

func TestIsValidOnionV3(t *testing.T) {
	valid := "vww6ybal4bd7szmgncyruucpgfkqahzddi37ktceo3ah7ngmcopnpyyd.onion"
	invalid := "short.onion"
	if !IsValidOnionV3(valid) {
		t.Errorf("expected valid onion address")
	}
	if IsValidOnionV3(invalid) {
		t.Errorf("expected invalid onion address")
	}
}
