package t00055

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	for _, tCase := range []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{1, 2}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 0}, true},
	} {
		assert.Equal(t, tCase.expected, canJump(tCase.nums), fmt.Sprintf("%+v", tCase))
	}
}
