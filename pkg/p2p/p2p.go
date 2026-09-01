package p2p

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Peer struct {
	ID       string
	Address  string
	LastSeen time.Time
}

type SwarmManager struct {
	peers map[string]*Peer
	mu    sync.RWMutex
}

func NewSwarmManager() *SwarmManager {
	return &SwarmManager{peers: make(map[string]*Peer)}
}

func (s *SwarmManager) AnnounceChunk(chunkHash string, port int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	peerID := fmt.Sprintf("peer_%d", time.Now().UnixNano())
	s.peers[peerID] = &Peer{
		ID:       peerID,
		Address:  fmt.Sprintf("127.0.0.1:%d", port),
		LastSeen: time.Now(),
	}
	return nil
}

func (s *SwarmManager) DiscoverPeers(ctx context.Context, chunkHash string) []*Peer {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var list []*Peer
	for _, p := range s.peers {
		list = append(list, p)
	}
	return list
}
