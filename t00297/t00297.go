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
	type entry struct {
		node  *TreeNode
		depth int
	}

	var stopDepth = maxDepth(root)
	var builder strings.Builder
	var queue = NewQueue[entry]()

	queue.Enqueue(entry{
		depth: 0,
		node:  root,
	})

	builder.WriteByte('[')

	var isFirst = true

	for !queue.IsEmpty() {
		entryInstance, err := queue.Dequeue()

		if err != nil {
			panic(err)
		}

		if entryInstance.depth == stopDepth {
			break
		}

		if !isFirst {
			builder.WriteString(",")
		}

		isFirst = false

		if entryInstance.node == nil {
			builder.WriteString("null")
		} else {
			builder.WriteString(fmt.Sprintf("%d", entryInstance.node.Val))
		}

		queue.Enqueue(entry{
			depth: entryInstance.depth + 1,
			node:  left(entryInstance.node),
		})

		queue.Enqueue(entry{
			depth: entryInstance.depth + 1,
			node:  right(entryInstance.node),
		})
	}

	builder.WriteByte(']')
	return builder.String()
}

// Deserializes your encoded data to tree.
func (r *Codec) deserialize(data string) *TreeNode {

	iter := NewIterator(data)

	first, active := iter.Next()

	if !active {
		return nil
	}

	root := new(TreeNode)
	root.Val = *first

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		parent, _ := queue.Dequeue()

		// left
		leftValue, isActive := iter.Next()

		if !isActive {
			break
		}

		var leftNode *TreeNode

		if parent != nil && leftValue != nil {
			parent.Left = new(TreeNode)
			parent.Left.Val = *leftValue
			leftNode = parent.Left
		}

		queue.Enqueue(leftNode)

		// right
		rightValue, isActive := iter.Next()

		if !isActive {
			break
		}

		var rightNode *TreeNode

		if parent != nil && rightValue != nil {
			parent.Right = new(TreeNode)
			parent.Right.Val = *rightValue
			rightNode = parent.Right
		}

		queue.Enqueue(rightNode)
	}

	return root
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
	case isDigit(ch) || ch == '-':
		var number int
		var isNegative = ch == '-'

		if isNegative {
			r.pos++
			ch = r.data[r.pos]
		}

		for isDigit(ch) {
			number *= 10
			number += int(ch - '0')
			r.pos++
			ch = r.data[r.pos]
		}

		if isNegative {
			number *= -1
		}

		return &number, true
	default:
		panic(fmt.Sprintf("unexpected invariant char=%s", string(ch)))
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

type (
	Queue[T any] struct {
		length int
		head   *queueNode[T]
		tail   *queueNode[T]
	}

	queueNode[T any] struct {
		val  T
		prev *queueNode[T]
		next *queueNode[T]
	}

	qError struct {
		msg string
	}
)

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) Len() int {
	return q.length
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) Enqueue(val T) {
	node := new(queueNode[T])
	node.val = val

	switch q.length {
	case 0:
		q.head = node
		q.tail = node
	case 1:
		q.tail = node
		q.head.next = q.tail
		q.tail.prev = q.head
	default:
		q.tail.next = node
		node.prev = q.tail
		q.tail = node
	}

	q.length++
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.length == 0 {
		var zero T
		return zero, newQError("empty queue")
	}

	val := q.head.val

	switch q.length {
	case 1:
		q.head = nil
		q.tail = nil
	case 2:
		q.head = q.tail
		q.head.next = nil
		q.head.prev = nil
	default:
		q.head = q.head.next
		q.head.prev = nil
	}

	q.length--

	return val, nil
}

func newQError(msg string) qError {
	return qError{
		msg: msg,
	}
}

func (e qError) Error() string {
	return e.msg
}
