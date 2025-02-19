package t00063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	dp := make([][]int, len(obstacleGrid))

	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}

	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}

		dp[i][0] = 1
	}

	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}

		dp[0][j] = 1
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	a, b := len(obstacleGrid)-1, len(obstacleGrid[0])-1

	return dp[a][b]
}
