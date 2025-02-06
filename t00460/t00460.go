package t00460

// LFUCache cache
type LFUCache struct {
	capacity int
	data     map[int]cacheEntry
	root     *mainList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		data:     make(map[int]cacheEntry),
		root:     newMainList(),
	}
}

func (r *LFUCache) Get(key int) int {
	rNode, ok := r.data[key]

	if !ok {
		return -1
	}

	r.riseNode(rNode)

	return rNode.value
}

func (r *LFUCache) Put(key int, value int) {
	entry, ok := r.data[key]
	entry.key = key
	entry.value = value

	if ok {
		r.riseNode(entry)
		return
	}

	r.evictLeastUsedIfLimitExceeded()

	list := r.root.getFirst()
	newNode := list.val.prepend(entry.key)

	entry.n = newNode
	entry.listNode = list

	r.data[key] = entry
}

func (r *LFUCache) riseNode(entry cacheEntry) {
	lst := entry.listNode.val
	n := entry.n

	lst.removeNode(n)
	newList := r.root.getOrCreateNext(entry.listNode)
	entry.listNode = newList
	entry.n = newList.val.prepend(entry.key)

	if lst.isEmpty() {
		r.root.removeNode(entry.listNode)
		lst = nil
	}

	r.data[entry.key] = entry
}

func (r *LFUCache) evictLeastUsedIfLimitExceeded() {
	if len(r.data) < r.capacity {
		return
	}

	key := r.root.evictNode()
	delete(r.data, key)
}

type node[T any] struct {
	val  T
	prev *node[T]
	next *node[T]
}

type linkedList[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func (r *linkedList[T]) isEmpty() bool {
	return r.size == 0
}

func (r *linkedList[T]) prepend(val T) *node[T] {
	newNode := &node[T]{val: val}
	r.size++

	if r.head == nil {
		r.head = newNode
		r.tail = newNode
		return newNode
	}

	newNode.prev = nil
	newNode.next = r.head
	r.head.prev = newNode
	r.head = newNode

	return newNode
}

func (r *linkedList[T]) append(val T) *node[T] {
	newNode := &node[T]{val: val}
	r.size++

	if r.head == nil {
		r.head = newNode
		r.tail = newNode
		return newNode
	}

	newNode.prev = r.tail
	newNode.next = nil
	r.tail.next = newNode
	r.tail = newNode

	return newNode
}

func (r *linkedList[T]) removeLast() *node[T] {
	rNode := r.tail
	r.removeNode(rNode)
	return rNode
}

func (r *linkedList[T]) removeNode(n *node[T]) {
	switch {
	case r.size == 1:
		r.head = nil
		r.tail = nil

	case r.size == 2 && n.prev == nil:
		r.head = r.tail
		r.tail.next = nil
		r.tail.prev = nil

	case r.size == 2 && n.next == nil:
		r.tail = r.head
		r.head.next = nil
		r.head.prev = nil

	case n.prev == nil:
		r.head = r.head.next
		r.head.prev = nil

	case n.next == nil:
		r.tail = r.tail.prev
		r.tail.next = nil
	}

	r.size--
}

func (r *linkedList[T]) insertAfter(n *node[T], val T) *node[T] {
	newNode := &node[T]{
		val:  val,
		prev: nil,
		next: nil,
	}

	switch {
	case r.size == 0:
		panic("broken invariant")
	case r.size == 1:
		r.tail = newNode
		r.tail.prev = r.head
		r.head.next = r.tail
	case n.next == nil:
		n.next = newNode
		newNode.prev = n
		r.tail = newNode
	default:
		newNode.prev = n
		newNode.next = n.next
		n.next.prev = newNode
		n.next = newNode
	}

	r.size++

	return newNode
}

type cacheEntry struct {
	key      int
	value    int
	n        *node[int]
	listNode *node[*entryList]
}

type entryList struct {
	*linkedList[int]
	rank int
}

func (l *entryList) takeLast() {

}

func newEntryList(rank int) *entryList {
	return &entryList{
		rank:       rank,
		linkedList: &linkedList[int]{},
	}
}

type mainList struct {
	*linkedList[*entryList]
}

func newMainList() *mainList {
	return &mainList{
		linkedList: &linkedList[*entryList]{},
	}
}

func (r *mainList) getOrCreateNext(n *node[*entryList]) *node[*entryList] {
	switch {
	case n.next != nil && n.next.val.rank == n.val.rank+1:
		return n.next

	case n.next != nil && n.next.val.rank != n.val.rank+1:
		return n.next

	case n.next == nil:
		newNode := &node[*entryList]{
			val: newEntryList(n.val.rank + 1),
		}
		newNode.prev = n
		n.next = newNode
		r.tail = newNode
		r.size++

		return newNode
	default:
		panic("unexpected case")
	}
}

func (r *mainList) getFirst() *node[*entryList] {
	switch {
	case r.head == nil:
		return r.prepend(newEntryList(1))
	case r.head != nil && r.head.val.rank == 1:
		return r.head
	case r.head != nil && r.head.val.rank != 1:
		return r.prepend(newEntryList(1))
	default:
		panic("unexpected case")
	}
}

func (r *mainList) evictNode() (key int) {
	key = r.head.val.tail.val

	r.head.val.removeNode(r.head.val.tail)

	if r.head.val.isEmpty() {
		r.removeNode(r.head)
	}

	return key
}
