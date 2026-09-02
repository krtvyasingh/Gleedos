package mqtt

import "testing"

func TestFormatMQTTTopic(t *testing.T) {
	top := FormatMQTTTopic("gleedos/downloads", "complete")
	if top != "gleedos/downloads/complete" {
		t.Errorf("unexpected topic: %s", top)
	}
}
