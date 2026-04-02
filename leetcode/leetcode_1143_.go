package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			if text1[x-1] == text2[y-1] {
				dp[x][y] = dp[x-1][y-1] + 1
			} else {
				dp[x][y] = max(dp[x-1][y], dp[x][y-1])
			}
		}
	}

	return dp[m][n]
}
