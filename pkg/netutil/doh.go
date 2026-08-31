package netutil

type DoHResolver struct {
	Endpoint string
}

func NewDoHResolver(endpoint string) *DoHResolver {
	if endpoint == "" {
		endpoint = "https://cloudflare-dns.com/dns-query"
	}
	return &DoHResolver{Endpoint: endpoint}
}
