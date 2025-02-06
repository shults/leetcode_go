package t00028

func strStr(haystack string, needle string) int {
	match := 0
	i := 0

	for {
		if len(needle) == match {
			return i - match
		}

		if i == len(haystack) {
			break
		}

		if needle[match] == haystack[i] {
			match++
			i++
		} else {
			i -= match
			i++
			match = 0
		}
	}

	return -1
}
