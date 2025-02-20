package common

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
