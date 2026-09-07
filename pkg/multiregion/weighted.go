package multiregion

type WeightedRegion struct {
	Region string
	Weight int
}

func PickWeightedRegion(regions []WeightedRegion, randVal int) string {
	if len(regions) == 0 {
		return ""
	}
	total := 0
	for _, r := range regions {
		total += r.Weight
	}
	if total == 0 {
		return regions[0].Region
	}
	target := randVal % total
	cum := 0
	for _, r := range regions {
		cum += r.Weight
		if target < cum {
			return r.Region
		}
	}
	return regions[0].Region
}
