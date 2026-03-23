package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// n 줄 수
	n := 0

	fmt.Fscan(reader, &n)

	grid := make([][]int, n+1)

	// 인접리스트로 만든다
	for i := 0; i < n-1; i++ {
		p, c := 0, 0
		fmt.Fscan(reader, &p, &c)
		grid[p] = append(grid[p], c)
		grid[c] = append(grid[c], p)
	}

	// 부모노드를 저장할 슬라이스
	parent := make([]int, n+1)
	var dfs func(int)
	dfs = func(curr int) {
		// 각 노드별로 연결된 노드를 순회한다
		// 연결된 노드가 루트노드인 경우는 제외한다
		for _, next := range grid[curr] {
			if parent[next] == 0 && next != 0 {
				// 다음 노드의 부모를 현재 노드로 해준다
				parent[next] = curr
				dfs(next)
			}
		}
	}

	// 첫 노드부터 시작한다.
	dfs(1)
	for i := 2; i <= n; i++ {
		fmt.Fprintln(writer, parent[i])
	}
}
