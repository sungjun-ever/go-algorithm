package main

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)

	if sLen > tLen {
		return false
	}

	sp, tp := 0, 0

	for sp < sLen && tp < tLen {
		if t[tp] == s[sp] {
			sp++
		}
		tp++
	}

	return sp == sLen

}
