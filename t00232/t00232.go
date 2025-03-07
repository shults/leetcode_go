package t00232

import "errors"

type MyQueue struct {
	tail *Stack[int]
	head *Stack[int]
}

func Constructor() MyQueue {
	return MyQueue{
		head: NewStack[int](),
		tail: NewStack[int](),
	}
}

func (r *MyQueue) Push(x int) {
	r.tail.Push(x)
}

func (r *MyQueue) Pop() int {
	r.migrateToHead()
	return r.head.MustPop()
}

func (r *MyQueue) migrateToHead() {
	if r.head.IsEmpty() {
		for !r.tail.IsEmpty() {
			r.head.Push(r.tail.MustPop())
		}
	}
}

func (r *MyQueue) Peek() int {
	r.migrateToHead()
	return r.head.Peak()
}

func (r *MyQueue) Empty() bool {
	return r.head.IsEmpty() && r.tail.IsEmpty()
}

type Stack[T any] struct {
	length int
	last   *StackNode[T]
}

type StackNode[T any] struct {
	val  T
	prev *StackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peak() T {
	return s.last.val
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[T]) Pop() (T, error) {

	if s.IsEmpty() {
		var zero T
		return zero, errors.New("fuck")
	}

	s.length--
	item := s.last
	s.last = s.last.prev

	return item.val, nil
}

func (s *Stack[T]) MustPop() T {
	val, err := s.Pop()

	if err != nil {
		panic(err)
	}

	return val
}

func (s *Stack[T]) Push(val T) {
	last := &StackNode[T]{
		val:  val,
		prev: s.last,
	}

	s.last = last
	s.length++
}
