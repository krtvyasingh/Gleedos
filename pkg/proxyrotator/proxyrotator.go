package proxyrotator

import "sync"

type ProxyNode struct {
	URL      string
	Country  string
	Failures int
}

type ProxyPool struct {
	proxies []*ProxyNode
	mu      sync.Mutex
	index   int
}

func NewProxyPool(proxies []string) *ProxyPool {
	var nodes []*ProxyNode
	for _, p := range proxies {
		nodes = append(nodes, &ProxyNode{URL: p})
	}
	return &ProxyPool{proxies: nodes}
}

func (p *ProxyPool) Next() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.proxies) == 0 {
		return ""
	}
	node := p.proxies[p.index%len(p.proxies)]
	p.index++
	return node.URL
}
