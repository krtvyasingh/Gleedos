package webrtc

import "strings"

func HasValidFingerprint(sdp string) bool {
	return strings.Contains(sdp, "a=fingerprint:sha-256") || strings.Contains(sdp, "a=fingerprint:sha-512")
}
