package main

import "fmt"

func NubmerOfIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	cnt := 0
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cnt++

				queue := [][]int{{i, j}}
				grid[i][j] = 0

				for len(queue) > 0 {
					curr := queue[0]
					queue = queue[1:]

					for k := 0; k < 4; k++ {
						nx, ny := curr[0]+dx[k], curr[1]+dy[k]

						if nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == 1 {
							grid[nx][ny] = 0
							queue = append(queue, []int{nx, ny})
						}
					}
				}
			}
		}
	}

	return cnt
}

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	fmt.Println(NubmerOfIslands(grid))
}
