package main

func findCircleNum(isConnected [][]int) int {
	length := len(isConnected)
	visited := make([]bool, length)
	provinceCount := 0

	var dfs func(city int)
	dfs = func(city int) {
		visited[city] = true

		// 연결되어있고 방문안한경우
		for n := 0; n < length; n++ {
			if isConnected[city][n] == 1 && !visited[n] {
				dfs(n)
			}
		}
	}

	for i := 0; i < length; i++ {
		if !visited[i] {
			provinceCount++
			dfs(i)
		}
	}

	return provinceCount
}
