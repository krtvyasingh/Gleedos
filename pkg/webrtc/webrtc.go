package webrtc

type WebRTCSession struct {
	SessionID string
	SDP       string
}

func CreateOffer(sessionID string) WebRTCSession {
	return WebRTCSession{SessionID: sessionID, SDP: "v=0\r\no=Gleedos..."}
}
