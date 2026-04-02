package main

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	// 1. 왼쪽 곱셈 (Prefix Product) 처리
	// result[i]에는 i의 왼쪽 숫자들의 모든 곱을 담습니다.
	result[0] = 1 // 0번 인덱스의 왼쪽엔 아무것도 없으므로 1
	for i := 1; i < length; i++ {
		fmt.Println("res = result * nums", result[i-1], nums[i-1])
		result[i] = result[i-1] * nums[i-1]
	}

	// 2. 오른쪽 곱셈 (Suffix Product) 처리
	// 이번에는 오른쪽에서 거꾸로 오면서, 아까 구한 result[i]에
	// i의 오른쪽 숫자들의 모든 곱을 곱해줍니다.
	rightMul := 1
	for i := length - 1; i >= 0; i-- {
		fmt.Println("res = result * rightMul", result[i-1], rightMul)
		result[i] = result[i] * rightMul
		rightMul *= nums[i] // 다음 칸(왼쪽 칸)을 위해 현재 숫자를 곱해둠
	}

	return result
}
