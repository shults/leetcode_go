package t00414

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 1, thirdMax([]int{3, 2, 1}))
	assert.Equal(t, 1, thirdMax([]int{2, 2, 3, 1}))
}
