package tlscloak

func ValidateALPN(protos []string) bool {
	for _, p := range protos {
		if p == "h2" || p == "http/1.1" || p == "h3" {
			return true
		}
	}
	return false
}
