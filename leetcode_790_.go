package main

func numTilings(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n == 3 {
		return 5
	}

	prev3 := 1 // 전전전
	prev2 := 2 // 전전
	prev1 := 5 // 전

	const mod = 1000000007

	for i := 4; i <= n; i++ {
		curr := (2*prev1 + prev3) % mod
		prev3 = prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
