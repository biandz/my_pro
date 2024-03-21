package main

import (
	"fmt"
	"time"
)

//写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出。
//package main
//func main() {
//	go func() {
//		// 1 在这里需要你写算法
//		// 2 要求每秒钟调用一次proc函数
//		// 3 要求程序不能退出
//	}()
//	select {}
//}
//func proc() {
//	panic("ok")
//}

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		t := time.NewTicker(1 * time.Second)
		for {
			<-t.C
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				proc()
			}()
		}
	}()
	select {}
}
func proc() {
	panic("ok")
}
