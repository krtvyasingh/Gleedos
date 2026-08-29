package ipc

import (
	"context"
	"encoding/json"
	"net"
	"path/filepath"
	"testing"
	"time"
)

func TestIPCServer(t *testing.T) {
	tmpDir := t.TempDir()
	sockPath := filepath.Join(tmpDir, "gleedos.sock")

	srv := NewServer(sockPath)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_ = srv.Start(ctx, func(m Message) Message {
			return Message{Command: "ACK", Payload: m.Payload}
		})
	}()

	time.Sleep(50 * time.Millisecond)

	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()

	_ = json.NewEncoder(conn).Encode(Message{Command: "PING", Payload: "123"})
	var res Message
	_ = json.NewDecoder(conn).Decode(&res)

	if res.Command != "ACK" || res.Payload != "123" {
		t.Errorf("unexpected response: %+v", res)
	}
}
