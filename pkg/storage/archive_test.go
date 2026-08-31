package storage

import "testing"

func TestIsSupportedArchive(t *testing.T) {
	if !IsSupportedArchive("bundle.zip") || !IsSupportedArchive("media.tar.gz") {
		t.Errorf("expected true for supported archives")
	}
}
