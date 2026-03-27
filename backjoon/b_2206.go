package main

import (
	"bufio"
	"fmt"
	"os"
)

// 0은 이동, 1은 벽
// 1,1 -> n, m 최단경로
// 이동은 4방향 가능
// 벽을 최대 하나까지는 부수기 가능
type State struct {
	x      int
	y      int
	broken int
}

func main() {
	var (
		n       int
		m       int
		grid    [][]int
		queue   []State
		visited [1001][1001][2]int
		dirs    = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &m)
	grid = make([][]int, n)

	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		var row string
		fmt.Fscan(reader, &row)
		for j := 0; j < m; j++ {
			grid[i][j] = int(row[j] - '0')
		}
	}

	queue = append(queue, State{0, 0, 0})
	visited[0][0][0] = 1
	visited[0][0][1] = 1

	var bfs func() int
	bfs = func() int {
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nx, ny := curr.x+d[0], curr.y+d[1]

				if curr.x == n-1 && curr.y == m-1 {
					return visited[curr.x][curr.y][curr.broken]
				}

				if nx >= 0 && nx < n && ny >= 0 && ny < m {
					// 다음이 벽이고 아직 벽을 부순적 없을때
					if grid[nx][ny] == 1 && curr.broken == 0 {
						if visited[nx][ny][1] == 0 {
							visited[nx][ny][1] = visited[curr.x][curr.y][0] + 1
							queue = append(queue, State{nx, ny, 1})
						}
					}

					// 다음이 빈 공간이고 아직 방문한 적 없을때
					if grid[nx][ny] == 0 && visited[nx][ny][curr.broken] == 0 {
						visited[nx][ny][curr.broken] = visited[curr.x][curr.y][curr.broken] + 1
						queue = append(queue, State{nx, ny, curr.broken})
					}
				}
			}

		}

		return -1
	}

	fmt.Fprintln(writer, bfs())
}
