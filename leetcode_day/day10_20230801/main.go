package main

import (
	"fmt"
)

//2681. 英雄的力量
func main() {
	nums := []int{2, 1, 4}
	fmt.Println(sumOfPower(nums))
}

func sumOfPower(nums []int) int {
	total := 1
	for i := 0; i < len(nums); i++ {
		total *= 2
	}
	return total
}

func genChildSet(nums []int) {

}
