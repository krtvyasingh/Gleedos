package webrtc

import "strings"

type ICECandidate struct {
	Foundation string
	Protocol   string
	IP         string
	Port       int
}

func ParseICECandidate(cand string) bool {
	return strings.HasPrefix(cand, "candidate:")
}
