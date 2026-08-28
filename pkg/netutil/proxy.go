package netutil

import (
	"net/http"
	"net/url"
)

func ConfigureProxy(proxyURL string) (*http.Transport, error) {
	t := http.DefaultTransport.(*http.Transport).Clone()
	if proxyURL != "" {
		parsed, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		t.Proxy = http.ProxyURL(parsed)
	}
	return t, nil
}
