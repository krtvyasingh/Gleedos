package webhook

import "testing"

func TestFormatSlackMessage(t *testing.T) {
	msg := FormatSlackMessage("Podcast", "/audio/pod.mp3")
	if msg.Text == "" {
		t.Errorf("empty slack message")
	}
}
