package torrentdav

import "fmt"

type MagnetGateway struct {
	Port int
}

func (m *MagnetGateway) GetStreamURL(infoHash string) string {
	return fmt.Sprintf("http://127.0.0.1:%d/stream/%s", m.Port, infoHash)
}
