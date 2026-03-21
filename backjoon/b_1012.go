package main

import "fmt"

func main() {
	// 서로 인접한 배추들의 묶음만큼 지렁이가 필요하다

	// 테스트 케이스 수
	cases := 0

	if _, err := fmt.Scan(&cases); err != nil {
		fmt.Println(err)
		return
	}

	// 케이스 별 정답을 저장할 슬라이스
	ans := make([]int, cases)

	// 배추밭 배열
	grid := [50][50]int{}

	// 4방향으로 퍼질 배열
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 케이스 만큼 루프
	loop := 0

	for loop < cases {
		// 케이스 별 루프당 배추 밭 크기 및 배추 수
		m, n, nums := 0, 0, 0
		if _, err := fmt.Scan(&m, &n, &nums); err != nil {
			fmt.Println(err)
		}

		// 배추 밭 초기화
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				grid[i][j] = 0
			}
		}

		// 배추 수 만큼 위치 받아옴
		for i := 0; i < nums; i++ {
			x, y := 0, 0
			if _, err := fmt.Scan(&x, &y); err != nil {
				fmt.Println(err)
				return
			}
			grid[x][y] = 1
		}

		var dfs func(int, int)

		dfs = func(x, y int) {
			// 탐색한 곳은 배추 없음 처리
			grid[x][y] = 0

			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]

				// 범위 안에 있고 배추가 있으면 계속 탐색
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					dfs(nx, ny)
				}
			}
		}

		// 필요한 지렁이 수
		count := 0

		// 배추가 심어져 있는 곳을 찾는다
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// 구역을 찾으면 지렁이 +1, 인접한 배추 탐색
				if grid[i][j] == 1 {
					count++
					dfs(i, j)
				}
			}
		}

		// 케이스 탐새기 끝나면 ans에 넣어줌
		ans[loop] = count
		loop++
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
