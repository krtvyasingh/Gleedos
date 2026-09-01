package webdav

import "testing"

func TestWebDAVClient(t *testing.T) {
	client := NewWebDAVClient("https://nextcloud.example.com/remote.php/webdav")
	req, err := client.CreatePutRequest("/video.mp4")
	if err != nil || req.Method != "PUT" {
		t.Fatalf("CreatePutRequest failed: %v", err)
	}
}
