package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsTemporaryOutput(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"video.mp4.part", true},
		{"video.ytdl", true},
		{"video.temp", true},
		{"video.tmp", true},
		{"video.mp4", false},
		{"video.mkv", false},
	}

	for _, tt := range tests {
		if got := isTemporaryOutput(tt.name); got != tt.want {
			t.Fatalf(
				"isTemporaryOutput(%q) = %v, want %v",
				tt.name,
				got,
				tt.want,
			)
		}
	}
}

func TestSnapshotOutputIgnoresTemporaryFiles(t *testing.T) {
	dir := t.TempDir()

	files := map[string]string{
		"video.mp4":      "completed",
		"video.mp4.part": "partial",
		"video.ytdl":     "temporary",
		"video.temp":     "temporary",
		"video.tmp":      "temporary",
	}

	for name, contents := range files {
		if err := os.WriteFile(
			filepath.Join(dir, name),
			[]byte(contents),
			0644,
		); err != nil {
			t.Fatal(err)
		}
	}

	got := snapshotOutput(dir)

	if len(got) != 1 {
		t.Fatalf("snapshot contains %d files, want 1", len(got))
	}

	expected := filepath.Join(dir, "video.mp4")

	if _, ok := got[expected]; !ok {
		t.Fatalf("completed output %q missing from snapshot", expected)
	}
}

func TestCleanupTemporaryFiles(t *testing.T) {
	dir := t.TempDir()

	keep := filepath.Join(dir, "video.mp4")

	files := []string{
		keep,
		filepath.Join(dir, "video.mp4.part"),
		filepath.Join(dir, "video.ytdl"),
		filepath.Join(dir, "video.temp"),
		filepath.Join(dir, "video.tmp"),
	}

	for _, path := range files {
		if err := os.WriteFile(path, []byte("data"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	cleanupTemporaryFiles(dir)

	if _, err := os.Stat(keep); err != nil {
		t.Fatalf("completed output was removed: %v", err)
	}

	for _, path := range files[1:] {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("temporary file still exists: %s", path)
		}
	}
}

func TestFindCompletedOutput(t *testing.T) {
	dir := t.TempDir()

	before := snapshotOutput(dir)

	output := filepath.Join(dir, "video.mp4")

	if err := os.WriteFile(
		output,
		[]byte("real video"),
		0644,
	); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(
		filepath.Join(dir, "video.mp4.part"),
		[]byte("partial"),
		0644,
	); err != nil {
		t.Fatal(err)
	}

	found := findCompletedOutput(before, dir)

	if found != output {
		t.Fatalf("found %q, want %q", found, output)
	}
}

func TestFindCompletedOutputRejectsEmptyFile(t *testing.T) {
	dir := t.TempDir()

	before := snapshotOutput(dir)

	output := filepath.Join(dir, "empty.mp4")

	if err := os.WriteFile(output, nil, 0644); err != nil {
		t.Fatal(err)
	}

	found := findCompletedOutput(before, dir)

	if found != "" {
		t.Fatalf("found empty output %q", found)
	}
}
