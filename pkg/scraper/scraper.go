package scraper

import "regexp"

var mediaURLRegex = regexp.MustCompile(`https?://[^\s"']+\.(?:mp4|m3u8|mp3|flac|mkv|webm)`)

func ExtractMediaURLs(html string) []string {
	return mediaURLRegex.FindAllString(html, -1)
}
