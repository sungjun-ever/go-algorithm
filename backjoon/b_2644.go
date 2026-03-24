package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1 <= n <= 100
// 입력 첫 줄은 전체 사람의 수 n
// 둘째 줄은 촌수를 계산해야하는 서로 다른 두 사람의 번호 n1, n2
// 셋째 줄에는 부모 자식들 관계의 개수 m
// 넷째 줄에는 부모 자식관의 관계를 나타내는 두 번호 x, y
// x는 y의 부모
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, n1, n2, m int

	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &n1, &n2)
	fmt.Fscan(reader, &m)

	grid := [100][100]bool{}
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		grid[x][y] = true
		grid[y][x] = true
	}

	depth := -1
	var dfs func(int, int, int)
	dfs = func(num, target, count int) {
		if num == target {
			depth = count
			return
		}

		for i := 1; i <= n; i++ {
			if grid[num][i] {
				grid[num][i] = false
				grid[i][num] = false
				dfs(i, target, count+1)
				grid[num][i] = true
				grid[i][num] = true
			}
		}
	}

	dfs(n1, n2, 0)

	fmt.Fprintln(writer, depth)
}
