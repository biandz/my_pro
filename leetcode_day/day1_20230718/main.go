package main

import (
	"fmt"
	"math"
)

//1851. 包含每个查询的最小区间
func main() {
	intervals := [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}
	queries := []int{2, 3, 4, 5}
	fmt.Println(minInterval(intervals, queries))
}

func minInterval(intervals [][]int, queries []int) []int {
	var rst = make([]int, 0, len(queries))
	for i := 0; i < len(queries); i++ {
		flag := false //表示有没有在某一个区间
		temp := math.MaxInt
		for j := 0; j < len(intervals); j++ {
			l := intervals[j][0]
			r := intervals[j][1]
			if l <= queries[i] && queries[i] <= r { //表示在区间
				//计算区间值
				temp = min(temp, r-l+1)
				flag = true
			}
		}
		if !flag {
			rst = append(rst, -1)
		} else {
			rst = append(rst, temp)
		}
	}
	return rst
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
