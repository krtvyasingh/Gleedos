package webhook

import "testing"

func TestFormatTelegramMessage(t *testing.T) {
	msg := FormatTelegramMessage("123456", "Done")
	if msg.ChatID != "123456" || msg.Text != "Done" {
		t.Errorf("unexpected telegram message: %+v", msg)
	}
}
