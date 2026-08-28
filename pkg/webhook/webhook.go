package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Event struct {
	Event     string    `json:"event"`
	URL       string    `json:"url"`
	Path      string    `json:"path"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

func Dispatch(ctx context.Context, webhookURL string, event Event) error {
	if webhookURL == "" {
		return nil
	}
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
