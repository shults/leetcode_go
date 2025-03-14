package t01092

import "strings"

func shortestCommonSupersequence(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	} else if len(str1) > len(str2) && strings.Contains(str1, str2) {
		return str1
	} else if len(str1) < len(str2) && strings.Contains(str2, str1) {
		return str2
	}

	// var a = concatMin(str1, str2)
	// var b = concatMin(str2, str1)
	// if len(a) > len(b) {
	// 	return b
	// } else {
	// 	return a
	// }

	var a = concatMin2(str1, str2)
	var b = concatMin2(str2, str1)

	if len(a) > len(b) {
		return b
	} else {
		return a
	}

	// return concatMin2(str1, str2)
}

func concatMin2(a, b string) string {
	var maxOverlap = 0

	for i := 0; i < len(a); i++ {

		var ai = i
		var bi = 0

		for {
			if ai == len(a) || bi == len(b) {
				break
			}

			if a[ai] == b[bi] {
				ai++
				bi++
			} else {
				ai++
			}
		}

		maxOverlap = max(maxOverlap, bi)
	}

	return a + b[maxOverlap:]
}

func concatMin(prefix, suffix string) string {
	var maxOverlap = min(len(prefix), len(suffix)) - 1

	for maxOverlap > 0 {
		var a = prefix[len(prefix)-maxOverlap:]
		var b = suffix[:maxOverlap]

		if a == b {
			return prefix[:len(prefix)-maxOverlap] + suffix
		}

		maxOverlap--
	}

	return prefix + suffix
}
