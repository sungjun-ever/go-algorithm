package main

import "fmt"

const (
	unvisited = 0
	visiting  = 1
	visited   = 2
)

func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)

	// 인접리스트 형태로 선행과목안에 수강 과목을 넣어준다
	for _, pre := range prerequisites {
		course, precourese := pre[0], pre[1]
		adj[precourese] = append(adj[precourese], course)
	}

	state := make([]int, numCourses)

	// 하나의 선행 과목에 연결되어있는 과목들을 계속 따라간다.
	// 끝까지 따라갔으면 완료 처리
	// 따라 가는 중에 완료 처리가 안된? 수강을 만나면 불가능?

	var dfs func(curr int) bool

	dfs = func(curr int) bool {
		// 현재 수강이 아직 방문 중이라면 순환이 발생
		if state[curr] == visiting {
			return false
		}

		// 방문이 끝났으면 가능한 수강
		if state[curr] == visited {
			return true
		}

		// 현재 수강을 방문 중 처리
		state[curr] = visiting

		// 아무 처리가 되지 않는 수강 검사 시작
		for _, next := range adj[curr] {
			if !dfs(next) {
				return false
			}
		}

		// 순회가 끝나면 현재 수강 방문 완료 처리
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

func main() {
	fmt.Println(CourseSchedule(2, [][]int{{1, 0}}))
	fmt.Println(CourseSchedule(2, [][]int{{1, 0}, {0, 1}}))
}
