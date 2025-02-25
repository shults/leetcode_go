package t00367

func isPerfectSquare(num int) bool {
	var x = 1 << 15
	var top = int(float64(x) * 1.42)

	var r = sqrt2(num, 0, top)
	r = r * r

	return r == num
}

func sqrt2(num int, start, end int) int {
	if end-start <= 1 {
		return start
	}

	var mid = (start + end) / 2

	if (mid * mid) > num {
		return sqrt2(num, start, mid)
	} else {
		return sqrt2(num, mid, end)
	}
}
