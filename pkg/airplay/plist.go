package airplay

import "fmt"

func FormatPlaybackInfoPlist(duration, position float64) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><plist version="1.0"><dict><key>duration</key><real>%f</real><key>position</key><real>%f</real></dict></plist>`, duration, position)
}
