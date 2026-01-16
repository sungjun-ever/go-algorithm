package main

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	// 정렬된 문자열을 키로 가지는 map 생성
	ascMap := make(map[string][]string)

	for _, s := range strs {
		// 문자열을 문자 배열로 변환 -> 정렬 -> 재합성
		convS := strings.Split(s, "")
		sort.Strings(convS)
		joinS := strings.Join(convS, "")

		ascMap[joinS] = append(ascMap[joinS], s)
	}

	var res [][]string

	for _, ss := range ascMap {
		res = append(res, ss)
	}

	return res
}

func groupAnagrams2(strs []string) [][]string {
	// 각 문자열 알파벳 횟수 배열을 키로하는 map 생성
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		count := [26]int{}
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
		}

		// count 배열을 키로 가지도록
		groups[count] = append(groups[count], s)
	}

	var res [][]string

	for _, ss := range groups {
		res = append(res, ss)
	}

	return res
}
