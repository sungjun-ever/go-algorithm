package programmers

func solution(n int, left int64, right int64) []int {
	/*
	   00 1, 01 2, 02 3, 03 4, 04 5, 05 6
	   10 2, 11 2, 12 3, 13 4, 14 5, 15 6
	   max(row,col)+1
	   row: 1   2   3   4   5   6
	   row: 2   2   3   4   5   6
	   row: 3   3   3   4   5   6
	   row: 4   4   4   4   5   6
	   row: 5   5   5   5   5   6
	   row: 6   6   6   6   6   6
	*/
	answer := make([]int, 0)

	for idx := left; idx <= right; idx++ {
		row := idx / int64(n)
		col := idx % int64(n)

		answer = append(answer, int(max(row, col)+1))
	}

	return answer
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
