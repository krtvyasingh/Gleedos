package chunker

import (
	"bytes"
	"testing"

	"github.com/krtvysingh/gleedos/pkg/tagger"
	"github.com/krtvysingh/gleedos/pkg/urlutil"
)

func BenchmarkDetectMediaType(b *testing.B) {
	mp4Header := []byte{0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p', 'm', 'p', '4', '2'}
	for i := 0; i < b.N; i++ {
		_, _ = tagger.DetectMediaType(bytes.NewReader(mp4Header))
	}
}

func BenchmarkCleanURL(b *testing.B) {
	raw := "https://example.com/video?utm_source=test&v=123"
	for i := 0; i < b.N; i++ {
		_ = urlutil.CleanURL(raw)
	}
}
