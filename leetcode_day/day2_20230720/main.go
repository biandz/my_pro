package main

import (
	"fmt"
	"math"
)

//918. 环形子数组的最大和
func main() {
	nums := []int{-2, 4, -5, 4, -5, 9, 4}
	fmt.Println(maxSubarraySumCircular(nums))
}

func maxSubarraySumCircular(nums []int) int {
	var l = len(nums)
	//m := make(map[int]int)  //k代表第i个元素，v代表前i个元素（包括第i个）的最大和
	m0 := make([]int, len(nums)) //k代表第i个元素，v代表前i个元素（包括第i个）且必须包含第0个元素的最大和
	m0[0] = nums[0]
	var max1, max2, tempMax1Total = nums[0], math.MinInt, nums[0]
	var sum3 = nums[0]
	//第一种i情况的最大值sli[len(sli)-1]
	for i := 1; i < l; i++ {
		sum3 += nums[i]
		tempMax1Total = max(nums[i], nums[i]+tempMax1Total)
		max1 = max(tempMax1Total, max1)
		m0[i] = max(sum3, m0[i-1])
	}
	//第二种情况
	tempSum := 0
	for j := l - 1; j > 0; j-- {
		tempSum += nums[j]
		max2 = max(max2, tempSum+m0[j-1])
	}
	return max(max1, max2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
