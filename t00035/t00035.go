package t00035

func searchInsert(nums []int, target int) int {
	return searchInsertInternal(nums, target, 0, len(nums))
}

func searchInsertInternal(nums []int, target int, start, end int) int {
	mid := (end + start) / 2

	if end-start == 1 {
		val := nums[mid]

		if val > target {
			return mid
		} else {
			return mid + 1
		}
	}

	if target > nums[mid] {
		return searchInsertInternal(nums, target, mid, end)
	} else {
		return searchInsertInternal(nums, target, start, mid)
	}
}
