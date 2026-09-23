package falloc

type PreallocPlan struct {
	TotalBytes int64
	BlockSize  int64
}

func PlanPreallocation(total int64) PreallocPlan {
	return PreallocPlan{TotalBytes: total, BlockSize: 4096}
}
