package t00143

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	lst := NewListNode(1, 2, 3, 4, 5)
	fmt.Printf("res=%+v\n", lst.ToSlice())
	reorderList(lst)
	fmt.Printf("res=%+v\n", lst.ToSlice())
}
