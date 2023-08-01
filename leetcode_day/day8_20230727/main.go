package main

import (
	"fmt"
	"sort"
)

//2500. 删除每行中的最大值
func main() {
	grid := [][]int{
		{10},
	}
	fmt.Println(deleteGreatestValue(grid))
}

func deleteGreatestValue(grid [][]int) int {
	var rst int
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}
	for i := len(grid[0]) - 1; i >= 0; i-- {
		tempMax := 0
		for j := 0; j < len(grid); j++ {
			tempMax = max(tempMax, grid[j][i])
		}
		rst += tempMax
	}
	return rst
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
