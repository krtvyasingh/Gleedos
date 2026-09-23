package gdrive

import "testing"

func TestNewUploadRequest(t *testing.T) {
	req := NewUploadRequest("video.mp4", "root")
	if req.FileName != "video.mp4" {
		t.Errorf("unexpected upload request: %+v", req)
	}
}
