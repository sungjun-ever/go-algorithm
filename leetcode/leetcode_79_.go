package main

func exist(board [][]byte, word string) bool {
	// 셀의 문자는 한 번만 사용 되어야함
	// word는 인접한 셀과 연속 되어야함
	// 셀 확인 여부를 저장할 그리드를 하나 만든다
	m := len(board)
	n := len(board[0])

	wLen := len(word)

	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		if index == wLen {
			return true
		}

		// 범위 밖으로 나가면, 글자가 일치하지 않으면 종료
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[index] {
			return false
		}

		// 방문 시에 임시로 글자를 바꿔줌
		temp := board[x][y]
		board[x][y] = '.'

		isFound := dfs(x+1, y, index+1) ||
			dfs(x-1, y, index+1) ||
			dfs(x, y+1, index+1) ||
			dfs(x, y-1, index+1)

		// 복구
		board[x][y] = temp

		return isFound
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false

}
