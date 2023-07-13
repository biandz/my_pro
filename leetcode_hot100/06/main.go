package main

import (
	"fmt"
)

func main() {
	s1 := "ab"
	s2 := ".*c"
	fmt.Println(isMatch(s1, s2))
}
func isMatch(s string, p string) bool {
	tempP := deal(p)
	for i := 0; i < len(tempP); i++ {
		if (tempP[i] == s[0] || tempP[i] == '.') && len(tempP)-i >= len(s) {
			ss := tempP[i : i+len(s)]
			fmt.Println(ss)
			if judge(s, ss) {
				return true
			}
		}
	}
	return false
}

func deal(str string) string {
	sli := make([]uint8, 0, len(str))
	for i, v := range str {
		if i > 0 && v == '*' {
			sli = append(sli, str[i-1])
		} else {
			sli = append(sli, str[i])
		}
	}
	return string(sli)
}

func judge(s1, s2 string) bool {
	for i := 0; i < len(s2); i++ {
		if s2[i] == '.' || s2[i] == s1[i] {
			continue
		}
		return false
	}
	return true
}
