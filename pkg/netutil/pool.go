package netutil

import "net/http"

func TuneTransport(t *http.Transport) {
	if t != nil {
		t.MaxIdleConns = 100
		t.MaxIdleConnsPerHost = 10
		t.MaxConnsPerHost = 50
	}
}
