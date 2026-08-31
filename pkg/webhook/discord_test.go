package webhook

import "testing"

func TestFormatDiscordMessage(t *testing.T) {
	p := FormatDiscordMessage("Video Title", "/downloads/vid.mp4")
	if len(p.Embeds) != 1 || p.Embeds[0].Title != "Video Title" {
		t.Errorf("unexpected Discord payload: %+v", p)
	}
}
