package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 서로 연결된 덩어리의 수를 구함
	reader := bufio.NewReader(os.Stdin)

	var n, m int // n: 정점의 개수, m: 간선의 개수
	fmt.Fscan(reader, &n, &m)

	// 연결을 기록할 grid
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		x, y := 0, 0
		fmt.Fscan(reader, &x, &y)

		grid[x-1][y-1] = 1
		grid[y-1][x-1] = 1
	}

	// 덩어리 수 기록
	count := 0

	// 방문 여부 체크
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(curr int) {
		visited[curr] = true

		for next := 0; next < n; next++ {
			if grid[curr][next] == 1 && !visited[next] {
				dfs(next)
			}
		}
	}

	for i := 0; i < n; i++ {
		// 현재 정점을 방문하지 않았다면 카운트
		if !visited[i] {
			count++
			dfs(i)
		}
	}

	fmt.Println(count)
}
