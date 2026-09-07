package tailscale

import "strings"

func IsTailnetIP(ip string) bool {
	return strings.HasPrefix(ip, "100.64.") || strings.HasPrefix(ip, "100.")
}
