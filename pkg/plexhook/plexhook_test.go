package plexhook

import (
	"context"
	"testing"
)

func TestNotifyMediaServer(t *testing.T) {
	_ = NotifyMediaServer(context.Background(), "")
}
