package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(num1, num2 []int) float64 {
	var rst float64
	num1 = append(num1, num2...)
	sort.Ints(num1)
	if len(num1)%2 == 1 {
		rst = float64(num1[len(num1)/2])
	} else {
		rst = float64(num1[len(num1)/2]+num1[len(num1)/2-1]) / 2
	}
	return rst
}
