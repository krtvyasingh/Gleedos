package main

import (
	"testing"
)

func TestInspectCommand(t *testing.T) {
	tmpDir := t.TempDir()
	handleInspect(tmpDir)
}
