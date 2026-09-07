package carplay

import "fmt"

func FormatPlaylistItem(title, url, duration string) string {
	return fmt.Sprintf("<item><title>%s</title><enclosure url=\"%s\" type=\"audio/mpeg\"/><itunes:duration>%s</itunes:duration></item>", title, url, duration)
}
