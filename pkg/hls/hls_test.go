package hls

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParsePlaylist(t *testing.T) {
	m3u8Content := `#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:10
#EXT-X-MEDIA-SEQUENCE:0
#EXTINF:9.009,
segment0.ts
#EXTINF:9.009,
segment1.ts
#EXTINF:3.003,
segment2.ts
#EXT-X-ENDLIST`

	baseURL, _ := url.Parse("https://example.com/hls/stream.m3u8")
	playlist, err := ParsePlaylist(baseURL, strings.NewReader(m3u8Content))
	if err != nil {
		t.Fatalf("ParsePlaylist failed: %v", err)
	}

	if playlist.IsMaster {
		t.Errorf("expected media playlist, got master")
	}

	if len(playlist.Segments) != 3 {
		t.Fatalf("expected 3 segments, got %d", len(playlist.Segments))
	}

	if playlist.Segments[0].URL != "https://example.com/hls/segment0.ts" {
		t.Errorf("segment 0 URL unexpected: %s", playlist.Segments[0].URL)
	}
}

func TestHLSDownloader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "playlist.m3u8") {
			w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
			fmt.Fprintln(w, "#EXTM3U")
			fmt.Fprintln(w, "#EXTINF:2.0,")
			fmt.Fprintln(w, "chunk_0.ts")
			fmt.Fprintln(w, "#EXTINF:2.0,")
			fmt.Fprintln(w, "chunk_1.ts")
			fmt.Fprintln(w, "#EXT-X-ENDLIST")
			return
		}

		if strings.HasSuffix(r.URL.Path, "chunk_0.ts") {
			w.Write([]byte("VIDEO_PART_0_DATA"))
			return
		}

		if strings.HasSuffix(r.URL.Path, "chunk_1.ts") {
			w.Write([]byte("VIDEO_PART_1_DATA"))
			return
		}

		http.NotFound(w, r)
	}))
	defer server.Close()

	tmpDir, err := os.MkdirTemp("", "gleedos_hls_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	targetFile := filepath.Join(tmpDir, "output.ts")

	downloader := New(Config{
		URL:         server.URL + "/playlist.m3u8",
		TargetPath:  targetFile,
		Concurrency: 2,
		Quiet:       true,
	})

	err = downloader.Download(context.Background())
	if err != nil {
		t.Fatalf("HLS download failed: %v", err)
	}

	result, err := os.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("read target file failed: %v", err)
	}

	expected := "VIDEO_PART_0_DATAVIDEO_PART_1_DATA"
	if string(result) != expected {
		t.Fatalf("HLS merged data mismatch! expected %q, got %q", expected, string(result))
	}
}
