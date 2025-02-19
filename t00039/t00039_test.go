package t00039

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

type testCase struct {
	input          []int
	target         int
	expectedResult [][]int
}

func (r *testCase) accepts(results [][]int) bool {
	if len(results) != len(r.expectedResult) {
		return false
	}

	for _, entry := range results {
		if !r.contains(entry) {
			return false
		}
	}

	return true
}

func (r *testCase) contains(val []int) bool {
	for _, entry := range r.expectedResult {
		if slices.Equal(entry, val) {
			return true
		}
	}

	return false
}

func Test(t *testing.T) {
	for _, tCase := range []testCase{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tCase), func(t *testing.T) {
			results := combinationSum(tCase.input, tCase.target)
			assert.True(t, tCase.accepts(results), fmt.Sprintf("received: %+v\nexpected: %+v", results, tCase.expectedResult))
		})
	}
}
