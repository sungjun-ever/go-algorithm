package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 상어의 기본 크기는 2
// 자신보다 작은 물고기만 먹고
// 자신보다 큰 물고기칸은 못 지나간다. 자신과 같은 사이즈는 통과는 가능하다
// 자신과 크기와 같은 수의 물고기를 먹어야 크기 + 1
// 먹을 수 있는 물고기를 다 먹으려면?

type State struct {
	x    int
	y    int
	dist int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	grid := make([][]int, n)
	var sx, sy int // 상어 위치

	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &grid[i][j])
			if grid[i][j] == 9 {
				sx, sy = i, j
				grid[i][j] = 0 // 상어 위치는 빈칸으로 처리
			}
		}
	}

	sharkSize := 2
	eatCount := 0
	totalTime := 0

	dirs := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for {
		queue := []State{{sx, sy, 0}}
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		visited[sx][sy] = true

		var eatable []State // 먹을 수 있는 물고기 후보들

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nx, ny := curr.x+d[0], curr.y+d[1]

				// 범위 안 + 방문 안 함 + 나보다 작거나 같은 물고기
				if nx >= 0 && nx < n && ny >= 0 && ny < n && !visited[nx][ny] && grid[nx][ny] <= sharkSize {
					visited[nx][ny] = true

					// 먹을 수 있는 물고기 조건: 0보다 크고 나보다 작아야 함
					if grid[nx][ny] > 0 && grid[nx][ny] < sharkSize {
						eatable = append(eatable, State{nx, ny, curr.dist + 1})
					}
					// 이동은 계속 함
					queue = append(queue, State{nx, ny, curr.dist + 1})
				}
			}
		}

		// 먹을 수 있는 물고기가 없으면 종료
		if len(eatable) == 0 {
			break
		}

		// 우선순위에 따라 정렬 (거리 -> 위 -> 왼쪽)
		sort.Slice(eatable, func(i, j int) bool {
			if eatable[i].dist != eatable[j].dist {
				return eatable[i].dist < eatable[j].dist
			}
			if eatable[i].x != eatable[j].x {
				return eatable[i].x < eatable[j].x
			}
			return eatable[i].y < eatable[j].y
		})

		// 가장 우선순위 높은 물고기 섭취
		target := eatable[0]
		totalTime += target.dist
		sx, sy = target.x, target.y // 상어 이동
		grid[sx][sy] = 0            // 먹힌 물고기 제거
		eatCount++

		// 성장 체크
		if eatCount == sharkSize {
			sharkSize++
			eatCount = 0
		}
	}

	fmt.Fprintln(writer, totalTime)

}
