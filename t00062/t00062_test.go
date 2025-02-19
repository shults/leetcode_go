package t00062

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 6, uniquePaths(3, 3))
	assert.Equal(t, 28, uniquePaths(3, 7))
}
