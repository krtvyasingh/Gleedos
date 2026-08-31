package netutil

import (
	"net/http"
	"testing"
)

func TestTuneTransport(t *testing.T) {
	tr := &http.Transport{}
	TuneTransport(tr)
	if tr.MaxIdleConns != 100 || tr.MaxIdleConnsPerHost != 10 {
		t.Errorf("unexpected transport tuning: %+v", tr)
	}
}
