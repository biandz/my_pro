package main

import (
	"fmt"
	"sort"
)

/*
解题思路：
1、暴力：需要四重循环，时间复杂度：O(n^4)
2、两层循环+双指针（需要排序）
两种方式都需要注意去重处理
*/
func main() {
	nums := []int{2, 2, 2, 2, 2}
	target := 8
	fmt.Println(fourSum(nums, target))
}
func fourSum(nums []int, target int) [][]int {
	//先排序 -2 -1 -1 1 1 2 2
	var rst = make([][]int, 0, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		//最外层去重处理
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//用于第二层去重处理
		var flag = false
		for j := i + 1; j < len(nums)-2; j++ {
			//第二层去重处理，添加flag字段防止第一次该循环的时候去重（因为只有第二次进入该循环才判断）
			if j > 1 && nums[j] == nums[j-1] && flag {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				one := nums[i]
				two := nums[j]
				three := nums[l]
				four := nums[r]
				tempRst := one + two + three + four
				//双指针去重处理
				if l > j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r < len(nums)-1 && nums[r] == nums[r+1] {
					r--
					continue
				}
				//比对结果
				if tempRst == target {
					rst = append(rst, []int{one, two, three, four})
					l++
					continue
				}
				if tempRst > target {
					r--
				} else {
					l++
				}
			}
			flag = true
		}
	}
	return rst
}
