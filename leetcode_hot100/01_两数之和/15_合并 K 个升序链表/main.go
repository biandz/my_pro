package main

import "fmt"

func main() {
	n := 4
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(num int) []string {
	m := make(map[int][]string)
	for i := 1; i <= num; i++ {
		ts := []string{}
		if i == 1 {
			ts = append(ts, "()")
		} else {
			ss := m[i-1] //上一个数字的数组集合
			tm := make(map[string]struct{})
			for _, s := range ss {
				for _, s2 := range insert(s) {
					tm[s2] = struct{}{}
				}
			}
			for s, _ := range tm {
				ts = append(ts, s)
			}
		}
		m[i] = ts
	}
	return m[num]
}

func insert(s string) []string {
	var rst []string
	for i := 0; i < len(s)+1; i++ {
		l := s[:i]
		r := s[i:]
		newStr := l + "()" + r

		if !judge(newStr, rst) {
			rst = append(rst, newStr)
		}
	}

	return rst
}

func judge(s string, s1 []string) bool {
	for _, s2 := range s1 {
		if s2 == s {
			return true
		}
	}
	return false
}
