package t00414

import (
	"cmp"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"math"
)

func thirdMax(nums []int) int {
	heap := newUniqHeap(3)

	for _, val := range nums {
		heap.Push(val)
	}

	return heap.Top()
}

type uniqHeap struct {
	heap    *priorityqueue.Queue
	items   map[int]struct{}
	maxSize int
	maxVal  int
}

func newUniqHeap(maxSize int) *uniqHeap {
	return &uniqHeap{
		heap: priorityqueue.NewWith(func(a, b interface{}) int {
			return cmp.Compare(a.(int), b.(int))
		}),
		items:   make(map[int]struct{}, maxSize),
		maxSize: maxSize,
		maxVal:  math.MinInt,
	}
}

func (h *uniqHeap) Push(val int) {
	_, ok := h.items[val]

	if ok {
		return
	}

	h.heap.Enqueue(val)
	h.items[val] = struct{}{}
	h.maxVal = max(h.maxVal, val)

	for h.heap.Size() > h.maxSize {
		removedVal, _ := h.heap.Dequeue()
		delete(h.items, removedVal.(int))
	}
}

func (h *uniqHeap) Top() int {
	if h.heap.Size() == h.maxSize {
		res, _ := h.heap.Peek()
		return res.(int)
	} else {
		return h.maxVal
	}
}
