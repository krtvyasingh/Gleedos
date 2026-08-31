package storage

import "testing"

func TestPickBestVolume(t *testing.T) {
	vols := []Volume{
		{Path: "/mnt/disk1", FreeBytes: 1000},
		{Path: "/mnt/disk2", FreeBytes: 50000},
	}
	if PickBestVolume(vols) != "/mnt/disk2" {
		t.Errorf("expected /mnt/disk2")
	}
}
