package actorindex

import "time"

type ActorAppearance struct {
	ActorName string
	Start     time.Duration
	End       time.Duration
}

func GroupAppearances(apps []ActorAppearance) map[string]int {
	counts := make(map[string]int)
	for _, a := range apps {
		counts[a.ActorName]++
	}
	return counts
}
