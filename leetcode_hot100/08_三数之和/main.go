package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4} //-4 -1 -1 0 1 2
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	var rst = [][]int{}
	//先排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			//左边去重处理
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			//右边去重处理
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}

			if nums[i]+nums[l]+nums[r] > 0 {
				r--
				continue
			}

			if nums[i]+nums[l]+nums[r] < 0 {
				l++
				continue
			}

			if nums[i]+nums[l]+nums[r] == 0 {
				rst = append(rst, []int{nums[i], nums[l], nums[r]})
				l++
			}
		}
	}
	return rst
}
