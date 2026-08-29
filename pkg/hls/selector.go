package hls

import "sort"

type VariantStream struct {
	URL        string
	Bandwidth  int64
	Resolution string
}

func SelectBestVariant(variants []VariantStream) VariantStream {
	if len(variants) == 0 {
		return VariantStream{}
	}
	sorted := make([]VariantStream, len(variants))
	copy(sorted, variants)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Bandwidth > sorted[j].Bandwidth
	})

	return sorted[0]
}
