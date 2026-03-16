package main

// bfs 방식으로 접근함
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	letterMap := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	// 큐에 조합 문자열들을 계속 넣어줌
	queue := make([]string, 0, 144)
	queue = append(queue, "")

	for _, d := range digits {
		curr := letterMap[d-'0']

		n := len(queue)
		for n > 0 {
			popped := queue[0]
			queue = queue[1:]

			for _, c := range curr {
				queue = append(queue, popped+string(c))
			}
			n--
		}
	}

	return queue
}

// dfs, backtracking 방식 (ai)
func letterCombinationsDFS(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	letterMap := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	ans := make([]string, 0, 144)

	var dfs func(index int, letter string)

	dfs = func(index int, letters string) {
		// 길이가 같아지면 종료
		if index == len(digits) {
			ans = append(ans, letters)
			return
		}

		// 번호에 해당하는 문자 묶음
		curr := letterMap[digits[index]-'0']

		for _, c := range curr {
			dfs(index+1, letters+string(c))
		}

	}

	dfs(0, "")

	return ans

}

func letterCombinationsBackTrack(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	letterMap := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	var ans []string
	// 하나의 공유 배열만 사용
	var path []byte

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			ans = append(ans, string(path))
			return
		}

		currLetters := letterMap[digits[index]-'0']

		for i := 0; i < len(currLetters); i++ {
			// 글자를 넣고
			path = append(path, currLetters[i])
			// 탐색
			backtrack(index + 1)
			// 방금 넣었던 글자를 뺌
			path = path[:len(path)-1]
		}
	}

	backtrack(0)

	return ans
}
