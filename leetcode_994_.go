package main

func orangesRotting(grid [][]int) int {
	// 썩은 오렌지를 큐에 넣는다
	// 큐에 있는 썩은 오렌지들은 동시에 전파 시킨 후 1분 증가 시켜야함
	// 이동 시키면서 신선한 오렌지면 썩게하고 신선한 오렌지 카운터 감소
	// 썩은 오렌지 다시 큐에 넣어줌
	// 하나라도 전염되면 시간 증가

	rows := len(grid)
	cols := len(grid[0])
	minute := 0     // 걸린 시간
	freshCount := 0 // 신선한 오렌지 수

	type Point struct{ r, c int }

	queue := make([]Point, 0, rows*cols)

	// 썩은 오렌지는 큐에 넣어줌
	// 신선하면 카운트
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, Point{i, j})
			}

			if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	// 신선한 오렌지가 없으면 종료
	if freshCount == 0 {
		return 0
	}

	// 큐를 순회하며 오염전파
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		size := len(queue)
		becomeRotten := false

		// 큐에 있는 썩은 오렌지를 같은 1분안에 전파 시켜야함
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nr, nc := curr.r+d[0], curr.c+d[1]

				// 범위 안에 있고 신선한 오렌지이면
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					freshCount--
					queue = append(queue, Point{nr, nc})
					becomeRotten = true
				}
			}
		}

		// 하나라도 상했으면 1분 올려줌
		if becomeRotten {
			minute++
		}

	}

	if freshCount > 0 {
		return -1
	}

	return minute
}
