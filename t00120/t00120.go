package t00120

func minimumTotal(triangle [][]int) int {
	var dpPrev = make([]int, len(triangle))
	var dp = make([]int, len(triangle))

	for _, row := range triangle {
		for i := 0; i < len(row); i++ {
			if i == 0 {
				dp[i] = row[i] + dpPrev[i]
			} else if i == len(row)-1 {
				dp[i] = row[i] + dpPrev[i-1]
			} else {
				dp[i] = row[i] + min(dpPrev[i-1], dpPrev[i])
			}
		}

		copy(dpPrev, dp)
	}

	var res = dp[0]

	for i := 1; i < len(dp); i++ {
		res = min(res, dp[i])
	}

	return res
}
