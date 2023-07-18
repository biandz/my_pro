package main

import (
	"fmt"
)

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(str string) string {
	var rst string
	for i := len(str); i > 0; i-- {
		for j := 0; j <= len(str)-i; j++ {
			tempStr := str[j : i+j]
			if judgePalindrome(tempStr) {
				rst = tempStr
				return rst
			}
		}
	}
	return rst
}
func judgePalindrome(s string) bool {
	times := (len(s) + 1) / 2
	for i := 0; i < times; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
