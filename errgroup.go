package waitgroup

import (
	"context"
	"sync"
)

// An ErrorGroup is a collection of goroutines working on subtasks that are part of
// the same overall task.
type ErrorGroup struct {
	size int
	pool chan byte

	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

// NewErrorGroup returns a new ErrorGroup instance
func NewErrorGroup(ctx context.Context, size int) (*ErrorGroup, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	wg := &ErrorGroup{
		size:   size,
		cancel: cancel,
	}
	if size > 0 {
		wg.pool = make(chan byte, size)
	}
	return wg, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *ErrorGroup) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Add calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *ErrorGroup) Add(f func() error) {

	if g.size > 0 {
		g.pool <- 1
	}
	g.wg.Add(1)

	go func() {
		defer func() {
			if g.size > 0 {
				<-g.pool
			}
			g.wg.Done()
		}()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
