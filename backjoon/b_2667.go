package main

import (
	"fmt"
	"sort"
)

func main() {
	// 총 단지 수 입력 받음
	n := 0
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)
		return
	}

	// 단지 내 집이 있는지 입력 받음
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		houseCount := ""
		if _, err := fmt.Scan(&houseCount); err != nil {
			fmt.Println(err)
			return
		}
		grid[i] = make([]int, n)

		for j := 0; j < n; j++ {
			grid[i][j] = int(houseCount[j] - '0')
		}
	}

	// 단지내 집의 수 저장용
	ans := make([]int, 0, n)

	// 4방향 배열
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		grid[x][y] = 0
		count := 1

		// 집을 찾으면 4방향으로 퍼지면서 찾는다
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 1 {
				count += dfs(nx, ny)
			}

		}

		return count
	}

	// 같은 단지에 속하는 집들끼리 구분을 지어야함
	// 먼저 집이 있는 좌표를 찾음
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 1 {
				ans = append(ans, dfs(x, y))
			}
		}
	}

	sort.Ints(ans)
	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Println(v)
	}
}
