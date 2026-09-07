package tor

func BuildSOCKS5AuthRequest() []byte {
	return []byte{0x05, 0x01, 0x00} // SOCKS5, 1 auth method, NO AUTH
}
