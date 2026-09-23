package retention

import "time"

type RetentionPolicy struct {
	MaxAgeDays      int
	MinFreeDiskGB   int
}

func ShouldPruneFile(age time.Duration, policy RetentionPolicy) bool {
	return age > time.Duration(policy.MaxAgeDays)*24*time.Hour
}
