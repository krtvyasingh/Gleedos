package webrtc

import "testing"

func TestParseICECandidate(t *testing.T) {
	cand := "candidate:1 1 UDP 2130706431 192.168.1.1 5000 typ host"
	if !ParseICECandidate(cand) {
		t.Errorf("expected valid candidate")
	}
}
