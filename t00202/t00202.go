package t00202

func isHappy(n int) bool {
	var cache = make(map[int]bool)
	return isHappyInternal(cache, n)
}

func isHappyInternal(cache map[int]bool, n int) bool {
	if res, ok := cache[n]; ok {
		return res
	}

	var visited = make(map[int]struct{})

retry:
	sum := 0
	visited[n] = struct{}{}
	for n > 0 {
		var rest = n % 10
		sum += rest * rest
		n /= 10
	}

	if sum == 1 {
		for k, _ := range visited {
			cache[k] = true
		}
		return true
	}

	if isHappy, ok := cache[sum]; ok {
		for k, _ := range visited {
			cache[k] = isHappy
		}
		return cache[sum]
	}

	if _, alreadyVisited := visited[sum]; alreadyVisited {
		for k, _ := range visited {
			cache[k] = false
		}
		return false
	}

	n = sum
	goto retry
}
