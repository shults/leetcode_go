package t00441

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 2, arrangeCoins(3))
	assert.Equal(t, 2, arrangeCoins(5))
	assert.Equal(t, 3, arrangeCoins(6))
	assert.Equal(t, 3, arrangeCoins(8))
	assert.Equal(t, 65535, arrangeCoins(2147483647))
}
