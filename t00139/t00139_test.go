package t00139

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		input    string
		dict     []string
		expected bool
	}

	for _, tCase := range []testCase{
		//{"leetcode", []string{"leet", "code"}, true},
		//{"applepenapple", []string{"apple", "pen"}, true},
		//{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		//{"bb", []string{"a", "b", "bbb", "bbbb"}, true},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			false,
		},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"},
			false, // ????
		},
		{
			"aaaaaaa",
			[]string{"aaa", "aaaa"},
			true, // ????
		},
	} {
		t.Run(fmt.Sprintf("%+v\n", tCase), func(t *testing.T) {
			assert.Equal(t, tCase.expected, wordBreak(tCase.input, tCase.dict))
		})
	}

}
