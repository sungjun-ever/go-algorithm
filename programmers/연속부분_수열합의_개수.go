package programmers

func solution(elements []int) int {
	// 여기서 말하는 길이가 N이라는 뜻은
	// N개의 연속된 수열 합이다

	// 수들의 합을 저장하는 맵을 만들어서 총 길이를 리턴한다
	sumMap := make(map[int]bool)

	// 각 인덱스를 시작 지점으로 잡고 시작 지점부터 길이를 늘려나가며 나오는 조합을 기록한다
	for start := 0; start < len(elements); start++ {
		// 길이별로 나눈다
		sum := 0
		// start부터 시작해서 길이별로 순열합을 구했을 때 나오는 수
		for length := 1; length <= len(elements); length++ {
			// 연속된 수열이기 때문에 인덱스는 돌아온다
			idx := (start + length - 1) % len(elements)
			// 1, 1+1, 1+1+4, 1+1+4+7, 1+1+4+7+9
			// 4, 4+7, 4+7+9, 4+7+9+1, 4+7+9+1+1
			sum += elements[idx]
			sumMap[sum] = true
		}
	}

	return len(sumMap)
}
