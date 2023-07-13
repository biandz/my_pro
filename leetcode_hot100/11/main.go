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
	s := "2"
	fmt.Println(letterCombinations(s))
}
func letterCombinations(digits string) []string {
	var rst = []string{}
	for _, v := range digits {
		rst = combination(m[string(v)], rst)
	}
	return rst
}

func combination(s1, s2 []string) []string {
	var rst = make([]string, 0, len(s1)*len(s2))
	for _, s := range s1 {
		if len(s2) > 0 {
			for _, s3 := range s2 {
				rst = append(rst, s3+s)
			}
		} else {
			rst = append(rst, s)
		}
	}
	return rst
}
