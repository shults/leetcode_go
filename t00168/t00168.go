package t00168

func convertToTitle(columnNumber int) string {
	var res string

	const base = 26

	for columnNumber > 0 {
		columnNumber -= 1
		rest := columnNumber % base
		res = string([]byte{'A' + byte(rest)}) + res
		columnNumber /= base
	}

	return res
}
