package t00168

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "Z", convertToTitle(26))
	assert.Equal(t, "AA", convertToTitle(27))
	assert.Equal(t, "AB", convertToTitle(28))
	assert.Equal(t, "ZY", convertToTitle(701))
}
