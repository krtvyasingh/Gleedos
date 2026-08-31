package security

import "bytes"

func IsExecutable(data []byte) bool {
	if len(data) >= 4 && bytes.Equal(data[:4], []byte{0x7F, 'E', 'L', 'F'}) {
		return true // Linux ELF
	}
	if len(data) >= 2 && bytes.Equal(data[:2], []byte{'M', 'Z'}) {
		return true // Windows PE
	}
	if len(data) >= 4 && (bytes.Equal(data[:4], []byte{0xFE, 0xED, 0xFA, 0xCE}) || bytes.Equal(data[:4], []byte{0xCF, 0xFA, 0xED, 0xFE})) {
		return true // macOS Mach-O
	}
	return false
}
