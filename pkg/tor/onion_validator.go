package tor

import "strings"

func IsValidOnionV3(addr string) bool {
	if !strings.HasSuffix(addr, ".onion") {
		return false
	}
	host := strings.TrimSuffix(addr, ".onion")
	return len(host) == 56
}
