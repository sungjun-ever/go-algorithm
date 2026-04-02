package main

type Path struct {
	to         int  // 다음 도시
	isOriginal bool // 정방향인지 역방향인지
}

func minReorder(n int, connections [][]int) int {
	// 시작지점과 다음지점을 저장하는 slice
	graph := make([][]Path, n)

	for _, conn := range connections {
		from, to := conn[0], conn[1]

		graph[from] = append(graph[from], Path{to: to, isOriginal: true})
		graph[to] = append(graph[to], Path{to: from, isOriginal: false})
	}

	count := 0

	var dfs func(curr, parent int)

	dfs = func(curr, parent int) {
		for _, path := range graph[curr] {
			if path.to != parent {
				if path.isOriginal {
					count++
				}

				dfs(path.to, curr)
			}
		}
	}

	dfs(0, -1)

	return count
}
