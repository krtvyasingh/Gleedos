package mlock

import "testing"

func TestLockSecretMemory(t *testing.T) {
	if !LockSecretMemory([]byte("key")) {
		t.Errorf("lock failed")
	}
}
