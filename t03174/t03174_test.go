package t03174

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "abc", clearDigits("abc"))
	assert.Equal(t, "", clearDigits("ab12"))
	assert.Equal(t, "de", clearDigits("ab12de"))
}
