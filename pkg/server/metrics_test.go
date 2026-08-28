package server

import (
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

func TestHandleMetrics(t *testing.T) {
	atomic.StoreUint64(&TotalDownloadsSuccess, 5)
	atomic.StoreUint64(&TotalBytesDownloaded, 102400)

	req := httptest.NewRequest("GET", "/metrics", nil)
	rec := httptest.NewRecorder()
	HandleMetrics(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, `gleedos_downloads_total{status="success"} 5`) {
		t.Errorf("missing success metrics in output: %s", body)
	}
}
