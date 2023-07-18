package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	if len(s)%2 == 1 { //是基数的肯定不满足条件
		return false
	}
	m := map[uint8]uint8{')': '(', '}': '{', ']': '['}

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
