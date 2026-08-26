package chunker

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChunkerSecurityInvalidURLs(t *testing.T) {
	invalidURLs := []string{
		"",
		"file:///etc/passwd",
		"ftp://example.com/file",
		"javascript:alert(1)",
		"--help",
		"-o",
		"http://",
	}

	for _, u := range invalidURLs {
		d := New(Config{
			URL:        u,
			TargetPath: "/tmp/out.bin",
			Quiet:      true,
		})
		err := d.Download(context.Background())
		if err == nil {
			t.Errorf("expected error for invalid URL %q, got nil", u)
		}
	}
}

func TestChunkerServerErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	d := New(Config{
		URL:        server.URL,
		TargetPath: "/tmp/out.bin",
		Quiet:      true,
	})

	err := d.Download(context.Background())
	if err == nil {
		t.Errorf("expected error on 500 server response, got nil")
	}
}
