package main

func uniquePaths(m int, n int) int {
	// 로봇은  0 0에서 시작해, m-1 n-1 으로 가야함
	// 로봇은 한 번에 오른쪽 또는 아래로만 이동 가능
	// 도착하기 위한 모든 경우의 수 탐색

	// dp[x][y] = dp[x][y-1] + dp[x-1][y]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if x == 0 || y == 0 {
				dp[x][y] = 1
			} else {
				dp[x][y] = dp[x][y-1] + dp[x-1][y]
			}

		}
	}

	return dp[m-1][n-1]
}

// 첫 줄을 계속 갱신 시켜서 1차원 배열로 하는 방법
func uniquePathsInPlace(m int, n int) int {
	// 1차원 배열로 풀이하면
	dp := make([]int, n)

	// 첫 row를 1로 초기화
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	// 한 줄의 배열만 계속 갱신하고 마지막 값을 가져오는 접근
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			// 나의 경우의 수는 내 윗줄 + 내 왼쪽
			dp[y] = dp[y] + dp[y-1]
		}
	}

	return dp[n-1]
}
