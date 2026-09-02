package pivoting

type PivotManager struct {
	mirrors []string
	current int
}

func NewPivotManager(mirrors []string) *PivotManager {
	return &PivotManager{mirrors: mirrors}
}

func (p *PivotManager) PivotOnFailure() string {
	if len(p.mirrors) == 0 {
		return ""
	}
	p.current = (p.current + 1) % len(p.mirrors)
	return p.mirrors[p.current]
}
