package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향으로 하나라도 물에 잠기지 않은 영역이 존재하면 안전한 영역

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	if _, err := fmt.Fscan(reader, &n); err != nil {
		fmt.Fprintln(writer, err)
	}

	grid := make([][]int, n)
	minH, maxH := 1, 1

	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &grid[i][j])
			minH = min(minH, grid[i][j])
			maxH = max(maxH, grid[i][j])
		}
	}

	// 최소 높이에서 최대 높이 사이에서
	// 높이 별 안정한 영역이 계산되고
	// 그 중 안전한 영역의 최대 개수
	var visited [100][100]bool

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(int, int, int)

	dfs = func(r, c, h int) {
		visited[r][c] = true

		for _, d := range dirs {
			nx, ny := r+d[0], c+d[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < n && !visited[nx][ny] && grid[nx][ny] >= h {
				dfs(nx, ny, h)
			}
		}
	}

	maxArea := 0
	for h := minH; h <= maxH; h++ {
		area := 0
		visited = [100][100]bool{}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !visited[i][j] && grid[i][j] >= h {
					area++
					dfs(i, j, h)
				}
			}
		}
		maxArea = max(maxArea, area)
	}

	fmt.Fprintln(writer, maxArea)
}
