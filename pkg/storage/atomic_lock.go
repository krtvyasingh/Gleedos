package storage

import (
	"os"
	"path/filepath"
)

func SafeWriteAtomic(targetPath string, data []byte) error {
	tmpPath := targetPath + ".tmp"
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil && filepath.Dir(targetPath) != "." {
		return err
	}
	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmpPath, targetPath)
}
