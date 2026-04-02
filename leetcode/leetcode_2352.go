package main

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	rowMap := make(map[string]int)
	count := 0

	// 각 row의 수를 키로 가져감
	for i := 0; i < len(grid); i++ {
		rowKey := arrayToString(grid[i])
		rowMap[rowKey]++
	}

	for j := 0; j < len(grid); j++ {
		col := make([]int, len(grid))
		for i := 0; i < len(grid); i++ {
			col[i] = grid[i][j]
		}
		colkey := arrayToString(col)

		if val, ok := rowMap[colkey]; ok {
			count += val
		}
	}

	return count
}

func arrayToString(arr []int) string {
	var str strings.Builder

	i := 0
	for i < len(arr)-1 {
		str.WriteString(strconv.Itoa(arr[i]))
		str.WriteString(",")
		i++
	}

	str.WriteString(strconv.Itoa(arr[i]))
	return str.String()
}
