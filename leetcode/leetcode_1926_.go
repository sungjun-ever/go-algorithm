package main

// 최단거리 탐색은 bfs 이용
func nearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	cols := len(maze[0])

	type State struct {
		r, c, steps int
	}

	queue := []State{{entrance[0], entrance[1], 0}}
	maze[entrance[0]][entrance[1]] = '+'

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			// 현재 좌표에서 이동 좌표를 구함
			nr, nc := curr.r+d[0], curr.c+d[1]

			// 이동 좌표가 배열안에 있고 빈 공간인 경우
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && maze[nr][nc] == '.' {
				// 이동 좌표가 가장자리인 경우 나감
				if nr == 0 || nr == rows-1 || nc == 0 || nc == cols-1 {
					return curr.steps + 1
				}

				// 출구가 아니면 지나감 표시 및 큐에 넣음
				maze[nr][nc] = '+'
				queue = append(queue, State{nr, nc, curr.steps + 1})
			}

		}
	}

	return -1

}
