package t00022

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		size     int
		expected []string
	}

	for _, tCase := range []testCase{
		{0, []string{}},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	} {
		t.Run(fmt.Sprintf("%+v", tCase), func(t *testing.T) {
			assert.Subset(t, generateParenthesis(tCase.size), tCase.expected)
			assert.Subset(t, tCase.expected, generateParenthesis(tCase.size))
		})
	}
}
