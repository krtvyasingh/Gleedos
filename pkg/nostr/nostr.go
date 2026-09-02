package nostr

import "fmt"

type NostrMediaEvent struct {
	PubKey    string
	MediaURL  string
	SHA256Sum string
	MimeType  string
}

func FormatNIP94Tag(e NostrMediaEvent) string {
	return fmt.Sprintf("[\"url\",\"%s\"],[\"m\",\"%s\"],[\"x\",\"%s\"]", e.MediaURL, e.MimeType, e.SHA256Sum)
}
