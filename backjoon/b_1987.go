package main

import (
	"bufio"
	"fmt"
	"os"
)

// 세로 r, 가로 c
// 보드 각 판에는 대문자 알파벳
// 0,0 에는 말이 있음
// 말은 4방향 이동 가능하고 지나가는 모든 알파벳은 유니크 해야한다
// 1 <= r, c <= 20
// 말이 최대 몇칸을 갈 수 있는지 구한다
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var r, c int
	fmt.Fscan(reader, &r, &c)

	grid := [20][20]byte{}
	for i := 0; i < r; i++ {
		row := ""
		fmt.Fscan(reader, &row)
		for j := 0; j < c; j++ {
			grid[i][j] = row[j]
		}
	}

	alphaMap := map[byte]bool{}

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= r || y < 0 || y >= c || alphaMap[grid[x][y]] {
			return 0
		}

		alphaMap[grid[x][y]] = true
		count := max(dfs(x+1, y), dfs(x-1, y), dfs(x, y+1), dfs(x, y-1)) + 1
		alphaMap[grid[x][y]] = false
		return count
	}

	fmt.Fprintln(writer, dfs(0, 0))
}
