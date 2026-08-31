package security

import "testing"

func TestIsExecutable(t *testing.T) {
	if !IsExecutable([]byte{0x7F, 'E', 'L', 'F'}) {
		t.Errorf("expected true for ELF binary")
	}
	if IsExecutable([]byte("ftypmp42")) {
		t.Errorf("expected false for MP4 media")
	}
}
