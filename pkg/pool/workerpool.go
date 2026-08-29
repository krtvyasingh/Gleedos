package pool

import (
	"context"
	"sync"
)

type Task func(ctx context.Context) error

type WorkerPool struct {
	concurrency int
	tasks       chan Task
	wg          sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	if concurrency <= 0 {
		concurrency = 4
	}
	return &WorkerPool{
		concurrency: concurrency,
		tasks:       make(chan Task, 1000),
	}
}

func (p *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < p.concurrency; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case t, ok := <-p.tasks:
					if !ok {
						return
					}
					_ = t(ctx)
				}
			}
		}()
	}
}

func (p *WorkerPool) Submit(t Task) {
	p.tasks <- t
}

func (p *WorkerPool) Stop() {
	close(p.tasks)
	p.wg.Wait()
}
