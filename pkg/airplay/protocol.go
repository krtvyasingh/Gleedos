package airplay

import "strings"

func ParseAirPlayCommand(req string) string {
	if strings.Contains(req, "POST /play") {
		return "play"
	}
	if strings.Contains(req, "POST /stop") {
		return "stop"
	}
	return "unknown"
}
