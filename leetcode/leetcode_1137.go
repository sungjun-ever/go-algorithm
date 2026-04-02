package main

func tribonacci(n int) int {
	f := make([]int, 38)
	f[0], f[1], f[2] = 0, 1, 1

	for i := 3; i <= 37; i++ {
		f[i] = f[i-3] + f[i-2] + f[i-1]
	}

	return f[n]
}
