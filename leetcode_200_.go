package main

func numIslands(grid [][]byte) int {
	// 먼저 땅을 찾음
	// 땅의 4방향을 탐색하며 1인경우 0으로 바꿔줌
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		// 범위 나가면 종료
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return
		}

		// 0으로 변경
		grid[r][c] = '0'

		// 상하좌우 탐색
		dfs(r, c-1)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r+1, c)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}
