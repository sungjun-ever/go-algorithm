package main

func uniqueOccurrences(arr []int) bool {
	// 모든 숫자의 중복 수가 유니크한 경우 true, 반대면 false
	// 숫자의 빈도수를 체크하는 map 하나
	// 빈도수가 key인 map 하나
	cnt := make(map[int]int)
	feq := make(map[int]int)

	for _, v := range arr {
		cnt[v]++
	}

	for _, v := range cnt {
		if feq[v] == 0 {
			feq[v] = 1
		} else {
			return false
		}
	}

	return true

}
