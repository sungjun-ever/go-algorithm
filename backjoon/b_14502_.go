package main

import (
	"bufio"
	"fmt"
	"os"
)

// 연구소의 크기는 n * m, 3 <= n,m <= 8
// 바이러스는 인접한 4방향으로 퍼질 수 있다
// 0 빈칸, 1 벽, 2 바이러스
// 벽을 세웠을 때 얻을 수 있는 안전영역의 최대 크기
// 바이러스의 개수는 2 <= v <= 10
// 빈칸의 개수 3 <= b

var (
	n, m        int
	grid        [][]int
	queue       [][2]int
	maxSafeArea int
	dirs        = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &m)
	grid = make([][]int, n)
	queue = make([][2]int, 0, 10)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
			// 바이러스 좌표를 큐에 넣어줌
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	setWall(0, 0)
	fmt.Fprintln(writer, maxSafeArea)

}

func setWall(count, start int) {
	if count == 3 {
		bfs()
		return
	}

	// 1차원 인덱스로 변환
	for i := start; i < n*m; i++ {
		r, c := i/m, i%m
		if grid[r][c] == 0 {
			grid[r][c] = 1
			setWall(count+1, i+1)
			grid[r][c] = 0
		}
	}

}

func bfs() {
	// 그리드 현재상태 복사
	tempGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		tempGrid[i] = make([]int, m)
		copy(tempGrid[i], grid[i])
	}

	// 바이러스 전파
	tempQueue := make([][2]int, len(queue))
	copy(tempQueue, queue)

	for len(tempQueue) > 0 {
		curr := tempQueue[0]
		tempQueue = tempQueue[1:]

		for _, d := range dirs {
			nx, ny := curr[0]+d[0], curr[1]+d[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && tempGrid[nx][ny] == 0 {
				tempGrid[nx][ny] = 2
				tempQueue = append(tempQueue, [2]int{nx, ny})
			}
		}
	}

	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if tempGrid[i][j] == 0 {
				count++
			}
		}
	}

	maxSafeArea = max(maxSafeArea, count)
}
