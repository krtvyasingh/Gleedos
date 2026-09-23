package mmapstream

import "testing"

func TestOpenVirtualMMap(t *testing.T) {
	mm := OpenVirtualMMap(1000)
	if mm.Length != 1000 {
		t.Errorf("unexpected mmap length")
	}
}
