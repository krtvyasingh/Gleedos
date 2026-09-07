package mpquic

import "testing"

func TestRoundRobinBalancer(t *testing.T) {
	b := NewRoundRobinBalancer()
	sf1 := &Subflow{InterfaceName: "wlan0"}
	sf2 := &Subflow{InterfaceName: "eth0"}
	list := []*Subflow{sf1, sf2}

	first := b.NextSubflow(list)
	second := b.NextSubflow(list)
	if first == second {
		t.Errorf("expected alternating subflows")
	}
}
