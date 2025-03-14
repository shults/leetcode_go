package t00434

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 5, countSegments("Hello, my name is John"))
	assert.Equal(t, 0, countSegments(""))
	assert.Equal(t, 1, countSegments("Hello"))

}
