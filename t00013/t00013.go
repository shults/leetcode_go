package t00013

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

//Input: s = "MCMXCIV"
//Output: 1994

// MCMXCIV => 1994
// +1000 -100 +1000 -10 +100 -1 5
// MMXXIV

func val(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		panic("error")
	}
}

func romanToInt(s string) int {
	acc := 0
	res := 0

	for i := 0; i < len(s); i++ {
		curr := val(s[i])

		if i == 0 {
			acc = curr
			continue
		}

		prev := val(s[i-1])

		switch {
		case prev == curr:
			acc += curr
		case prev < curr:
			res += curr - acc
			acc = 0
		case prev > curr:
			res += acc
			acc = curr
		}
	}

	return res + acc
}
