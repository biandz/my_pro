package main

import "fmt"

func main() { //1 2 3 5 4
	s := ")()())"
	fmt.Println(longestValidParentheses(s))
}
func longestValidParentheses(s string) int {
	var rst int
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			newStr := s[j : j+i]
			if len(newStr)%2 == 0 && judge(newStr) {
				return len(newStr)
			}
		}
	}
	return rst
}
func judge(s string) bool {
	m := map[uint8]uint8{')': '('}
	heap := []uint8{}
	for k, _ := range s {
		if m[s[k]] > 0 { //说明是右括号
			if len(heap) == 0 || heap[len(heap)-1] != m[s[k]] {
				return false
			}
			heap = heap[:len(heap)-1]
		} else { //说明是左括号
			heap = append(heap, s[k])
		}
	}
	return len(heap) == 0
}
