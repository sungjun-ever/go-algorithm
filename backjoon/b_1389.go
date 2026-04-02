package main

import (
	"bufio"
	"fmt"
	"os"
)

// 유저 수 2 <= N <= 100
// 친구 관계 수 1 <= M <= 5000
// 케빈 베이컨이 가장 작은 사람을 출력
// 여러 명일 경우에는 번호가 가장 작은 사람 출력
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int

	fmt.Fscan(reader, &n, &m)

	adj := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
	}

}
