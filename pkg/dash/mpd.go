package dash

import (
	"encoding/xml"
	"io"
)

type MPD struct {
	XMLName xml.Name `xml:"MPD"`
	Type    string   `xml:"type,attr"`
	Period  []Period `xml:"Period"`
}

type Period struct {
	AdaptationSet []AdaptationSet `xml:"AdaptationSet"`
}

type AdaptationSet struct {
	MimeType       string           `xml:"mimeType,attr"`
	Representation []Representation `xml:"Representation"`
}

type Representation struct {
	ID        string `xml:"id,attr"`
	Bandwidth int64  `xml:"bandwidth,attr"`
	Width     int    `xml:"width,attr"`
	Height    int    `xml:"height,attr"`
}

func ParseMPD(r io.Reader) (*MPD, error) {
	var mpd MPD
	if err := xml.NewDecoder(r).Decode(&mpd); err != nil {
		return nil, err
	}
	return &mpd, nil
}
