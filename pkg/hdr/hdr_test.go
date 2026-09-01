package hdr

import "testing"

func TestDetectHDRProfile(t *testing.T) {
	if DetectHDRProfile([]byte("...mdcv...")) != HDR10 {
		t.Errorf("expected HDR10 profile")
	}
	if DetectHDRProfile([]byte("...dvh1...")) != DolbyVision {
		t.Errorf("expected DolbyVision profile")
	}
}
