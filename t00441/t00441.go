package t00441

func arrangeCoins(n int) int {
	return arrangeCoinsAux(n, 1, 1<<16)
}

func arrangeCoinsAux(expected, start, end int) int {
	if end-start <= 1 {
		return start
	}

	var mid = (start + end) / 2
	var sumOf = sum(mid)

	if expected == sumOf {
		return mid
	} else if sumOf < expected {
		return arrangeCoinsAux(expected, mid, end)
	} else {
		return arrangeCoinsAux(expected, start, mid)
	}
}

func sum(n int) int {
	return ((n + 1) * n) / 2
}
