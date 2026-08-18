package main

import (
	"strings"
	"testing"
)

func strategyArgValue(args []string, key string) string {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == key {
			return args[i+1]
		}
	}
	return ""
}

func strategyHasArg(args []string, value string) bool {
	for _, arg := range args {
		if arg == value {
			return true
		}
	}
	return false
}

func TestBuildDownloadStrategiesAudio(t *testing.T) {
	base := []string{"--newline", "-o", "output.%(ext)s"}

	strategies := buildDownloadStrategies(
		base,
		"best",
		true,
		"https://example.com/watch?v=test",
	)

	if len(strategies) != 3 {
		t.Fatalf("audio strategy count = %d, want 3", len(strategies))
	}

	wantNames := []string{
		"mweb audio",
		"web_safari audio",
		"android audio",
	}

	wantClients := []string{
		"youtube:player_client=mweb",
		"youtube:player_client=web_safari",
		"youtube:player_client=android",
	}

	for i, strategy := range strategies {
		if strategy.name != wantNames[i] {
			t.Fatalf(
				"audio strategy %d name = %q, want %q",
				i,
				strategy.name,
				wantNames[i],
			)
		}

		if strategyArgValue(strategy.args, "--extractor-args") != wantClients[i] {
			t.Fatalf(
				"audio strategy %d client = %q, want %q",
				i,
				strategyArgValue(strategy.args, "--extractor-args"),
				wantClients[i],
			)
		}

		if strategyArgValue(strategy.args, "-f") != "ba/b" {
			t.Fatalf(
				"audio strategy %d format = %q, want ba/b",
				i,
				strategyArgValue(strategy.args, "-f"),
			)
		}

		if !strategyHasArg(strategy.args, "--extract-audio") {
			t.Fatalf("audio strategy %d missing --extract-audio", i)
		}

		if strategyArgValue(strategy.args, "--audio-format") != "mp3" {
			t.Fatalf("audio strategy %d missing mp3 format", i)
		}

		if strategyArgValue(strategy.args, "--audio-quality") != "0" {
			t.Fatalf("audio strategy %d missing audio quality 0", i)
		}

		if strategy.args[len(strategy.args)-1] != "https://example.com/watch?v=test" {
			t.Fatalf("audio strategy %d missing URL at end", i)
		}
	}
}

func TestBuildDownloadStrategiesVideo(t *testing.T) {
	base := []string{"--newline", "-o", "output.%(ext)s"}
	url := "https://example.com/watch?v=test"

	strategies := buildDownloadStrategies(
		base,
		"bestvideo+bestaudio/best",
		false,
		url,
	)

	if len(strategies) != 6 {
		t.Fatalf("video strategy count = %d, want 6", len(strategies))
	}

	wantNames := []string{
		"mweb best",
		"web safari progressive MP4",
		"android progressive MP4",
		"web safari 720p",
		"android 720p",
		"web safari best",
	}

	for i, strategy := range strategies {
		if strategy.name != wantNames[i] {
			t.Fatalf(
				"video strategy %d name = %q, want %q",
				i,
				strategy.name,
				wantNames[i],
			)
		}

		if strategy.args[len(strategy.args)-1] != url {
			t.Fatalf("video strategy %d URL is not last argument", i)
		}

		if !strategyHasArg(strategy.args, "--merge-output-format") {
			t.Fatalf("video strategy %d missing MP4 merge option", i)
		}

		if strategyArgValue(strategy.args, "--merge-output-format") != "mp4" {
			t.Fatalf("video strategy %d merge format is not mp4", i)
		}
	}
}

func TestBuildDownloadStrategiesPreservesBaseArguments(t *testing.T) {
	base := []string{
		"--newline",
		"--progress",
		"--no-playlist",
		"-o",
		"test-output.%(ext)s",
	}

	original := append([]string(nil), base...)

	_ = buildDownloadStrategies(
		base,
		"best",
		false,
		"https://example.com/watch?v=test",
	)

	if strings.Join(base, "\x00") != strings.Join(original, "\x00") {
		t.Fatalf("base arguments were mutated")
	}
}

func TestBuildDownloadStrategiesUsesRequestedQuality(t *testing.T) {
	base := []string{"--newline"}
	quality := "bv*[height<=1080]+ba/b"

	strategies := buildDownloadStrategies(
		base,
		quality,
		false,
		"https://example.com/watch?v=test",
	)

	if len(strategies) == 0 {
		t.Fatal("no video strategies returned")
	}

	if strategyArgValue(strategies[0].args, "-f") != quality {
		t.Fatalf(
			"first video format = %q, want %q",
			strategyArgValue(strategies[0].args, "-f"),
			quality,
		)
	}
}

func TestBuildDownloadStrategiesKeepsProvenAndroidFallback(t *testing.T) {
	base := []string{"--newline"}

	strategies := buildDownloadStrategies(
		base,
		"best",
		false,
		"https://example.com/watch?v=test",
	)

	var found bool

	for _, strategy := range strategies {
		if strategy.name != "android progressive MP4" {
			continue
		}

		found = true

		if strategyArgValue(
			strategy.args,
			"--extractor-args",
		) != "youtube:player_client=android" {
			t.Fatal("Android fallback uses the wrong extractor client")
		}

		if strategyArgValue(
			strategy.args,
			"-f",
		) != "18/b[ext=mp4]/b" {
			t.Fatal("Android fallback format changed")
		}
	}

	if !found {
		t.Fatal("proven Android progressive MP4 fallback is missing")
	}
}
