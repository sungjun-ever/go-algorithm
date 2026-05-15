package main

import "fmt"

func FindWordInGrid(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	var dfs func(int, int, int) bool

	dfs = func(x, y, index int) bool {
		if board[x][y] != word[index] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		temp := board[x][y]
		board[x][y] = '#'

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx >= 0 && ny >= 0 && nx < m && ny < n {
				if dfs(nx, ny, index+1) {
					return true
				}
			}
		}

		board[x][y] = temp
		return false
	}

	// 문자열의 시작지점을 찾는다
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(FindWordInGrid(board, "SEE"))

}
