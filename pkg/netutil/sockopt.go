package netutil

const (
	DefaultReadBufferSize  = 64 * 1024
	DefaultWriteBufferSize = 64 * 1024
)

type SocketOptions struct {
	ReadBufferSize  int
	WriteBufferSize int
	TCPNoDelay      bool
}

func DefaultSocketOptions() SocketOptions {
	return SocketOptions{
		ReadBufferSize:  DefaultReadBufferSize,
		WriteBufferSize: DefaultWriteBufferSize,
		TCPNoDelay:      true,
	}
}
