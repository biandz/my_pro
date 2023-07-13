package main

import (
	"fmt"
)

/*
解题思路：
1、暴力：需要四重循环，时间复杂度：O(n^4)
2、两层循环+双指针（需要排序）
两种方式都需要注意去重处理
*/
var m = make(map[string][]string)

func init() {
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
}
func main() {
	s := "23"
	fmt.Println(letterCombinations(s))
}
func letterCombinations(digits string) []string {
	for _, v := range digits {
		fmt.Println(m[string(v)])
	}

	return []string{}
}
