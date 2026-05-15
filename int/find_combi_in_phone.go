package main

import "fmt"

func FindCombiInPhone(digits string) []string {
	var ans []string

	size := len(digits)

	if size == 0 {
		return []string{}
	}

	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(index int, path []byte)
	backtrack = func(index int, path []byte) {
		if len(path) == size {
			ans = append(ans, string(path))
			return
		}

		num := digits[index]
		strs := digitMap[num]

		for i := 0; i < len(strs); i++ {
			path = append(path, strs[i])
			// 다음 숫자를 backtrack에 넣는다
			backtrack(index+1, path)
			path = path[:len(path)-1]

		}

	}

	backtrack(0, []byte{})

	return ans
}

func main() {
	fmt.Println(FindCombiInPhone("23"))
}
