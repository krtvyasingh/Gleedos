package ipc

import (
	"context"
	"encoding/json"
	"net"
	"os"
)

type Message struct {
	Command string `json:"command"`
	Payload string `json:"payload"`
}

type Server struct {
	SocketPath string
	listener   net.Listener
}

func NewServer(socketPath string) *Server {
	return &Server{SocketPath: socketPath}
}

func (s *Server) Start(ctx context.Context, handler func(msg Message) Message) error {
	_ = os.Remove(s.SocketPath)
	l, err := net.Listen("unix", s.SocketPath)
	if err != nil {
		return err
	}
	s.listener = l

	go func() {
		<-ctx.Done()
		_ = s.listener.Close()
		_ = os.Remove(s.SocketPath)
	}()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return nil
		}
		go func(c net.Conn) {
			defer c.Close()
			var msg Message
			if err := json.NewDecoder(c).Decode(&msg); err == nil {
				res := handler(msg)
				_ = json.NewEncoder(c).Encode(res)
			}
		}(conn)
	}
}
