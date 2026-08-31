package hls

import (
	"bufio"
	"io"
	"strings"
)

type IPTVChannel struct {
	Name string
	URL  string
}

func ParseIPTV(r io.Reader) []IPTVChannel {
	var channels []IPTVChannel
	scanner := bufio.NewScanner(r)
	var currentName string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#EXTINF:") {
			parts := strings.Split(line, ",")
			if len(parts) > 1 {
				currentName = strings.TrimSpace(parts[1])
			}
		} else if strings.HasPrefix(line, "http") && currentName != "" {
			channels = append(channels, IPTVChannel{Name: currentName, URL: line})
			currentName = ""
		}
	}
	return channels
}
