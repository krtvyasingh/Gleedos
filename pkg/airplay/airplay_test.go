package airplay

import "testing"

func TestDiscoverDevices(t *testing.T) {
	devs := DiscoverDevices()
	if len(devs) != 1 || devs[0].Type != "airplay" {
		t.Errorf("unexpected devices: %+v", devs)
	}
}
