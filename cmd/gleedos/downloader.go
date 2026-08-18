package main

import (
	"os"
	"os/exec"
)

func runDownload(strategy downloadStrategy) error {
	cmd := exec.Command("yt-dlp", strategy.args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
