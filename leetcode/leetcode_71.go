package main

import "strings"

func simplifyPath(path string) string {
	// 절대 경로를 canonical path로 만들어야 한다
	// .은 현재 문서
	// .. 은 부모 문서
	// N개의 슬래시는 단일 슬래시로 취급한다
	// 위 사항에 해당되지 않는 것들은 문서 또는 파일명으로 취급한다

	// 경로는 단일 슬래시로 시작해야한다
	// 경로에 있는 문서들은 슬래시로 구분된다
	// 경로는 슬래시로 끝날 수 없다

	// path를 / 기준으로 나눠준다
	paths := strings.Split(path, "/")

	// 스택을 준비해서 순차적으로 넣어준다
	stack := make([]string, 0, len(paths))
	for _, p := range paths {
		// 슬래시인 경우 넘어간다
		if p == "/" || p == "" {
			continue
		}

		// 부모 경로인 경우 현재 스택에 있는 것을 제거한다
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else if p == "." { // 현재 경로라면 넘어간다
			continue
		} else { // 그 외의 경우 스택에 넣는다
			stack = append(stack, p)
		}
	}

	return "/" + strings.Join(stack, "/")
}
