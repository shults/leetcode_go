package t00055

func canJump(nums []int) bool {
	st := newStack(len(nums))
	st.push(0)
	dest := len(nums) - 1

	for {
		if st.isEmpty() {
			return false
		}

		currIndex := st.pop()

		if currIndex == dest {
			return true
		}

		currValue := nums[currIndex]

		for i, upto := currIndex, min(currIndex+currValue, len(nums)-1); i < upto; upto-- {
			if !st.push(upto) {
				break
			}
		}
	}
}

type stack struct {
	items    []int
	set      []bool
	capacity int
	length   int
}

func newStack(size int) *stack {
	return &stack{
		capacity: size,
		length:   0,
		items:    make([]int, size),
		set:      make([]bool, size),
	}
}

func (s *stack) push(index int) bool {
	if s.set[index] {
		return false
	}

	s.items[s.length] = index
	s.length++
	s.set[index] = true

	return true
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic("non expected tu call")
	}

	s.length--
	return s.items[s.length]
}

func (s *stack) isEmpty() bool {
	return s.length == 0
}
