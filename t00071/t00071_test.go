package t00071

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	for _, c := range []testCase{
		{"/home//foo/", "/home/foo"},
		{"/home/", "/home"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
		{"/", "/"},
		{"/.././GVzvE/./xBjU///../..///././//////T/../../.././zu/q/e", "/zu/q/e"},
	} {
		t.Run(c.input, func(t *testing.T) {
			assert.Equal(t, c.expected, simplifyPath(c.input))
		})

	}
}
