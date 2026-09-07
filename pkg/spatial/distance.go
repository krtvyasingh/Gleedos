package spatial

func CalculateDistanceGain(distanceMeters, referenceDistance, maxDistance float64) float64 {
	if distanceMeters <= referenceDistance {
		return 1.0
	}
	if distanceMeters >= maxDistance {
		return 0.0
	}
	return referenceDistance / (referenceDistance + (distanceMeters - referenceDistance))
}
