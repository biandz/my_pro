package main

import (
	"fmt"
)

//5. 最长回文子串
func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	var maxLenStr string
	//长度为1的s，直接返回
	if len(s) == 1 {
		return s
	}

	//遍历单个长度的s肯定是回文
	for i := 0; i < len(s); i++ {
		maxLenStr = string(s[i])
	}
	//取长度为2的字串
	for l := 2; l <= len(s); l++ {
		//双指针获取字串判断是否为回文
		for i := 0; i < len(s); i++ {
			j := l + i
			//越界处理
			if j >= len(s) {
				break
			}

			//对比第一个和最后一个是否相等
			if s[i] == s[j] {
				maxLenStr = MaxLenStr(maxLenStr, s[i:j])
			}
		}
	}
	return maxLenStr
}

func MaxLenStr(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
