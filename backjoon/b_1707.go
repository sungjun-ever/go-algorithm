package main

import (
	"bufio"
	"fmt"
	"os"
)

// 케이스 수 2 <= k <= 5
// 정점의 개수 1 <= v <= 20000, 간선의 개수 1 <= e <= 200000
// e만큼 간선 정보
func main() {
	reader := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()

	var k int
	fmt.Fscan(reader, &k)

	for k > 0 {
		var v, e int
		fmt.Fscan(reader, &v, &e)
		grid := make([][]int, v)
		colors := make([]int, v)

		// 두 정점 연결 표시
		for i := 0; i < e; i++ {
			var x, y int
			fmt.Fscan(reader, &x, &y)
			grid[x-1] = append(grid[x-1], y-1)
			grid[y-1] = append(grid[y-1], x-1)
		}

		isOk := true
		var dfs func(int, int)
		dfs = func(curr, color int) {
			colors[curr] = color

			for _, next := range grid[curr] {
				if colors[next] == 0 {
					dfs(next, 3-colors[curr])
				} else if colors[next] == colors[curr] {
					isOk = false
					return
				}
			}

			if !isOk {
				return
			}
		}

		for i := 0; i < v; i++ {
			if colors[i] == 0 {
				dfs(i, 1)
			}

			if !isOk {
				break
			}
		}

		if isOk {
			fmt.Fprintln(write, "YES")
		} else {
			fmt.Fprintln(write, "NO")
		}

		k--
	}
}
