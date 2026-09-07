package nostr

import "strings"

func IsNpub(key string) bool {
	return strings.HasPrefix(key, "npub1") && len(key) == 63
}

func IsNsec(key string) bool {
	return strings.HasPrefix(key, "nsec1") && len(key) == 63
}
