package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
func maxArea(height []int) int {
	var area int //面积
	var l, r = 0, len(height) - 1
	for l < r {
		h := min(height[l], height[r])
		w := r - l
		area = max(h*w, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
