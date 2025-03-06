package common

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

func (s *Stack[T]) Empty() bool {
	return s.length == 0
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("stack is empty")
	}

	s.length--
	item := s.last
	s.last = s.last.prev

	return item.val
}

func (s *Stack[T]) Push(val T) {
	last := &StackNode[T]{
		val:  val,
		prev: s.last,
	}

	s.last = last
	s.length++
}

type PriorityStack[T any] struct {
	items *Stack[T]
}

func NewPriorityStack[T any]() *PriorityStack[T] {
	return &PriorityStack[T]{
		items: nil,
	}
}

func (s *PriorityStack[T]) Push(priority int, val T) {

}
