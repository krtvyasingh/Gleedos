package security

import "regexp"

var tokenReg = regexp.MustCompile(`(?i)(token|key|secret|password|api_key)=([^&\s]+)`)

func RedactSecrets(s string) string {
	return tokenReg.ReplaceAllString(s, "$1=REDACTED")
}
