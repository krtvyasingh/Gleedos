package server

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var (
	TotalDownloadsSuccess uint64
	TotalDownloadsFailed  uint64
	TotalBytesDownloaded  uint64
)

func HandleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "# HELP gleedos_downloads_total Total completed downloads\n")
	fmt.Fprintf(w, "# TYPE gleedos_downloads_total counter\n")
	fmt.Fprintf(w, "gleedos_downloads_total{status=\"success\"} %d\n", atomic.LoadUint64(&TotalDownloadsSuccess))
	fmt.Fprintf(w, "gleedos_downloads_total{status=\"failed\"} %d\n", atomic.LoadUint64(&TotalDownloadsFailed))
	fmt.Fprintf(w, "# HELP gleedos_bytes_total Total downloaded bytes\n")
	fmt.Fprintf(w, "# TYPE gleedos_bytes_total counter\n")
	fmt.Fprintf(w, "gleedos_bytes_total %d\n", atomic.LoadUint64(&TotalBytesDownloaded))
}
