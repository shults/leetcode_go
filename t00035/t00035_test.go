package t00035

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	for _, c := range []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", c), func(t *testing.T) {
			receiver := searchInsert(c.input, c.target)
			if receiver != c.expected {
				t.Errorf("fail received=%d expected=%d", receiver, c.expected)
			}
		})
	}

}
