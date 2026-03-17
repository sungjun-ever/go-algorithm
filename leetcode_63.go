package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[i] = 0
		} else {
			if i > 0 {
				dp[i] = dp[i] + dp[i-1]
			} else {
				dp[i] = 1
			}
		}
	}

	// 첫 줄 먼저 초기화
	// 장애물이 있으면 0, 아니면1
	for x := 1; x < m; x++ {
		for y := 0; y < n; y++ {
			if obstacleGrid[x][y] == 1 {
				dp[y] = 0
			} else {
				if y > 0 {
					dp[y] = dp[y] + dp[y-1]
				}
			}
		}

	}

	return dp[n-1]
}
