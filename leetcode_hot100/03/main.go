package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "au"
	fmt.Println(lengthOfLongestSubstring1(s))
}

//滑动窗口解法
func lengthOfLongestSubstring1(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var maxLen int
	var tempStr string
	for _, i := range s {
		str := string(i)
		firstPos := strings.Index(tempStr, str)
		if firstPos == -1 {
			tempStr += str
		} else {
			tempStr = tempStr[firstPos+1:] + str
		}
		maxLen = max(maxLen, len(tempStr))
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//暴力超时
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if !judgeIsRepeat(s[j : j+i]) {
				maxLen = len(s[j : j+i])
				return maxLen
			}
		}
	}
	return maxLen
}
func judgeIsRepeat(s string) bool {
	m := make(map[rune]struct{})
	for _, i := range s {
		if _, ok := m[i]; ok {
			return true
		} else {
			m[i] = struct{}{}
		}
	}
	return false
}
