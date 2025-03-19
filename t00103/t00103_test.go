package t00103

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var root = node(3)
	root.Left = node(9)
	root.Right = node(20)
	root.Right.Left = node(15)
	root.Right.Right = node(7)

	var res = zigzagLevelOrder(root)
	fmt.Printf("result=%+v\n", res)
}

func node(val int) *TreeNode {
	var n = new(TreeNode)
	n.Val = val
	return n
}
