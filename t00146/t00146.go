package t00146

type LRUCache struct {
	capacity int
	data     map[int]*node
	head     *node
	tail     *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*node, capacity),
	}
}

func (r *LRUCache) Get(key int) int {
	rNode, ok := r.data[key]

	if !ok {
		return -1
	}

	if len(r.data) == 1 {
		goto ret
	}

	if r.head == rNode {
		goto ret
	}

	if r.tail == rNode {
		r.tail = r.tail.prev
		r.tail.next = nil

		rNode.prev = nil
		rNode.next = r.head
		r.head.prev = rNode
		r.head = rNode

		goto ret
	}

	rNode.next.prev = rNode.prev
	rNode.prev.next = rNode.next
	rNode.prev = nil
	rNode.next = r.head
	r.head.prev = rNode
	r.head = rNode
ret:
	return rNode.val
}

func (r *LRUCache) Put(key int, value int) {
	rNode, ok := r.data[key]

	if ok {
		// key already exist wy should only update its value and put it in front of queue
		rNode.val = value
		_ = r.Get(key) // put key into head of queue
		return
	}

	r.evictLastIfIsFull()

	nItem := &node{
		key:  key,
		val:  value,
		next: nil,
		prev: nil,
	}

	r.data[key] = nItem

	if r.head == nil {
		r.head = nItem
		r.tail = nItem
	} else {
		nItem.next = r.head
		r.head.prev = nItem
		r.head = nItem
	}
}

// evictLastIfIsFull evicts last entry
func (r *LRUCache) evictLastIfIsFull() {
	if r.capacity != len(r.data) {
		return
	}

	evictNode := r.tail
	delete(r.data, evictNode.key)

	switch r.capacity {
	case 1:
		r.head = nil
		r.tail = nil
	case 2:
		r.head.next = nil
		r.head.prev = nil
		r.tail = r.head
	default:
		r.tail = evictNode.prev
		r.tail.next = nil
	}
}

type node struct {
	key  int
	val  int
	prev *node
	next *node
}
