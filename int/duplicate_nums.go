package main

import "fmt"

// 배열 내에서 어떤 값이든 최소 두 번 이상 나타나면 true
// 모든 값이 서로 다르다면 false

func DuplicatedNums(nums []int) bool {
	// 정렬을 해서 순차 체크를 할 것이냐 -> 데이터 정렬에 O(nlogn)이 소요돼서 데이터가 많아지면 느릴 수 있다.
	// 맵을 하나 만들어서 빈도수를 체크할 것이냐 -> 한 번의 순회로 O(n)으로 가능, 추가 공간 O(n)이 필요
	feq := make(map[int]int, len(nums))

	for _, n := range nums {
		if _, ok := feq[n]; ok {
			return true
		}

		feq[n]++
	}

	return false
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(DuplicatedNums(nums))
}
