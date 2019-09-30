package waitgroup_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-waitgroup"
)

func Test_ErrorGroup_Add(t *testing.T) {

	type test struct {
		name string
		size int
	}

	var tests = []test{
		{"single", 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ctx := context.Background()

			wg, _ := waitgroup.NewErrorGroup(ctx, tc.size)
			assert.NotNil(t, wg)

			wg.Add(func() error {
				return nil
			})

			wg.Add(func() error {
				return errors.New("An error occurred")
			})

			err := wg.Wait()
			assert.Error(t, err)

		})
	}

}
