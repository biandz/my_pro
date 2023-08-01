package main

import (
	"fmt"
	"strconv"
)

//874. 模拟行走机器人
func main() {
	cmd := []int{4, -1, 3}
	obstacles := [][]int{{}}
	fmt.Println(robotSim(cmd, obstacles))
}

func robotSim(commands []int, obstacles [][]int) int {
	var maxL int
	m := make(map[string]struct{})
	for _, v := range obstacles {
		if len(v) == 0 {
			continue
		}
		key := strconv.Itoa(v[0]) + "." + strconv.Itoa(v[1])
		m[key] = struct{}{}
	}

	//初始化方向为北
	var direction = 0 //0代表北，1代表东，2代表南，3代表西
	//初始化位置
	var position = []int{0, 0}
	for _, command := range commands {
		if command == -1 {
			if direction < 3 {
				direction++
			} else {
				direction = 0
			}
			continue
		}

		if command == -2 {
			if direction > 0 {
				direction--
			} else {
				direction = 3
			}
			continue
		}

		for i := 0; i < command; i++ {
			switch direction {
			case 0:
				position[1]++
				tk := strconv.Itoa(position[0]) + "." + strconv.Itoa(position[1])
				if _, ok := m[tk]; ok {
					position[1]--
					break
				}
			case 1:
				position[0]++
				tk := strconv.Itoa(position[0]) + "." + strconv.Itoa(position[1])
				if _, ok := m[tk]; ok {
					position[0]--
					break
				}
			case 2:
				position[1]--
				tk := strconv.Itoa(position[0]) + "." + strconv.Itoa(position[1])
				if _, ok := m[tk]; ok {
					position[1]++
					break
				}
			case 3:
				position[0]--
				tk := strconv.Itoa(position[0]) + "." + strconv.Itoa(position[1])
				if _, ok := m[tk]; ok {
					position[0]++
					break
				}
			}
		}
		l := position[0]*position[0] + position[1]*position[1]
		maxL = max(l, maxL)
	}
	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
