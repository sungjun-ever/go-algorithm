package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2 <= m, n <= 1000
// m은 가로 n은 세로
// 1은 익은 토마토, 0은 안익은 토마토, -1은 없는 칸
// 익은 토마토 주변에 안익은 토마토는 동시에 익는다
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m, n int
	fmt.Fscan(reader, &m, &n)

	grid := make([][]int, n)
	queue := make([][2]int, 0)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 모든 익은 토마토에서 동시에 출발함
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nx, ny := curr[0]+d[0], curr[1]+d[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == 0 {
				grid[nx][ny] = grid[curr[0]][curr[1]] + 1
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}

	maxDay := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 안 익은게 존재한다면
			if grid[i][j] == 0 {
				fmt.Fprintln(writer, -1)
				return
			} else {
				maxDay = max(maxDay, grid[i][j])
			}
		}
	}

	if maxDay == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, maxDay-1)
	}
}
