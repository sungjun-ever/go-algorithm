package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms), len(rooms))
	visited[0] = true

	var dfs func(room int)

	dfs = func(room int) {
		for _, k := range rooms[room] {
			if !visited[k] {
				visited[k] = true
				dfs(k)
			}
		}
	}

	dfs(0)

	for _, v := range visited {
		if v == false {
			return false
		}
	}

	return true
}
