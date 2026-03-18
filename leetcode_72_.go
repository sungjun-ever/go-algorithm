package main

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// 문자열 2개를 2차원 배열로 생각
	// 빈 문자열을 하나 추가해준다
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// word1 또는 word2가 비어져있다면. 지우거나 채워야한다
	// N번째 글자에는 N개의 행동이 생긴다
	for x := 0; x <= m; x++ {
		dp[x][0] = x
	}

	for y := 0; y <= n; y++ {
		dp[0][y] = y
	}

	// 넣는 경우: dp[i][j] = dp[i][j-1]+1
	// 바꾸는 경우: dp[i][j] = dp[i-1][j-1]+1
	// 삭제하는 경우: dp[i][j] = dp[i-1][j] + 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j-1], dp[i-1][j]) + 1
			}
		}
	}

	return dp[m][n]

}
