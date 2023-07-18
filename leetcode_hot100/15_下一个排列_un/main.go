package main

import "fmt"

func main() { //1 2 3 5 4
	nums := []int{2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
func nextPermutation(nums []int) {
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	temp1 := nums[i+1:]
	temp2 := nums[:i+1]
	rst := []int{}
	rst = append(rst, temp2...)
	rst = append(rst, reverseSLi(temp1)...)
	for i2, i3 := range rst {
		nums[i2] = i3
	}
}
func reverseSLi(nums []int) []int {
	rst := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		rst = append(rst, nums[i])
	}
	return rst
}
