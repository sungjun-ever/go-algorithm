package main

// 3 x 3 박스 안에 1~9  중복으로 들어가면 안된다
// 한 행과 열에 1~9가 중복으로 들어가면 안된다
func isValidSudoku(board [][]byte) bool {
	var rows [9][9]bool  // 9줄의 행의 숫자 출현 여부 기록용
	var cols [9][9]bool  // 9줄의 열의 숫자 출현 여부 기록용
	var boxes [9][9]bool // 33 박스 숫자 출현 여부 기록용

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			// 빈 공간이면 검사할 필요가 없음
			if board[r][c] == '.' {
				continue
			}

			// board 셀의 숫자를 인덱스로 변환
			idx := board[r][c] - '1' // 0~8

			// 박스
			// 각 박스는 3으로 나눈 크기
			// 그 박스에 숫자가 나온지 확인
			// 아래 공식은 2차원 좌표를 1차원 좌표로 표현하는 공식
			boxIdx := (r/3)*3 + (c / 3)

			// 하나라도 이미 출현한 숫자라면 종료
			if rows[r][idx] || cols[c][idx] || boxes[boxIdx][idx] {
				return false
			}

			rows[r][idx] = true
			cols[c][idx] = true
			boxes[boxIdx][idx] = true
		}
	}
	return true
}
