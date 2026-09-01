package bot

import "testing"

func TestBotDaemon(t *testing.T) {
	d := NewBotDaemon("token_tg", "token_dc")
	res := d.HandleMessage(BotMessage{ChatID: "123", URL: "https://example.com/video"})
	if res != "Queued download for: https://example.com/video" {
		t.Errorf("unexpected bot response: %s", res)
	}
}
