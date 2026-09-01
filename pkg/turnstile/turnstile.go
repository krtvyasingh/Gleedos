package turnstile

import (
	"encoding/json"
	"net/http"
	"sync"
)

type ChallengeSolution struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
}

type Harvester struct {
	solutions map[string]string
	mu        sync.RWMutex
}

func NewHarvester() *Harvester {
	return &Harvester{solutions: make(map[string]string)}
}

func (h *Harvester) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var sol ChallengeSolution
	if err := json.NewDecoder(r.Body).Decode(&sol); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.mu.Lock()
	h.solutions[sol.SessionID] = sol.Token
	h.mu.Unlock()
	w.WriteHeader(http.StatusOK)
}
