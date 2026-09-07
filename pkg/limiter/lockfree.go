package limiter

import "sync/atomic"

type LockFreeBucket struct {
	tokens int64
}

func NewLockFreeBucket(capacity int64) *LockFreeBucket {
	return &LockFreeBucket{tokens: capacity}
}

func (l *LockFreeBucket) TryConsume(count int64) bool {
	for {
		curr := atomic.LoadInt64(&l.tokens)
		if curr < count {
			return false
		}
		if atomic.CompareAndSwapInt64(&l.tokens, curr, curr-count) {
			return true
		}
	}
}
