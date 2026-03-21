package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 적록색약은 빨간색과 초록색 차이를 못느낀다
	// 색은 빨, 초, 파가 존재한다
	// 인접한 같은 색은 같은 구역으로 한다.
	// 빨초가 인접한 경우에도 같은 구역으로 한다

	// 적록색약인 사람과 아닌 사람이 봤을 때 구역 수

	reader := bufio.NewReader(os.Stdin)

	var n int // 줄 수
	fmt.Fscan(reader, &n)
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
		var paintLine string
		fmt.Fscan(reader, &paintLine)
		for j := 0; j < n; j++ {
			grid[i][j] = paintLine[j]
		}
	}

	// 4방향으로 퍼짐
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 방문 여부
	visited := [100][100]bool{}

	// 일반인 a, 색맹 b
	a, b := 0, 0

	// 구역을 나누려면
	// 이전꺼와 비교해서 같으면 같은 구역, 다르면 다른 구역
	var dfs func(int, int, byte, bool)
	dfs = func(x, y int, prev byte, isColor bool) {
		visited[x][y] = true

		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]

			// 범위 안이고 방문하지 않고 색이 같은 경우에만
			if nx >= 0 && nx < n && ny >= 0 && ny < n && !visited[nx][ny] && grid[nx][ny] == prev {
				dfs(nx, ny, grid[x][y], isColor)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				dfs(i, j, grid[i][j], true)
				a++
			}
		}
	}

	visited = [100][100]bool{}

	// 적록색약용 데이터 변환
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'G' {
				grid[i][j] = 'R'
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				dfs(i, j, grid[i][j], false)
				b++
			}
		}
	}

	fmt.Println(a, b)
}
