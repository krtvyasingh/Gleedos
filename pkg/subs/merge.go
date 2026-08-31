package subs

type SubTrack struct {
	Language string
	Path     string
}

func CombineTracks(tracks []SubTrack) int {
	return len(tracks)
}
