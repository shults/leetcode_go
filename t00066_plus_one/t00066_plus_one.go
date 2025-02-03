package t00066_plus_one

func plusOne(digits []int) []int {
	rest := 1

	for i := len(digits) - 1; i >= 0; i-- {
		if rest == 0 {
			break
		}

		digits[i] += rest

		if digits[i] >= 10 {
			rest = 1
			digits[i] -= 10
		} else {
			rest = 0
		}
	}

	if rest == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
