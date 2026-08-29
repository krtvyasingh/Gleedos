package notify

import "testing"

func TestSendNotification(t *testing.T) {
	_ = SendNotification("Gleedos Test", "Download Complete")
}
