package ws

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Hub struct {
	clients map[chan []byte]bool
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{clients: make(map[chan []byte]bool)}
}

func (h *Hub) Broadcast(v any) {
	data, _ := json.Marshal(v)
	h.mu.Lock()
	defer h.mu.Unlock()
	for ch := range h.clients {
		select {
		case ch <- data:
		default:
		}
	}
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
