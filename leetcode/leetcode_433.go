package main

func minMutation(startGene string, endGene string, bank []string) int {
	// startGene -> endGene으로 가기 위해 한글자씩 확인이 필요하다
	// 중간 과정 돌연변이와 engGene 모두 유전자 은행에 있어야한다.

	// endGene이 bank에 없다면 진행할 필요가 없다
	bankSet := make(map[string]bool)

	for _, gene := range bank {
		bankSet[gene] = true
	}

	if !bankSet[endGene] {
		return -1
	}

	// 변환 가능한 문자
	chars := []byte{'A', 'C', 'G', 'T'}

	// 돌연변이 횟수
	mutations := 0

	// 한 번 확인한 돌연변이는 다시 확인하지 않도록 방문처리용
	visited := map[string]bool{startGene: true}

	// 시작 유전자를 큐에 넣고 시작한다
	queue := []string{startGene}

	for len(queue) > 0 {
		// 돌연 변이 횟수 구분을 위해
		// 큐 현재 상태 사이즈 만큼 루프를 돈다
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			// 큐에서 뽑은 돌연변이가 답이라면 리턴
			if curr == endGene {
				return mutations
			}

			// 현재 문자열을 byte 배열로 변환한다
			geneBytes := []byte(curr)

			for pos := 0; pos < len(geneBytes); pos++ {
				original := geneBytes[pos]

				// 변환 및 체크 과정
				for _, ch := range chars {
					// 같은 문자는 하지 않는다
					if original == ch {
						continue
					}

					geneBytes[pos] = ch
					nextGene := string(geneBytes)

					if bankSet[nextGene] && !visited[nextGene] {
						visited[nextGene] = true
						queue = append(queue, nextGene)
					}
				}

				geneBytes[pos] = original
			}
		}

		mutations++
	}

	return -1
}
