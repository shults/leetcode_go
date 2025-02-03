package t00067_add_binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "11", addBinary("10", "01"))
	assert.Equal(t, "110", addBinary("11", "11"))
	assert.Equal(t, "0", addBinary("0", "0"))
}

func Test_getVal(t *testing.T) {
	assert.Equal(t, 1, getVal("1", 1, 2))
	assert.Equal(t, 0, getVal("1", 0, 2))
	assert.Equal(t, 1, getVal("1", 2, 3))
	assert.Equal(t, 0, getVal("1", 1, 3))
	assert.Equal(t, 0, getVal("1", 0, 3))

	assert.Equal(t, 1, getVal("101", 2, 3))
	assert.Equal(t, 0, getVal("101", 1, 3))
	assert.Equal(t, 1, getVal("101", 0, 3))
}
