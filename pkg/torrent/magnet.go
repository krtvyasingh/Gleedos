package torrent

import (
	"errors"
	"net/url"
	"strings"
)

type MagnetURI struct {
	ExactTopic  string
	DisplayName string
}

func ParseMagnet(uri string) (*MagnetURI, error) {
	if !strings.HasPrefix(uri, "magnet:?") {
		return nil, errors.New("invalid magnet URI prefix")
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	return &MagnetURI{
		ExactTopic:  q.Get("xt"),
		DisplayName: q.Get("dn"),
	}, nil
}
