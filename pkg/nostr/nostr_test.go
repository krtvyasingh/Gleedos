package nostr

import (
	"strings"
	"testing"
)

func TestFormatNIP94Tag(t *testing.T) {
	tag := FormatNIP94Tag(NostrMediaEvent{MediaURL: "https://example.com/v.mp4", MimeType: "video/mp4", SHA256Sum: "123456"})
	if !strings.Contains(tag, "video/mp4") {
		t.Errorf("unexpected NIP94 tag: %s", tag)
	}
}
