package hooks

import (
	"context"
	"os/exec"
	"time"
)

func ExecuteHook(scriptPath, filePath string) error {
	if scriptPath == "" {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return exec.CommandContext(ctx, scriptPath, filePath).Run()
}
