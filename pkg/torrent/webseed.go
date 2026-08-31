package torrent

type WebSeed struct {
	URL string
}

func NewWebSeed(u string) WebSeed {
	return WebSeed{URL: u}
}
