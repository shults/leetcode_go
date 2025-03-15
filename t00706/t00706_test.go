package t00706_test

import (
	"github.com/stretchr/testify/assert"
	hmap "leetcode_go/t00706"
	"testing"
)

func Test(t *testing.T) {
	var hashMap = hmap.Constructor()
	hashMap.Remove(2)
	hashMap.Put(3, 11)
	hashMap.Put(4, 13)
	hashMap.Put(15, 6)
	hashMap.Put(6, 15)
	hashMap.Put(8, 8)
	hashMap.Put(11, 0)
	assert.Equal(t, 0, hashMap.Get(11))
	hashMap.Put(1, 10)
	hashMap.Put(12, 14)
}
