package main

func solve(board [][]byte) {
	// 2차원 행렬 board는 X, O을 가지고 있다
	// X로 둘러싸인 O은 X로 변경된다, 벽 밖은 X로 취급하지 않는다
	// 둘러싸인 정의는 현재 셀 기준으로 4방향이 모두 X인 경우

	// 벽에 있는 O가 아닌 O들은 X로 바뀐다
	// 벽과 연결된 0이 아닌 0을 엑스로 바꿔준다
	// 먼저 벽에있는 0과 연결된 0들을 다른 문자로 바꿔주고
	// 다른 문자가 아닌건 X로 다른 문자는 다시 O로 바꾼다

	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		board[x][y] = 'C'

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx >= 0 && ny >= 0 && nx < len(board) && ny < len(board[0]) && board[nx][ny] == 'O' {
				dfs(nx, ny)
			}
		}
	}

	// 경계에 있는 O들만 C로 바꿔준다
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 경계에 있다면
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1) && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'C' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}

}
