package webrtc

import "testing"

func TestCreateOffer(t *testing.T) {
	s := CreateOffer("sess_1")
	if s.SessionID != "sess_1" || s.SDP == "" {
		t.Errorf("unexpected session: %+v", s)
	}
}
