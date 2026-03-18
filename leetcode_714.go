package main

func maxProfit(prices []int, fee int) int {
	// H: 들고 있을 때, F: 안 들고 있을때
	// 보유중이려면
	// 안 팔면 H[i] = H[i-1]
	// 산다면: H[i] = F[i-1] - prices[i]
	// H[i] = max(H[i-1], F[i-1] - prices[i])

	// 빈손이려면
	// 안 사면 F[i] = F[i-1]
	// 팔면: F[i] = H[i-1] - prices[i] - fee
	// F[i] = max(F[i-1], H[i-1] + prices[i] - fee)
	h := 0 - prices[0] // 첫 날 주식을 들고있으려면
	f := 0             // 첫 날 빈 손이면
	n := len(prices)

	for i := 1; i < n; i++ {
		h, f = max(h, f-prices[i]), max(f, h+prices[i]-fee)
	}

	return max(h, f)

}
