package limiter

func CalculateBurstCapacity(rateBps int64, burstSeconds int) int64 {
	if rateBps <= 0 {
		return 0
	}
	if burstSeconds <= 0 {
		burstSeconds = 1
	}
	return rateBps * int64(burstSeconds)
}
