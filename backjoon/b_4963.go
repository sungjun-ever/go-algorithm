package main

import (
	"bufio"
	"fmt"
	"os"
)

// 가로, 세로, 대각선으로 연결되어 있으면 같은 섬으로 취급
// 0 <= w, h <= 50
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	w, h := -1, -1
	// 4방향 + 대각선 4방향
	dirs := [8][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for true {
		fmt.Fscan(reader, &w, &h)

		if w == 0 && h == 0 {
			break
		}

		grid := make([][]int, h)
		for i := 0; i < h; i++ {
			grid[i] = make([]int, w)
			for j := 0; j < w; j++ {
				fmt.Fscan(reader, &grid[i][j])
			}
		}

		var dfs func(int, int)
		dfs = func(x, y int) {
			// 땅을 바다로 바꿔서 방문 처리
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]

				// 범위 안에 땅인경우에만
				if nx >= 0 && nx < h && ny >= 0 && ny < w && grid[nx][ny] == 1 {
					grid[nx][ny] = 0
					dfs(nx, ny)
				}
			}
		}

		islandCnt := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if grid[i][j] == 1 {
					grid[i][j] = 0
					islandCnt++
					dfs(i, j)
				}
			}
		}
		fmt.Fprintln(writer, islandCnt)
	}
}
