package pool

import (
	"context"
	"sync/atomic"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	pool := NewWorkerPool(3)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool.Start(ctx)

	var counter int32
	for i := 0; i < 10; i++ {
		pool.Submit(func(ctx context.Context) error {
			atomic.AddInt32(&counter, 1)
			return nil
		})
	}

	pool.Stop()

	if atomic.LoadInt32(&counter) != 10 {
		t.Errorf("expected 10 tasks executed, got %d", atomic.LoadInt32(&counter))
	}
}
