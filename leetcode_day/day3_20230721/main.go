package main

import (
	"fmt"
	"math"
)

//1499. 满足不等式的最大值
func main() {
	points := [][]int{{-19, -12}, {-13, -18}, {-12, 18}, {-11, -8}, {-8, 2}, {-7, 12}, {-5, 16}, {-3, 9}, {1, -7}, {5, -4}, {6, -20}, {10, 4}, {16, 4}, {19, -9}, {20, 19}}
	k := 6
	fmt.Println(findMaxValueOfEquation(points, k))
}

func findMaxValueOfEquation(points [][]int, k int) int {
	var maxRst = math.MinInt
	//x坐标有序，可以考虑滑动窗口
	i, j := 0, 1
	for j < len(points) {
		if i == j {
			j++
			continue
		}
		tv := abs(points[i][0], points[j][0])
		if tv <= k {
			maxRst = max(maxRst, tv+points[i][1]+points[j][1])
			j++
		} else {
			i++
		}
	}
	return maxRst
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
