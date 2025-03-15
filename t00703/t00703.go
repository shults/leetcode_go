package t00703

import (
	"cmp"
	"github.com/emirpasic/gods/queues/priorityqueue"
)

type KthLargest struct {
	k    int
	heap *priorityqueue.Queue
	item map[int]int
	size int
}

func Constructor(k int, nums []int) KthLargest {
	var heap = priorityqueue.NewWith(func(a, b interface{}) int {
		return cmp.Compare(a.(int), b.(int))
	})

	var instance = KthLargest{
		k:    k,
		heap: heap,
		item: make(map[int]int),
		size: 0,
	}

	for _, num := range nums {
		_ = instance.Add(num)
	}

	return instance
}

func (k *KthLargest) Add(val int) int {
	if k.size < k.k {

	} else {

	}

	return 0
}
