package notify

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func SendNotification(title, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch runtime.GOOS {
	case "darwin":
		script := fmt.Sprintf("display notification %q with title %q", message, title)
		return exec.CommandContext(ctx, "osascript", "-e", script).Run()
	case "linux":
		if _, err := exec.LookPath("notify-send"); err == nil {
			return exec.CommandContext(ctx, "notify-send", title, message).Run()
		}
	}
	return nil
}
