package grpcservice

import "testing"

func TestServiceServer(t *testing.T) {
	srv := &ServiceServer{}
	res := srv.ProcessDownload(DownloadRequest{URL: "https://example.com/video"})
	if !res.Success || res.Status != "Completed" {
		t.Errorf("unexpected response: %+v", res)
	}
}
