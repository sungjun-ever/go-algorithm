package main

import (
	"bufio"
	"fmt"
	"os"
)

// n개의 줄에 m개의 정수로 미로
// 2 <= n, m <= 100
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 이동할 수 있는 칸 저장용
	grid := make([][]byte, n)

	// 칸별 이동 수 저장용
	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		grid[i] = make([]byte, m)
		dist[i] = make([]int, m)
		var row string
		fmt.Fscan(reader, &row)
		for j := 0; j < m; j++ {
			grid[i][j] = row[j] - '0'
		}
	}

	// 4방향 정의
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := make([][2]int, 0, n*m)
	// 가장 첫 자리 넣어줌
	queue = append(queue, [2]int{0, 0})
	// 첫 자리 이동 수 1 초기화
	dist[0][0] = 1

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// 목적지 도착 시 종료
		if curr[0] == n-1 && curr[1] == m-1 {
			break
		}

		for _, d := range dirs {
			nx, ny := curr[0]+d[0], curr[1]+d[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == 1 && dist[nx][ny] == 0 {
				dist[nx][ny] = dist[curr[0]][curr[1]] + 1
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}

	fmt.Fprintln(writer, dist[n-1][m-1])
}
