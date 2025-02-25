package t00219

func containsNearbyDuplicate(nums []int, k int) bool {
	var w = newWindow()

	for i := 0; i < len(nums); i++ {
		var val = nums[i]

		if w.size == k+1 {
			w.remove(nums[i-k-1])
		}

		if w.add(val) {
			return true
		}
	}

	return false
}

type window struct {
	size  int
	items map[int]int
}

func newWindow() *window {
	return &window{
		size:  0,
		items: make(map[int]int),
	}
}

func (w *window) add(val int) bool {
	w.size++
	w.items[val] = w.items[val] + 1
	return w.items[val] > 1
}

func (w *window) remove(val int) {
	w.size--
	w.items[val] = w.items[val] - 1
}
