package feed

import (
	"encoding/xml"
	"io"
)

type RSSFeed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title     string    `xml:"title"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	URL string `xml:"url,attr"`
}

func ParseRSS(r io.Reader) (*RSSFeed, error) {
	var feed RSSFeed
	if err := xml.NewDecoder(r).Decode(&feed); err != nil {
		return nil, err
	}
	return &feed, nil
}
