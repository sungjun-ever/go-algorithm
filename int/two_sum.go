package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	// 숫자를 키로 인덱스를 값으로 하는 맵 생성
	feq := make(map[int]int, len(nums))

	for idx, n := range nums {
		// 목표와 현재 수의 차가 이미 나왔다면 리턴
		if _, ok := feq[target-n]; ok {
			return []int{feq[target-n], idx}
		}

		// 맵에 넣어줌
		feq[n] = idx
	}

	return []int{}

}

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(TwoSum(nums, 6))
}
