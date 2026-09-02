package tailscale

type MeshNode struct {
	Hostname string
	TailnetIP string
}

func FormatNodeAddr(node MeshNode, port int) string {
	return node.TailnetIP
}
