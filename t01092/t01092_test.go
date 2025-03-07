package t01092

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	for _, tCase := range []struct {
		expected string
		str1     string
		str2     string
	}{
		{"cabac", "abac", "cab"},
		{"aaaaaaaa", "aaaaaaaa", "aaaaaaaa"},
		{"bbbaaababbb", "bbbaaaba", "bbababbb"},
		{"bbabcacccaaab", "bcacaaab", "bbabaccc"},
	} {
		t.Run(fmt.Sprintf("%+v", tCase), func(t *testing.T) {
			assert.Equal(t, tCase.expected, shortestCommonSupersequence(tCase.str1, tCase.str2))
		})
	}
}

type Saver interface {
	Save()
}

type UserSaver struct{}

func (u *UserSaver) Save() {}

func calSaver(s Saver) {
	fmt.Println(s == nil)
	s.Save()
}

func TestInterface(t *testing.T) {
	var s *UserSaver
	calSaver(s)
}
