package t00202

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	require.True(t, isHappy(7))
}
