package t00297

import (
	"fmt"
	"strconv"
	"strings"
)

var null = []byte("null")

func newBuilder() *Builder {
	b := &Builder{
		numberOfNils: 0,
		isFirst:      true,
	}

	return b
}

type Builder struct {
	numberOfNils int
	sb           strings.Builder
	isFirst      bool
}

func (b *Builder) writeNil() {
	b.numberOfNils++
}

func (b *Builder) flushNils() {

	for i := 0; i < b.numberOfNils; i++ {
		if !b.isFirst {
			b.sb.WriteByte(',')
		}
		b.isFirst = false

		_, _ = b.sb.Write(null)
	}
	b.numberOfNils = 0
}

func (b *Builder) write(val int) {
	b.flushNils()

	if !b.isFirst {
		b.sb.WriteByte(',')
	}
	b.isFirst = false

	b.sb.WriteString(strconv.Itoa(val))
}

func (b *Builder) String() string {
	return fmt.Sprintf("[%s]", b.sb.String())
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (r *Codec) serialize(root *TreeNode) string {

	var queue = NewQueue[*TreeNode]()

	queue.Enqueue(root)
	var builder = newBuilder()

	for !queue.IsEmpty() {
		item, err := queue.Dequeue()

		if err != nil {
			panic(err)
		}

		if item == nil {
			builder.writeNil()
			continue
		} else {
			builder.write(item.Val)
		}

		queue.Enqueue(item.Left)
		queue.Enqueue(item.Right)
	}

	return builder.String()
}

// Deserializes your encoded data to tree.
func (r *Codec) deserialize(data string) *TreeNode {
	iter := NewParser(data)

	first, _, active := iter.Next()

	if !active {
		return nil
	}

	root := new(TreeNode)
	root.Val = first

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		parent, _ := queue.Dequeue()
		leftValue, isNil, isActive := iter.Next()

		if !isActive {
			break
		}

		if parent != nil && !isNil {
			parent.Left = new(TreeNode)
			parent.Left.Val = leftValue
			queue.Enqueue(parent.Left)
		}

		// right
		rightValue, isNil, isActive := iter.Next()

		if !isActive {
			break
		}

		if parent != nil && !isNil {
			parent.Right = new(TreeNode)
			parent.Right.Val = rightValue
			queue.Enqueue(parent.Right)
		}

	}

	return root
}

type Parser struct {
	data string
	pos  int
}

func NewParser(data string) *Parser {
	return &Parser{
		data: data,
		pos:  0,
	}
}

func (r *Parser) Next() (val int, isNil bool, isActive bool) {
	if r.pos >= len(r.data) {
		return 0, true, false
	}

	var ch = r.data[r.pos]

	switch {
	case ch == ']':
		r.pos++
		return 0, true, false
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
		return 0, true, true
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

		return number, false, true
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
