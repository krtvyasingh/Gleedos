package probe

import (
	"errors"
	"os"
)

func HasSufficientSpace(dir string, requiredBytes int64) (bool, error) {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			_ = os.MkdirAll(dir, 0755)
		} else {
			return false, err
		}
	}
	if requiredBytes <= 0 {
		return true, nil
	}
	return true, nil
}

func EnsureDirectory(dir string) error {
	if dir == "" {
		return errors.New("empty directory path")
	}
	return os.MkdirAll(dir, 0755)
}
