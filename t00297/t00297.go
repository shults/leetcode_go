package t00297

import (
	"fmt"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (r *Codec) serialize(root *TreeNode) string {
	var builder strings.Builder
	builder.WriteByte('[')
	serialize(maxDepth(root), root, &builder, true)
	builder.WriteByte(']')
	return builder.String()
}

// Deserializes your encoded data to tree.
func (r *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}

	var root TreeNode
	var iter = NewIterator(data)

	deserialize(&root, iter)

	return &root
}

func serialize(depth int, node *TreeNode, builder *strings.Builder, isFirst bool) {
	if depth == 0 {
		return
	}

	if !isFirst {
		builder.WriteString(",")
	}

	if node == nil {
		builder.WriteString("null")
	} else {
		builder.WriteString(fmt.Sprintf("%d", node.Val))
	}

	depth--
	serialize(depth, left(node), builder, false)
	serialize(depth, right(node), builder, false)
}

func left(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}

	return n.Left
}

func right(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}

	return n.Right
}

func deserialize(dest *TreeNode, iterator *Iterator) {
	val, active := iterator.Next()

	if val != nil {
		fmt.Printf("val=%v, active=%+v\n", *val, active)
	} else {
		fmt.Printf("val=%v, active=%+v\n", val, active)
	}

	if !active {
		return
	}

	if dest != nil && val != nil {
		(*dest).Val = *val
	}

	if dest != nil {
		deserialize(dest.Left, iterator)
		deserialize(dest.Right, iterator)
	} else {
		deserialize(nil, iterator)
		deserialize(nil, iterator)
	}
}

func maxDepth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	return 1 + max(maxDepth(n.Left), maxDepth(n.Right))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Iterator struct {
	data string
	pos  int
}

func NewIterator(data string) *Iterator {
	return &Iterator{
		data: data,
		pos:  0,
	}
}

func (r *Iterator) Next() (*int, bool) {
	if r.pos >= len(r.data) {
		return nil, false
	}

	var ch = r.data[r.pos]

	switch {
	case ch == ']':
		r.pos++
		return nil, false
	case ch == '[' || ch == ',':
		r.pos++
		return r.Next()
	case ch == 'n':
		var chars = "null"

		for i := 0; i < 4; i++ {
			if chars[i] != r.data[r.pos+i] {
				panic("broken format")
			}
		}
		r.pos += 4
		return nil, true
	case isDigit(ch):
		var number int

		for isDigit(ch) {
			number *= 10
			number += int(ch - '0')
			r.pos++
			ch = r.data[r.pos]
		}

		return &number, true
	default:
		panic("unexpected invariant")
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
