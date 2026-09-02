package tor

type AnonymousDialer struct {
	SOCKS5Proxy string
}

func NewAnonymousDialer(proxy string) *AnonymousDialer {
	if proxy == "" {
		proxy = "127.0.0.1:9050"
	}
	return &AnonymousDialer{SOCKS5Proxy: proxy}
}
