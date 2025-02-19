package t00063

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	_ = grid

	assert.Equal(t, 2, uniquePathsWithObstacles(grid))
}
