package storage

type Volume struct {
	Path      string
	FreeBytes int64
}

func PickBestVolume(vols []Volume) string {
	if len(vols) == 0 {
		return ""
	}
	best := vols[0]
	for _, v := range vols {
		if v.FreeBytes > best.FreeBytes {
			best = v
		}
	}
	return best.Path
}
