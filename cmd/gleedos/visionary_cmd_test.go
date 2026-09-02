package main

import (
	"path/filepath"
	"testing"
)

func TestVisionaryCommandHandlers(t *testing.T) {
	handlePQCDemo()
	handleWireframeDemo()
	tmpDir := t.TempDir()
	handleUniversalExportDemo(filepath.Join(tmpDir, "export.json"), "json")
}
