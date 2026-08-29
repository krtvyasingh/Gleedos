package demux

import (
	"bytes"
	"testing"
)

func TestReadTSPacket(t *testing.T) {
	packet := make([]byte, PacketSize)
	packet[0] = SyncByte
	packet[1] = 0x01
	packet[2] = 0x00

	pkt, err := ReadTSPacket(bytes.NewReader(packet))
	if err != nil {
		t.Fatalf("ReadTSPacket failed: %v", err)
	}
	if pkt.PID != 256 {
		t.Errorf("expected PID 256, got %d", pkt.PID)
	}
}
