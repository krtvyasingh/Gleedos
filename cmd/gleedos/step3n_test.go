package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestStep3NSameSizeReplacementIsDetected(t *testing.T) {
	dir := t.TempDir()
	output := filepath.Join(dir, "video.mp4")

	if err := os.WriteFile(output, []byte("before"), 0644); err != nil {
		t.Fatal(err)
	}
	before := snapshotOutput(dir)

	updated := time.Now().Add(time.Second)
	if err := os.WriteFile(output, []byte("after!"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(output, updated, updated); err != nil {
		t.Fatal(err)
	}

	if found := findCompletedOutput(before, dir); found != output {
		t.Fatalf("found %q, want same-size replacement %q", found, output)
	}
}

func TestStep3NFailedAttemptOutputIsNotDetectedAfterNewSnapshot(t *testing.T) {
	dir := t.TempDir()
	output := filepath.Join(dir, "failed-attempt.mp4")

	if err := os.WriteFile(output, []byte("incomplete output"), 0644); err != nil {
		t.Fatal(err)
	}

	beforeLaterAttempt := snapshotOutput(dir)
	if found := findCompletedOutput(beforeLaterAttempt, dir); found != "" {
		t.Fatalf("found stale failed-attempt output %q", found)
	}
}

func TestStep3NTemporaryOutputsAreIgnored(t *testing.T) {
	dir := t.TempDir()
	before := snapshotOutput(dir)

	for _, name := range []string{"video.mp4.part", "video.ytdl", "video.temp", "video.tmp"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("partial"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	if found := findCompletedOutput(before, dir); found != "" {
		t.Fatalf("found temporary output %q", found)
	}
}
