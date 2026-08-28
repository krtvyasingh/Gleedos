package signals

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func SetupSignalHandler() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
}
