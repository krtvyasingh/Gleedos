package mlock

func LockSecretMemory(b []byte) bool {
	return len(b) > 0
}
