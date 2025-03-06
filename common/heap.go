package common

type CompareFn[T any] func(a T, b T) int

type Heap[T comparable] struct {
	size    int
	compare CompareFn[T]
	items   []T
}

func NewHeap[T comparable](fn CompareFn[T]) *Heap[T] {
	return &Heap[T]{
		compare: fn,
		size:    0,
		items:   make([]T, 8),
	}
}

func (h *Heap[T]) Empty() bool {
	return h.size == 0
}

func (h *Heap[T]) Pop() T {
	if h.Empty() {
		panic("empty binary heap")
	}

	var val = h.items[0]

	h.size--
	h.items[0] = h.items[h.size]
	h.maxHeapify(0)

	return val
}

func (h *Heap[T]) Peek() T {
	if h.Empty() {
		panic("empty binary heap")
	}

	return h.items[0]
}

func (h *Heap[T]) Push(val T) {
	h.grow()
	h.items[h.size] = val
	h.bubleUp(h.size)
	h.size++
}

func (h *Heap[T]) bubleUp(i int) {
	for {
		if i < 0 {
			break
		}

		var parent = parentIndex(i)

		if h.compare(h.items[parent], h.items[i]) < 0 {
			h.items[parent], h.items[i] = h.items[i], h.items[parent]
		} else {
			return
		}

		i = parent
	}
}

func (h *Heap[T]) maxHeapify(i int) {
	left, right := children(i)
	//left := 2*(i+1) - 1
	//right := left + 1
	largest := i

	if left < h.size && h.compare(h.items[left], h.items[largest]) > 0 {
		largest = left
	}

	if right < h.size && h.compare(h.items[right], h.items[largest]) > 0 {
		largest = right
	}

	if largest != i {
		h.items[largest], h.items[i] = h.items[i], h.items[largest]
		h.maxHeapify(largest)
	}
}

// grow if limit exceeded
func (h *Heap[T]) grow() {
	if len(h.items) != h.size {
		return
	}

	var items = make([]T, h.size*2)
	copy(items, h.items)
	h.items = items
}

func children(i int) (left, right int) {
	left = 2*(i+1) - 1
	right = 2 * (i + 1)

	return
}

func parentIndex(i int) (pIndex int) {
	return ((i + 2) / 2) - 1
}
