package batch

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// HistoryRecord tracks a completed or attempted download.
type HistoryRecord struct {
	URL          string    `json:"url"`
	CompletedAt  time.Time `json:"completed_at"`
	OutputPath   string    `json:"output_path"`
	Success      bool      `json:"success"`
	ErrorMessage string    `json:"error_message,omitempty"`
}

// HistoryStore persists and queries download history to avoid redundant work.
type HistoryStore struct {
	filePath string
	records  map[string]HistoryRecord
	mu       sync.RWMutex
}

// NewHistoryStore initializes the history store.
func NewHistoryStore(customPath string) (*HistoryStore, error) {
	path := customPath
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(home, ".gleedos", "history.json")
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	store := &HistoryStore{
		filePath: path,
		records:  make(map[string]HistoryRecord),
	}

	if data, err := os.ReadFile(path); err == nil {
		var list []HistoryRecord
		if json.Unmarshal(data, &list) == nil {
			for _, rec := range list {
				store.records[rec.URL] = rec
			}
		}
	}

	return store, nil
}

// HasURL returns true if the URL was already successfully downloaded.
func (h *HistoryStore) HasURL(u string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	rec, exists := h.records[u]
	return exists && rec.Success
}

// Record saves a download outcome.
func (h *HistoryStore) Record(rec HistoryRecord) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.records[rec.URL] = rec

	var list []HistoryRecord
	for _, r := range h.records {
		list = append(list, r)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(h.filePath, data, 0644)
}

// GetAll returns all history records.
func (h *HistoryStore) GetAll() []HistoryRecord {
	h.mu.RLock()
	defer h.mu.RUnlock()

	list := make([]HistoryRecord, 0, len(h.records))
	for _, r := range h.records {
		list = append(list, r)
	}
	return list
}

// ParseURLFile reads URLs from a text file, skipping empty lines and comments (#).
func ParseURLFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		urls = append(urls, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

// JobHandler executes a single download job in the batch.
type JobHandler func(ctx context.Context, u string) (string, error)

// ProcessBatch runs batch downloads with concurrency and history deduplication.
func ProcessBatch(
	ctx context.Context,
	urls []string,
	concurrency int,
	store *HistoryStore,
	handler JobHandler,
) (successCount int, failCount int, err error) {
	if concurrency <= 0 {
		concurrency = 3
	}

	type queueItem struct {
		index int
		url   string
	}

	items := make(chan queueItem, len(urls))
	for i, u := range urls {
		items <- queueItem{index: i, url: u}
	}
	close(items)

	var wg sync.WaitGroup
	var mu sync.Mutex

	workerCount := concurrency
	if workerCount > len(urls) {
		workerCount = len(urls)
	}

	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range items {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if store != nil && store.HasURL(item.url) {
					mu.Lock()
					successCount++
					mu.Unlock()
					continue
				}

				outPath, err := handler(ctx, item.url)
				rec := HistoryRecord{
					URL:         item.url,
					CompletedAt: time.Now(),
					OutputPath:  outPath,
					Success:     err == nil,
				}
				if err != nil {
					rec.ErrorMessage = err.Error()
				}

				if store != nil {
					_ = store.Record(rec)
				}

				mu.Lock()
				if err == nil {
					successCount++
				} else {
					failCount++
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return successCount, failCount, nil
}
