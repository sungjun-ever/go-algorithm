package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// m, n, k <= 100
// k개의 직사강형을 그릴때 직사각형을 제외한 빈 공간 덩어리 수 및 각 덩어리 넓이 구하기
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m, n, k int
	fmt.Fscan(reader, &m, &n, &k)

	grid := [100][100]int{}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for j := 0; j < k; j++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				grid[y][x] = 1
			}
		}
	}

	visited := [100][100]bool{}
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		visited[x][y] = true
		count := 1
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 && !visited[nx][ny] {
				count += dfs(nx, ny)
			}
		}

		return count
	}

	ans := make([]int, 0, n*m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				ans = append(ans, dfs(i, j))
			}
		}
	}

	sort.Ints(ans)
	fmt.Fprintln(writer, len(ans))

	for i := 0; i < len(ans); i++ {
		fmt.Fprint(writer, ans[i])
		if i < len(ans)-1 {
			fmt.Fprint(writer, " ")
		}
	}

}
