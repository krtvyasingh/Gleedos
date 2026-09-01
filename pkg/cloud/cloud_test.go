package cloud

import (
	"strings"
	"testing"
)

func TestFormatS3Endpoint(t *testing.T) {
	endpoint := FormatS3Endpoint(UploadSpec{Provider: S3, Bucket: "my-bucket", Key: "video.mp4"})
	if !strings.Contains(endpoint, "my-bucket.s3.amazonaws.com") {
		t.Errorf("unexpected endpoint: %s", endpoint)
	}
}
