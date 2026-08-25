package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/krtvysingh/gleedos/pkg/batch"
)

// JobStatus tracks the state of a queued/running download.
type JobStatus string

const (
	StatusQueued      JobStatus = "queued"
	StatusDownloading JobStatus = "downloading"
	StatusCompleted   JobStatus = "completed"
	StatusFailed      JobStatus = "failed"
)

// DownloadRequest represents the incoming JSON body to /api/download.
type DownloadRequest struct {
	URL        string `json:"url"`
	Format     string `json:"format,omitempty"`
	AudioMode  bool   `json:"audio,omitempty"`
	Best       bool   `json:"best,omitempty"`
	Turbo      bool   `json:"turbo,omitempty"`
	OutputPath string `json:"output,omitempty"`
}

// JobResponse is a serializable DTO for job status.
type JobResponse struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Status     JobStatus `json:"status"`
	OutputPath string    `json:"output_path,omitempty"`
	Error      string    `json:"error,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// DownloadJob represents an active or finished job with internal synchronization.
type DownloadJob struct {
	ID         string
	URL        string
	Status     JobStatus
	OutputPath string
	Error      string
	CreatedAt  time.Time
	mu         sync.RWMutex
}

// Response returns a thread-safe copy of the job as JobResponse for serialization.
func (j *DownloadJob) Response() JobResponse {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return JobResponse{
		ID:         j.ID,
		URL:        j.URL,
		Status:     j.Status,
		OutputPath: j.OutputPath,
		Error:      j.Error,
		CreatedAt:  j.CreatedAt,
	}
}

// HandlerFunc executes the actual download for a job.
type HandlerFunc func(ctx context.Context, req DownloadRequest) (string, error)

// Server is the HTTP REST API server.
type Server struct {
	port    int
	handler HandlerFunc
	store   *batch.HistoryStore
	jobs    map[string]*DownloadJob
	mu      sync.RWMutex
	httpSrv *http.Server
}

// New creates a new Server instance.
func New(port int, handler HandlerFunc, store *batch.HistoryStore) *Server {
	if port <= 0 {
		port = 8080
	}
	return &Server{
		port:    port,
		handler: handler,
		store:   store,
		jobs:    make(map[string]*DownloadJob),
	}
}

// Start runs the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/api/download", s.handleDownload)
	mux.HandleFunc("/api/status", s.handleStatus)
	mux.HandleFunc("/api/history", s.handleHistory)

	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = s.httpSrv.Shutdown(shutdownCtx)
	}()

	err := s.httpSrv.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "ok",
		"service": "gleedos-api",
		"version": "2.0.0",
	})
}

func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req DownloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, `{"error": "invalid request body or missing url"}`, http.StatusBadRequest)
		return
	}

	jobID := fmt.Sprintf("job_%d", time.Now().UnixNano())
	job := &DownloadJob{
		ID:        jobID,
		URL:       req.URL,
		Status:    StatusQueued,
		CreatedAt: time.Now(),
	}

	s.mu.Lock()
	s.jobs[jobID] = job
	s.mu.Unlock()

	response := job.Response()

	go func() {
		job.mu.Lock()
		job.Status = StatusDownloading
		job.mu.Unlock()

		outPath, err := s.handler(context.Background(), req)

		job.mu.Lock()
		if err != nil {
			job.Status = StatusFailed
			job.Error = err.Error()
		} else {
			job.Status = StatusCompleted
			job.OutputPath = outPath
		}
		job.mu.Unlock()

		if s.store != nil {
			rec := batch.HistoryRecord{
				URL:         req.URL,
				CompletedAt: time.Now(),
				OutputPath:  outPath,
				Success:     err == nil,
			}
			if err != nil {
				rec.ErrorMessage = err.Error()
			}
			_ = s.store.Record(rec)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error": "missing id parameter"}`, http.StatusBadRequest)
		return
	}

	s.mu.RLock()
	job, exists := s.jobs[id]
	s.mu.RUnlock()

	if !exists {
		http.Error(w, `{"error": "job not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job.Response())
}

func (s *Server) handleHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if s.store == nil {
		json.NewEncoder(w).Encode([]batch.HistoryRecord{})
		return
	}
	json.NewEncoder(w).Encode(s.store.GetAll())
}
