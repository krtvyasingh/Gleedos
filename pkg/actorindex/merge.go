package actorindex

import "time"

func MergeIntervals(apps []ActorAppearance) []ActorAppearance {
	if len(apps) <= 1 {
		return apps
	}
	var merged []ActorAppearance
	curr := apps[0]
	for i := 1; i < len(apps); i++ {
		next := apps[i]
		if next.ActorName == curr.ActorName && next.Start <= curr.End+time.Second {
			if next.End > curr.End {
				curr.End = next.End
			}
		} else {
			merged = append(merged, curr)
			curr = next
		}
	}
	merged = append(merged, curr)
	return merged
}
