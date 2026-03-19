package main

func minDistance(word1 string, word2 string) int {
	/**
	  빈 문자열이라고 생각했을 때
	  글자수 만큼 행동이 발생해야함
	  3가지 상태가 있음
	  + 1은 행동 추가값
	  1. 삽입 I => dp[i][j] = dp[i][j-1] + 1
	  2. 삭제 D => dp[i][j] = dp[i-1][j] + 1
	  3. 바꾸기 R => dp[i][j] = dp[i-1][j-1] + 1
	  4. 같으면 dp[i][j] = dp[i-1][j-1]
	  5. 행동이 발생하면 세가지 상태 중 더 작은 수로

	      1   h   o   r   s   e
	  1   0   1   2   3   4   5
	  r   1
	  o   2
	  s   3
	*/

	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for x := 0; x <= m; x++ {
		dp[x][0] = x
	}

	for y := 0; y <= n; y++ {
		dp[0][y] = y
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}
