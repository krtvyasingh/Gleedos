package config

import (
	"path/filepath"
	"testing"
)

func TestConfigLoadSave(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "config.json")

	cfg := DefaultConfig()
	cfg.Threads = 8
	cfg.DefaultQuality = "best"

	if err := cfg.Save(cfgPath); err != nil {
		t.Fatalf("Save config failed: %v", err)
	}

	loaded, err := Load(cfgPath)
	if err != nil {
		t.Fatalf("Load config failed: %v", err)
	}

	if loaded.Threads != 8 || loaded.DefaultQuality != "best" {
		t.Errorf("Loaded config mismatch: %+v", loaded)
	}
}
