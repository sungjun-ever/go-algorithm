package main

func largestAltitude(gain []int) int {
	highest := 0
	sum := 0
	for _, n := range gain {
		sum += n
		highest = max(highest, sum)
	}
	return highest
}
