package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum1(nums, target))
}

//暴力破解O(n)*O(n)
func twoSum1(nums []int, target int) []int {
	var rst = make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				rst = append(rst, i, j)
				return rst
			}
		}
	}
	return rst
}

//hash解放O(n)
func twoSum2(nums []int, target int) []int {
	rst := make([]int, 0, 2)
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[num]; ok {
			rst = append(rst, idx, i)
		} else {
			m[target-num] = i
		}
	}
	return rst
}
