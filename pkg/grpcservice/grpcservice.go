package grpcservice

type DownloadRequest struct {
	URL        string
	Quality    string
	OutputPath string
}

type DownloadResponse struct {
	Status  string
	Bytes   int64
	Success bool
}

type ServiceServer struct{}

func (s *ServiceServer) ProcessDownload(req DownloadRequest) DownloadResponse {
	return DownloadResponse{Status: "Completed", Bytes: 1024, Success: true}
}
