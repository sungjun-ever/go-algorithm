func solution(topping []int) int {
	// 롤케이크 조각의 크기와 상관 없이
	// 토핑 종류의 개수를 공평하게 가져가기
	// 공평하게 나눌수 없는 경우도 있다

	// 왼쪽과 오른쪽 유니크 토핑의 개수가 같아야한다
	// 먼저 토핑을 한쪽으로 몰아준다
	left := make(map[int]int, len(topping))
	right := make(map[int]int, len(topping))

	for _, t := range topping {
		left[t]++
	}

	// 토핑을 반대쪽으로 하나씩 옮겨준다
	ans := 0
	for i := 0; i < len(topping); i++ {
		t := topping[i]

		right[t]++
		left[t]--
		if left[t] == 0 {
			delete(left, t)
		}

		if len(left) == len(right) {
			ans++
		}
	}

	return ans
}