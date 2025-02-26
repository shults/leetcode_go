package t00245

var vovels [256]bool

func init() {
	for _, ch := range []byte("aeouiAEOUI") {
		vovels[int(ch)] = true
	}
}

func reverseVowels(s string) string {
	var sb = []byte(s)
	var left = -1
	var right = len(s)

	nextVovel(&left, sb)
	prevVovel(&right, sb)

	for left < right {
		sb[left], sb[right] = sb[right], sb[left]
		nextVovel(&left, sb)
		prevVovel(&right, sb)
	}

	return string(sb)
}

func prevVovel(ptr *int, sb []byte) {
	for i := *ptr - 1; i > -1; i-- {
		*ptr = i

		if vovels[sb[i]] {
			break
		}
	}
}

func nextVovel(ptr *int, sb []byte) {
	for i := *ptr + 1; i < len(sb); i++ {
		*ptr = i

		if vovels[sb[i]] {
			break
		}
	}
}
