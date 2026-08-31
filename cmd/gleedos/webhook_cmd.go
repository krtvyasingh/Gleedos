package main

import (
	"context"
	"time"

	"github.com/krtvysingh/gleedos/pkg/webhook"
)

func notifyWebhook(webhookURL, url, path string, success bool) {
	if webhookURL == "" {
		return
	}
	_ = webhook.Dispatch(context.Background(), webhookURL, webhook.Event{
		Event:     "download.event",
		URL:       url,
		Path:      path,
		Success:   success,
		Timestamp: time.Now(),
	})
}
