package t00219

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		nums   []int
		k      int
		output bool
	}

	for _, tCase := range []testCase{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	} {
		t.Run(fmt.Sprintf("test-case: %+v", tCase), func(t *testing.T) {
			require.Equal(t, tCase.output, containsNearbyDuplicate(tCase.nums, tCase.k))
		})
	}
}
