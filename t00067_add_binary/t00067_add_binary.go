package t00067_add_binary

const (
	ZERO = "0"
	ONE  = "1"
)

func addBinary(a string, b string) string {
	maxLen := max(len(a), len(b))
	res := ""
	rest := 0

	for i := maxLen; i >= 0; i-- {
		sum := getVal(a, i-1, maxLen) + getVal(b, i-1, maxLen) + rest
		val := 0

		switch sum {
		case 0:
			val = 0
			rest = 0
		case 1:
			val = 1
			rest = 0
		case 2:
			val = 0
			rest = 1
		case 3:
			val = 1
			rest = 1
		}

		if val == 0 {
			res = ZERO + res
		} else {
			res = ONE + res
		}
	}

	if len(res) > maxLen && res[0:1] == "0" {
		return res[1:]
	} else {
		return res
	}
}

func getVal(src string, i int, maxLen int) int {
	i -= maxLen - len(src)

	if i >= len(src) || i < 0 {
		return 0
	}

	return int(src[i]) - int('0')
}
