package common

import (
	"cmp"
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap(func(a int, b int) int {
		return cmp.Compare(a, b)
	})

	//var origin = make([]int, 0)
	//var deheaped = make([]int, 0)

	for i := 0; i < 1000; i++ {
		var val = rand.Int() % 100
		//origin = append(origin, val)
		heap.Push(val)
		//fmt.Printf("push: %v\n", heap.items[:heap.size])
	}

	for !heap.Empty() {
		//deheaped = append(deheaped, heap.Pop())
		fmt.Printf("%d\n", heap.Pop())
		//fmt.Printf(" pop: %v\n", heap.items[:heap.size])
	}

	//fmt.Printf("origin=%+v\n", origin)
	//fmt.Printf("deheaped=%+v\n", deheaped)
}
