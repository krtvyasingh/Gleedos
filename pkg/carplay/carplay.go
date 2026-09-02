package carplay

import "fmt"

func GenerateCarPlayFeed(channelTitle, mediaURL string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>%s</title>
    <item>
      <enclosure url="%s" type="audio/mpeg"/>
    </item>
  </channel>
</rss>`, channelTitle, mediaURL)
}
