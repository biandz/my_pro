package main

import (
	"fmt"
	"strings"
)

//771. 宝石与石头
func main() {
	jewels := "z"
	stones := "ZZ"
	fmt.Println(numJewelsInStones(jewels, stones))
}

func numJewelsInStones(jewels string, stones string) int {
	var rst int
	for _, i := range jewels {
		rst += strings.Count(stones, string(i))
	}
	return rst
}
