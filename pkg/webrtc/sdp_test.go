package webrtc

import "testing"

func TestHasValidFingerprint(t *testing.T) {
	valid := "v=0\r\na=fingerprint:sha-256 12:34:56\r\n"
	if !HasValidFingerprint(valid) {
		t.Errorf("expected valid fingerprint")
	}
}
