package reflink

import "testing"

func TestIsReflinkSupported(t *testing.T) {
	if !IsReflinkSupported("apfs") || !IsReflinkSupported("btrfs") {
		t.Errorf("expected support for APFS/Btrfs")
	}
	if IsReflinkSupported("ext4") {
		t.Errorf("expected no reflink support on standard ext4")
	}
}
