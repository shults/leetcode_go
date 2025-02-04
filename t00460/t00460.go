package t00460

// LFUCache cache
type LFUCache struct {
	capacity int
	data     map[int]*node
	head     *node // most freq used
	tail     *node // least freq used
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		data:     make(map[int]*node, capacity),
		tail:     nil,
		head:     nil,
	}
}

func (r *LFUCache) Get(key int) int {
	rNode, ok := r.data[key]

	if !ok {
		return -1
	}

	rNode.ctr++
	r.riseNode(rNode)

	return rNode.val
}

func (r *LFUCache) Put(key int, value int) {
	rNode, ok := r.data[key]

	if ok {
		rNode.val = value
		rNode.ctr++
		r.riseNode(rNode)
		return
	}

	r.evictLeastUsedIfLimitExceeded()

	rNode = &node{
		ctr:  1,
		next: nil,
		prev: nil,
		key:  key,
		val:  value,
	}

	switch len(r.data) {
	case 0:
		r.head = rNode
		r.tail = rNode
	case 1:
		r.tail = rNode
		r.tail.prev = r.head
		r.head.next = r.tail
	default:
		rNode.prev = r.tail
		r.tail.next = rNode
		r.tail = rNode
	}

	r.data[rNode.key] = rNode
	r.riseNode(rNode)
}

func (r *LFUCache) evictLeastUsedIfLimitExceeded() {
	if len(r.data) < r.capacity {
		return
	}

	if r.tail == nil {
		return
	}

	delete(r.data, r.tail.key)

	if r.tail.prev == nil {
		r.tail = nil
		r.head = nil
		return
	}

	r.tail = r.tail.prev
	r.tail.next = nil
}

func (r *LFUCache) riseNode(n *node) {

	atLeastOneIterationMade := false

	for {
		// no place move forward
		if n.prev == nil {
			r.head = n
			return
		}

		// we are in proper place
		if n.ctr < n.prev.ctr {
			return
		}

		if atLeastOneIterationMade && n.ctr == n.prev.ctr {
			return
		}

		atLeastOneIterationMade = true

		switch {
		case len(r.data) == 1:
			return
		case len(r.data) == 2:
			r.head, r.tail = r.tail, r.head
			r.head.prev = nil
			r.head.next = r.tail
			r.tail.next = nil
			r.tail.prev = r.head
		case n == r.tail:
			preDown := r.tail.prev.prev
			down := r.tail.prev
			up := r.tail

			preDown.next = up
			up.prev = preDown
			up.next = down
			down.next = nil
			down.prev = up

			n = up
			r.tail = down
		default:

			top := n.prev.prev
			down := n.prev
			up := n
			bottom := n.next

			if top != nil {
				top.next = up
			}

			up.prev = top
			up.next = down
			down.prev = up
			down.next = bottom

			if bottom != nil {
				bottom.prev = down
			}
		}
	}

}

type node struct {
	val  int
	key  int
	ctr  int
	prev *node
	next *node
}
