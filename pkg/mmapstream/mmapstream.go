package mmapstream

type MMapBuffer struct {
	Length int64
}

func OpenVirtualMMap(length int64) *MMapBuffer {
	return &MMapBuffer{Length: length}
}
