package t00069_sqrtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(2))
	assert.Equal(t, 1, mySqrt(3))
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(5))
	assert.Equal(t, 2, mySqrt(6))
	assert.Equal(t, 2, mySqrt(7))
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 3, mySqrt(9))
	assert.Equal(t, 46340, mySqrt(2147395600)) // rakamakafo: it breaks requirements params cause of x should be in range 1 <= x 1 <= 2**31 - 1
}
