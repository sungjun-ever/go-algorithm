package main

func maxProfit(prices []int, fee int) int {
	/**
	  두 가지 상태가 존재한다.

	  1. 가지고 있는 상태 h
	      1. 사거나 h[i] = f[i-1] - prices[i] 빈손인 현재 상태 값에서 금액을 빼줌
	      2. 산 상태를 유지할 때 h[i] = h[i-1] 전 상태값을 그대로 가져감
	      3. 둘 중 더 큰 값을 가져감
	  2. 빈 손인 상태 f
	      1. 팔거나 f[i] = h[i-1] + prices[i] - fee
	      2. 빈손 유지 f[i] = f[i-1]
		  3. 둘 중 더 큰 값을 가져감
	*/

	h := 0 - prices[0] // 첫 날 보유 중
	f := 0             // 첫 날 미보유중

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		h, f = max(h, f-price), max(f, h+price-fee)
	}

	return max(h, f)
}
