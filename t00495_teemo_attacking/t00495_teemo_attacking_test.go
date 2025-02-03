package t00495_teemo_attacking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPoisonedDuration(t *testing.T) {
	assert.Equal(t, 4, findPoisonedDuration([]int{1, 4}, 2))
	assert.Equal(t, 3, findPoisonedDuration([]int{1, 2}, 2))
}
