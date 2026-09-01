package plugins

type FilterPlugin struct {
	Name    string
	Version string
}

func (p *FilterPlugin) TransformURL(rawURL string) string {
	return rawURL
}
