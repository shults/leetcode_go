package t00297

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var one = newNode(1)
	var two = newNode(2)
	var three = newNode(3)
	var four = newNode(4)
	var five = newNode(5)
	one.Right = three
	one.Left = two
	three.Left = four
	three.Right = five

	var codec = Constructor()
	var ser = codec.serialize(one)
	fmt.Printf("res=%+v\n", ser)
	var des = codec.deserialize(ser)
	fmt.Printf("res=%+v\n", codec.serialize(des))
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
