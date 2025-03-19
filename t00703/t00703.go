package t00703

import (
	"cmp"
	"github.com/emirpasic/gods/queues/priorityqueue"
)

type value int

type occurrences int

type KthLargest struct {
	k           int
	heap        *priorityqueue.Queue
	occurrences map[value]occurrences
	size        int
}

func Constructor(k int, nums []int) KthLargest {
	var heap = priorityqueue.NewWith(func(a, b interface{}) int {
		return cmp.Compare(a.(int), b.(int))
	})

	var instance = KthLargest{
		k:           k,
		heap:        heap,
		occurrences: make(map[value]occurrences),
		size:        0,
	}

	for _, num := range nums {
		_ = instance.Add(num)
	}

	return instance
}

func (k *KthLargest) Add(val int) int {
	if k.size < k.k {
		occur, _ := k.occurrences[value(val)]

		if occur == 0 {
			k.heap.Enqueue(val)
		}

		k.occurrences[value(val)]++
		k.size++

		peakVal, _ := k.heap.Peek()
		return peakVal.(int)
	}

	// k.size == k.k
	k.occurrences[value(val)]++
	k.heap.Enqueue(val)

	peakVal, _ := k.heap.Peek()
	peak := value(peakVal.(int))

	peakOccur := k.occurrences[peak]

	if peakOccur == 1 {
		_, _ = k.heap.Dequeue()
		delete(k.occurrences, peak)
	} else {
		k.occurrences[peak]--
	}

	return peakVal.(int)
}
