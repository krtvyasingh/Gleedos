package tlscloak

import "crypto/tls"

type BrowserProfile string

const (
	Chrome  BrowserProfile = "chrome"
	Firefox BrowserProfile = "firefox"
	Safari  BrowserProfile = "safari"
)

type JA4Fingerprint struct {
	Profile    BrowserProfile
	CipherSuites []uint16
	Curves     []tls.CurveID
	ALPN       []string
}

func GetCamouflageConfig(profile BrowserProfile) *tls.Config {
	switch profile {
	case Firefox:
		return &tls.Config{
			MinVersion: tls.VersionTLS13,
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256, tls.CurveP384},
			NextProtos: []string{"h2", "http/1.1"},
		}
	case Safari:
		return &tls.Config{
			MinVersion: tls.VersionTLS13,
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
			NextProtos: []string{"h2", "http/1.1"},
		}
	default:
		return &tls.Config{
			MinVersion: tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256, tls.CurveP384, tls.CurveP521},
			NextProtos: []string{"h2", "http/1.1"},
		}
	}
}
