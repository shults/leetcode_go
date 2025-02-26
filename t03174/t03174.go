package t03174

func clearDigits(s string) string {
	var res = []byte(s)

	left := 0
	i := 0

	for {

		if len(s) == i {
			return string(res[:left])
		}

		if s[i] >= '0' && s[i] <= '9' {
			left = max(left-1, 0)
			i++
			continue
		}

		res[left] = s[i]

		left++
		i++
	}

}
