package main

func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	const (
		unvisited = 0
		visiting  = 1
		visited   = 2
	)

	// 모든 수업을 수강 가능하면 true. 아니면 false
	graph := make([][]int, numCourses)

	// 선행 과목안에 수강들을 넣어준다
	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		graph[pre] = append(graph[pre], course)
	}

	state := make([]int, numCourses)

	var dfs func(curr int) bool

	dfs = func(curr int) bool {
		if state[curr] == visiting {
			return false
		}

		if state[curr] == visited {
			return true
		}

		state[curr] = visiting

		for _, next := range graph[curr] {
			if !dfs(next) {
				return false
			}
		}

		state[curr] = visited

		return true

	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}

func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	// 선행 과목안에 수강들을 넣어준다
	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		graph[pre] = append(graph[pre], course)
		// 수강의 선수 과목 수를 중가 시켜준다
		inDegree[course]++
	}

	queue := []int{}
	// 선수과목이 없는 과정들을 큐에 넣는다
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 수강 완료 과목 수
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1 : len(queue)-1]
	}

	return count == numCourses
}
