package t00109

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	bst := sortedListToBST(newTree([]int{-10, -3, 0, 5, 9}))
	_ = bst
}

func newTree(items []int) *ListNode {
	var root ListNode
	var curr = &root

	for _, item := range items {
		var node = new(ListNode)
		node.Val = item

		curr.Next = node
		curr = node
	}

	return root.Next
}

func TestM(t *testing.T) {
	var r = (1 << 31) - 1
	fmt.Printf("x=%d\n", r)
	fmt.Printf("x=%b\n", r)
}
