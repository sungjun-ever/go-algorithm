package main

import "fmt"

// 1부터 n까지의 수로 k를 뽑아 만들 수 있는 조합
func FindCombinations(n, k int) [][]int {
	// 하나의 숫자로 만들 수 있는 조합을 확인하고 k가 되면 답에 넣고 종료
	var ans [][]int

	var backtrack func(curr int, path []int)

	backtrack = func(curr int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := curr; i <= n; i++ {
			path = append(path, i)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(1, []int{})

	return ans
}

func main() {
	fmt.Print(FindCombinations(4, 2))
}
