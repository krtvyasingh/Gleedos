package turnstilehook

type SolverEvent struct {
	SiteKey string
	Action  string
}

func NewSolverEvent(siteKey, action string) SolverEvent {
	return SolverEvent{SiteKey: siteKey, Action: action}
}
