package mpquic

type PrioritySubflow struct {
	Subflow  *Subflow
	Priority int
}

func PickHighestPriority(subflows []PrioritySubflow) *Subflow {
	if len(subflows) == 0 {
		return nil
	}
	best := subflows[0]
	for _, sf := range subflows {
		if sf.Priority > best.Priority {
			best = sf
		}
	}
	return best.Subflow
}
