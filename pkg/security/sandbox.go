package security

import (
	"os"
	"path/filepath"
)

func CreateSandboxDir(base string) (string, error) {
	dir := filepath.Join(base, ".gleedos_sandbox")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	return dir, nil
}
