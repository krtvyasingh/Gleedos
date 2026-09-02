package hotreload

import "net"

type ReloadManager struct {
	Listener net.Listener
}

func (r *ReloadManager) IsListening() bool {
	return r.Listener != nil
}
