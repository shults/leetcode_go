package t00091

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var dp = make([]int, len(s))

	if s[0] >= '1' && s[0] <= '9' {
		dp[0] = 1
	} else {
		return 0
	}

	for i := 1; i < len(s); i++ {
		var ch = s[i]
		var chPrev = s[i-1]

		var eventualPrev = int(chPrev-'0') * 10
		eventualPrev += int(ch - '0')

		dp[i] = dp[i-1] + 1

		if eventualPrev <= 26 && (i-2) >= 0 {
			dp[i] += dp[i-2] + 1
		}
	}

	// 1 1 0
	return dp[len(s)-1]
}
