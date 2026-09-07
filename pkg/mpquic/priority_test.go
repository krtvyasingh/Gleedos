package mpquic

import "testing"

func TestPickHighestPriority(t *testing.T) {
	s1 := &Subflow{InterfaceName: "wlan0"}
	s2 := &Subflow{InterfaceName: "eth0"}
	list := []PrioritySubflow{
		{Subflow: s1, Priority: 1},
		{Subflow: s2, Priority: 10},
	}
	best := PickHighestPriority(list)
	if best != s2 {
		t.Errorf("expected eth0 to be selected")
	}
}
