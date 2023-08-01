package main

import (
	"fmt"
)

//1499. 满足不等式的最大值
func main() {
	bills := []int{5, 5, 5, 10, 20}
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	var rst bool
	for i, bill := range bills {
		if i == 0 && bill > 5 {
			return false
		}
	}
	return rst
}
