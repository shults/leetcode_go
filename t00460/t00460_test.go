package t00460_test

import (
	"github.com/stretchr/testify/assert"
	"leetcode_go/t00460"
	"testing"
)

func Test_LFU_cache(t *testing.T) {
	lfu := t00460.Constructor(2)
	lfu.Put(1, 1)                  // cache=[1,_], cnt(1)=1
	lfu.Put(2, 2)                  // cache=[2,1], cnt(2)=1, cnt(1)=1
	assert.Equal(t, 1, lfu.Get(1)) // return 1
	// cache=[1,2], cnt(2)=1, cnt(1)=2
	lfu.Put(3, 3) // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
	// cache=[3,1], cnt(3)=1, cnt(1)=2
	assert.Equal(t, -1, lfu.Get(2)) // return -1 (not found)
	assert.Equal(t, 3, lfu.Get(3))  // return 3
	// cache=[3,1], cnt(3)=2, cnt(1)=2
	lfu.Put(4, 4) // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	assert.Equal(t, -1, lfu.Get(1)) // return -1 (not found)
	assert.Equal(t, 3, lfu.Get(3))  // return 3
	// cache=[3,4], cnt(4)=1, cnt(3)=3
	assert.Equal(t, 4, lfu.Get(4)) // return 4
	// cache=[4,3], cnt(4)=2, cnt(3)=3
}

func TestLFUCache_Fail_Case(t *testing.T) {
	//["LFUCache", "put", "put", "put", "put","get", "get", "get", "get", "put","get", "get", "get", "get", "get"]
	//	[[3], [1, 1], [2, 2], [3, 3], [4, 4],[4], [3], [2], [1], [5, 5], [1], [2],[3], [4], [5]]

	lfu := t00460.Constructor(3)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	assert.Equal(t, 4, lfu.Get(4))
	assert.Equal(t, 3, lfu.Get(3))
	assert.Equal(t, 2, lfu.Get(2))
	assert.Equal(t, -1, lfu.Get(1))

}
