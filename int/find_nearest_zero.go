package main

import "fmt"

// 1에서 가장 가까운 0까자의 거리를 찾는다.
func FindNearestZero(mat [][]int) {
	// 0을 기준으로 출발하거나 1을 기준으로 출발할 수 있다.
	if len(mat) == 0 || len(mat[0]) == 0 {
		return
	}

	m, n := len(mat), len(mat[0])

	queue := make([][]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else { // 1이면 방문 확인을 위해 -1로 변경한다.
				mat[i][j] = -1
			}
		}
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y := curr[0], curr[1]

		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]

			// 방문하지 않고 그리드 안에 있는 경우에만 확인
			if nx >= 0 && ny >= 0 && nx < m && ny < n && mat[nx][ny] == -1 {
				// 현재 노드 기준으로 +1 한 값이 거리
				mat[nx][ny] = mat[x][y] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}
}

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	FindNearestZero(mat)
	fmt.Println(mat)
}
