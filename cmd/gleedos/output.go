package main

import (
	"os"
	"path/filepath"
	"strings"
)

func isTemporaryOutput(name string) bool {
	return strings.HasSuffix(name, ".part") ||
		strings.HasSuffix(name, ".ytdl") ||
		strings.HasSuffix(name, ".temp") ||
		strings.HasSuffix(name, ".tmp")
}

type outputSnapshot struct {
	size    int64
	modTime int64
}

func snapshotOutput(dir string) map[string]outputSnapshot {
	out := make(map[string]outputSnapshot)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return out
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if isTemporaryOutput(name) {
			continue
		}

		path := filepath.Join(dir, name)

		info, err := entry.Info()
		if err != nil {
			continue
		}

		out[path] = outputSnapshot{
			size:    info.Size(),
			modTime: info.ModTime().UnixNano(),
		}
	}

	return out
}

func cleanupTemporaryFiles(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if isTemporaryOutput(name) {
			_ = os.Remove(filepath.Join(dir, name))
		}
	}
}

func findCompletedOutput(
	before map[string]outputSnapshot,
	dir string,
) string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if isTemporaryOutput(name) {
			continue
		}

		path := filepath.Join(dir, name)

		info, err := entry.Info()
		if err != nil || info.Size() <= 0 {
			continue
		}

		oldSnapshot, existed := before[path]

		currentSnapshot := outputSnapshot{
			size:    info.Size(),
			modTime: info.ModTime().UnixNano(),
		}

		if !existed ||
			oldSnapshot.size != currentSnapshot.size ||
			oldSnapshot.modTime != currentSnapshot.modTime {
			return path
		}
	}

	return ""
}
