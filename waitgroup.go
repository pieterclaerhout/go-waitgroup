package waitgroup

import (
	"sync"
	"sync/atomic"
)

// WaitGroup implements a simple goruntine pool.
type WaitGroup struct {
	size      int
	pool      chan byte
	waitCount int64
	waitGroup sync.WaitGroup
}

// NewWaitGroup creates a waitgroup with a specific size (the maximum number of
// goroutines to run at the same time). If you use -1 as the size, all items
// will run concurrently (just like a normal sync.WaitGroup)
func NewWaitGroup(size int) *WaitGroup {
	wg := &WaitGroup{
		size: size,
	}
	if size > 0 {
		wg.pool = make(chan byte, size)
	}
	return wg
}

// BlockAdd pushes ‘one’ into the group. Blocks if the group is full.
func (wg *WaitGroup) BlockAdd() {
	if wg.size > 0 {
		wg.pool <- 1
	}
	wg.waitGroup.Add(1)
	atomic.AddInt64(&wg.waitCount, 1)
}

// Done pops ‘one’ out the group.
func (wg *WaitGroup) Done() {
	if wg.size > 0 {
		<-wg.pool
	}
	wg.waitGroup.Done()
	atomic.AddInt64(&wg.waitCount, -1)
}

// Wait waiting the group empty
func (wg *WaitGroup) Wait() {
	wg.waitGroup.Wait()
}

// PendingCount returns the number of pending operations
func (wg *WaitGroup) PendingCount() int64 {
	return wg.waitCount
}
