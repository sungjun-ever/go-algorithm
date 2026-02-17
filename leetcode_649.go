package main

func predictPartyVictory(senate string) string {
	rq := make([]int, 0, len(senate))
	dq := make([]int, 0, len(senate))
	n := len(senate)

	for i, c := range senate {
		if c == 'R' {
			rq = append(rq, i)
		} else {
			dq = append(dq, i)
		}
	}

	for len(rq) > 0 && len(dq) > 0 {
		r := rq[0]
		d := dq[0]
		rq = rq[1:]
		dq = dq[1:]

		if r < d {
			rq = append(rq, r+n)
		} else {
			dq = append(dq, d+n)
		}
	}

	if len(rq) > 0 {
		return "Radiant"
	}

	return "Dire"
}
