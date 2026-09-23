package dht

type PeerEndpoint struct {
	IP   string
	Port int
}

func FormatPeerCompact(p PeerEndpoint) string {
	return p.IP
}
