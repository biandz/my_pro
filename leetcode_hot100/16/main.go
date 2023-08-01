package main

import (
	"fmt"
	"strings"
)

func main() { //1 2 3 5 4
	s := ")))(    ()(())  (()"
	fmt.Println(longestValidParentheses1(s))
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

func longestValidParentheses1(s string) int {
	res := 0
	ln := 0
	temp := 0
	s = strings.TrimRight(s, "(")
	s = strings.TrimLeft(s, ")")
	n := 0
	carry := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ln++
			if carry == "(" {
				if n > 0 {
					temp = 0
					n = 0
				}
				n++
			}
			carry = "("
		}
		if s[i] == ')' {
			carry = ")"
			if ln != 0 {
				ln--
				temp += 2
			} else {
				temp = 0
			}
			res = max(temp, res)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
