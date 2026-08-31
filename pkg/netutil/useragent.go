package netutil

var UserAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15",
}

func GetUserAgent(idx int) string {
	if len(UserAgents) == 0 {
		return "Gleedos/3.0"
	}
	return UserAgents[idx%len(UserAgents)]
}
