package mobile

import "testing"

func TestParseMobileSharePayload(t *testing.T) {
	raw := "Check out this video: https://youtu.be/123456 shared via mobile"
	url := ParseMobileSharePayload(raw)
	if url != "https://youtu.be/123456" {
		t.Errorf("unexpected url: %s", url)
	}
}
