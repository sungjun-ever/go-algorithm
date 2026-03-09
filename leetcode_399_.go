package main

// 경로 문제로 접근
// a / b, b / c가 존재할 때
// a / c 면 a -> b -> c 로 가는 경로를 탐색하는걸로 접
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i := 0; i < len(values); i++ {
		equation := equations[i]
		val := values[i]
		x, y := equation[0], equation[1]

		if graph[x] == nil {
			graph[x] = make(map[string]float64)
		}

		if graph[y] == nil {
			graph[y] = make(map[string]float64)
		}

		graph[x][y] = val
		graph[y][x] = 1.0 / val
	}

	var dfs func(curr, target string, visited map[string]bool) float64

	dfs = func(curr, target string, visited map[string]bool) float64 {
		if curr == target {
			return 1.0
		}

		visited[curr] = true

		for neighbor, weight := range graph[curr] {
			if !visited[neighbor] {
				result := dfs(neighbor, target, visited)

				if result != -1.0 {
					return weight * result
				}
			}
		}

		return -1.0
	}

	answer := make([]float64, len(queries))
	for i, q := range queries {
		src, dst := q[0], q[1]

		if graph[src] == nil || graph[dst] == nil {
			answer[i] = -1.0
		} else {
			visited := make(map[string]bool)
			answer[i] = dfs(src, dst, visited)
		}
	}

	return answer
}
