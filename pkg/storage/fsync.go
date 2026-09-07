package storage

import "os"

func SyncFileToDisk(f *os.File) error {
	if f == nil {
		return nil
	}
	return f.Sync()
}
