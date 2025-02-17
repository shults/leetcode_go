package t00013

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type testCase struct {
		roman string
		i     int
	}

	for _, tc := range []testCase{
		{"XX", 20},
		{"XXI", 21},
		{"MCMXCIV", 1994},
	} {
		t.Run(fmt.Sprintf("%s => %d", tc.roman, tc.i), func(t *testing.T) {
			assert.Equal(t, tc.i, romanToInt(tc.roman))
		})
	}

}
