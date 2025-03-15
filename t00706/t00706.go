package t00706

const defaultCapacity = 4
const bucketSize = 8

type MyHashMap struct {
	buckets  []*bucket
	capacity int
	size     int
}

type bucket struct {
	size       int
	key, value int
	//items [bucketSize * 2]int
	next *bucket
}

func (b *bucket) put(key, value int) {
	// todo: rewrite to bucket approach
	// return
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets:  make([]*bucket, defaultCapacity),
		capacity: defaultCapacity,
		size:     0,
	}
}

func (m *MyHashMap) Put(key int, value int) {
	m.grow()

	var index = getIndex(key, m.capacity)
	var entry = m.buckets[index]

	if entry == nil {
		m.size++
		m.buckets[index] = &bucket{
			key:   key,
			value: value,
			next:  nil,
		}
	} else {
		for entry != nil {
			if entry.key == key {
				entry.value = value
				return
			}
			entry = entry.next
		}

		m.buckets[index] = &bucket{
			key:   key,
			value: value,
			next:  m.buckets[index],
		}
		m.size++
	}
}

func (m *MyHashMap) Get(key int) int {
	var index = getIndex(key, m.capacity)
	var entry = m.buckets[index]

	for entry != nil {
		if entry.key == key {
			return entry.value
		}

		entry = entry.next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	var index = getIndex(key, m.capacity)
	var entry = m.buckets[index]

	var prev *bucket
	for entry != nil {
		if entry.key == key {
			if prev == nil {
				m.buckets[index] = entry.next
			} else {
				prev.next = entry.next
			}

			m.size--
			return
		}

		prev = entry
		entry = entry.next
	}
}

func (m *MyHashMap) grow() {
	// todo: remove magic numbers
	if m.size < m.capacity {
		return
	}

	var newCapacity = m.capacity * 2
	var newItems = make([]*bucket, newCapacity)

	for _, item := range m.buckets {
		for item != nil {

			var index = getIndex(item.key, newCapacity)

			var entry = bucket{
				key:   item.key,
				value: item.value,
				next:  newItems[index],
			}

			newItems[index] = &entry
			item = item.next
		}
	}

	m.buckets = newItems
	m.capacity = newCapacity
}

func getIndex(key int, capacity int) int {
	return hash(key) % capacity
}

func hash(val int) int {
	return val
}
