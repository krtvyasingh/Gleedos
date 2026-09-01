package webdav

import "net/http"

type WebDAVClient struct {
	Endpoint string
}

func NewWebDAVClient(endpoint string) *WebDAVClient {
	return &WebDAVClient{Endpoint: endpoint}
}

func (w *WebDAVClient) CreatePutRequest(path string) (*http.Request, error) {
	return http.NewRequest("PUT", w.Endpoint+path, nil)
}
