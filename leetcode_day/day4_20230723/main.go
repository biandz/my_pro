package main

import (
	"fmt"
)

//771. 宝石与石头
func main() {
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0}
	height := []int{4, 2, 0, 3, 2, 5, 0}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	var rst int
	si := 0        //开始下标
	ei := 0        //结束下标
	isAsc := false //true代表升，false代表降
	for i := 0; i < len(height); i++ {
		if i == 0 {
			si = i
			continue
		}

		if height[i] >= height[i-1] {
			ei = i
			isAsc = true
		} else { //遇到单调降区间需要更新前一个v字形雨水空间和下一个v的开始坐标
			if isAsc == true {
				rst += (ei-si-1)*min(height[ei], height[si]) - sum(ei, si, height)

				si = ei
			}
			isAsc = false
		}
	}

	return rst
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(ei, si int, nums []int) int {
	ts := nums[si+1 : ei]
	total := 0
	for _, t := range ts {
		total += t
	}
	return total
}
