package plexhook

import (
	"context"
	"net/http"
	"time"
)

func NotifyMediaServer(ctx context.Context, webhookURL string) error {
	if webhookURL == "" {
		return nil
	}
	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, nil)
	if err != nil {
		return err
	}
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
