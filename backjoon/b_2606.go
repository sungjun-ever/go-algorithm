package main

import "fmt"

func main() {
	// 총 컴퓨터 수 m
	// 연결된 쌍의 수 n
	m, n := 0, 0

	if _, err := fmt.Scan(&m); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)
		return
	}

	// 컴퓨터 관계 설정 grid 만든다
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, m)
	}

	// 쌍의 수 만큼 관계를 입력 받는다
	for i := 0; i < n; i++ {
		x, y := 0, 0
		if _, err := fmt.Scan(&x, &y); err != nil {
			fmt.Println(err)
			return
		}

		grid[x-1][y-1] = 1
		grid[y-1][x-1] = 1
	}

	// 해당 컴퓨터 확인 여부용
	visited := make([]bool, m)

	useBfs(m, n, grid, visited)
}

func useBfs(m, n int, grid [][]int, visited []bool) {
	queue := make([]int, 0, m)
	queue = append(queue, 0)
	visited[0] = true
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for next := 1; next < m; next++ {
			// 연결 되어있는 경우
			if grid[curr][next] == 1 && !visited[next] {
				visited[next] = true
				grid[curr][next] = 0
				grid[next][curr] = 0
				queue = append(queue, next)
				count++
			}
		}
	}

	fmt.Println(count)
}

func useDfs(m, n int, grid [][]int, visited []bool) {

	// 오염된 컴퓨터 수 확인용
	count := 0

	var dfs func(int)
	dfs = func(curr int) {
		visited[curr] = true

		// 해당 행을 순회를 돌면서 연결 되어있고, 방문하지 않은 경우에만
		// 다음 행으로 이동
		for next := 0; next < m; next++ {
			if grid[curr][next] == 1 && !visited[next] {
				count++
				dfs(next)
			}
		}
	}

	dfs(0)
	fmt.Println(count)
}
