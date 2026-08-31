package security

import (
	"errors"
	"path/filepath"
	"strings"
)

func IsSafePath(baseDir, userPath string) (string, error) {
	cleanBase := filepath.Clean(baseDir)
	cleanPath := filepath.Clean(filepath.Join(cleanBase, userPath))
	if !strings.HasPrefix(cleanPath, cleanBase) {
		return "", errors.New("path traversal detected")
	}
	return cleanPath, nil
}
