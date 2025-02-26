package t00245

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "AceCreIm", reverseVowels("IceCreAm"))
	assert.Equal(t, " ", reverseVowels(" "))
}
