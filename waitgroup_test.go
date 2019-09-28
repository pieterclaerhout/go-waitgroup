package waitgroup_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-waitgroup"
)

func Test_Waitgroup(t *testing.T) {

	type test struct {
		name string
		size int
	}

	var tests = []test{
		{"single", 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			wg := waitgroup.NewWaitGroup(tc.size)
			assert.NotNil(t, wg)
			assert.Zero(t, wg.PendingCount(), "pending-before")

			wg.BlockAdd()
			go func() {
				defer wg.Done()
			}()

			assert.EqualValues(t, 1, wg.PendingCount(), "pending-during")

			wg.Wait()

			assert.Zero(t, wg.PendingCount(), "pending-after")

		})
	}

}
