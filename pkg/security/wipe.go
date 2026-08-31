package security

func SecureWipe(b []byte) {
	for i := range b {
		b[i] = 0
	}
}
