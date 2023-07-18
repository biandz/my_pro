package main

import (
	"fmt"
)

/*
解题思路：
判断元素的idx+val是否大于等于最后一个元素的idx（前提：需要判断之前的步数是否可达当前元素）
*/
func main() {
	nums := []int{0, 2, 3}
	fmt.Println(canJump(nums))
}
func canJump(nums []int) bool {
	var maxGap int
	for i, num := range nums {
		//跳过最后一个元素的遍历，只需要遍历倒数第二个元素即可
		if i == len(nums)-1 {
			continue
		}
		//判断记录的步数是否可达当前元素
		if maxGap >= i {
			maxGap = max(maxGap, i+num)
		}
	}
	return maxGap >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
