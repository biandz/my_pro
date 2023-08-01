package main

import (
	"fmt"
)

//2208. 将数组和减半的最少操作次数
func main() {
	nums1 := []int{1, 0, 1}
	nums2 := []int{0, 0, 0}
	queries := [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}}
	fmt.Println(handleQuery1(nums1, nums2, queries))
}

func handleQuery1(nums1 []int, nums2 []int, queries [][]int) []int64 {
	var rst []int64
	var nums2TempTotal = sum(nums2)
	for _, v := range queries {
		if v[0] == 1 { //将nums1 下标v[1]到v[2]的值进行反转0->1，1->0
			dealReverse(nums1, v[1], v[2])
		} else if v[0] == 2 { //令 nums2[i] = nums2[i] + nums1[i] * p  p=v[1]
			p := v[1]
			nums2TempTotal = dealNums2(nums1, nums2, p)
		} else if v[0] == 3 {
			rst = append(rst, int64(nums2TempTotal))
		}
	}
	return rst
}

func dealReverse(nums1 []int, v1, v2 int) {
	for i := v1; i <= v2; i++ {
		nums1[i] = 1 - nums1[i]
	}
}

func dealNums2(nums1, nums2 []int, p int) int {
	var total int
	for i := 0; i < len(nums2); i++ {
		tmp := nums2[i] + nums1[i]*p
		nums2[i] = tmp
		total += tmp
	}
	return total

}

func sum(nums2 []int) int {
	var total int
	var i, j = 0, len(nums2) - 1
	for i <= j {
		total += nums2[i]
		if i != j {
			total += nums2[j]
		}
		i++
		j--
	}
	return total
}
