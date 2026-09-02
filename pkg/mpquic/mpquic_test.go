package mpquic

import "testing"

func TestMultipathScheduler(t *testing.T) {
	s := NewMultipathScheduler()
	s.AddSubflow("wlan0", 50*1024*1024, 15)
	s.AddSubflow("rmnet0", 80*1024*1024, 25)
	if s.GetTotalBandwidth() != 130*1024*1024 {
		t.Errorf("unexpected total bandwidth: %d", s.GetTotalBandwidth())
	}
}
