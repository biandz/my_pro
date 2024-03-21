package main

import (
	"fmt"
	"time"
)

var (
	BucketSize          = 98
	CreateFlagPerSecond = 100
	Groutines           = 100
)

// 实现令牌桶算法，来限流
func main() {
	//定义桶大小
	ch := make(chan struct{}, BucketSize)
	//启动前事先放入BucketSize个令牌
	for i := 0; i < BucketSize; i++ {
		ch <- struct{}{}
	}

	//启动一个协程，定时往桶里放入令牌
	go func() {
		//这里模拟1s中只允许10个令牌产生
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			for i := 0; i < CreateFlagPerSecond; i++ {
				ch <- struct{}{}
			}
		}
	}()
	//启动一个协程来模拟大并发请求
	go func() {
		for i := 0; i < Groutines; i++ {
			go print_test(i, ch)
		}
	}()
	select {}
}

func print_test(i int, ch chan struct{}) {
	timer := time.NewTimer(500 * time.Millisecond) //模拟等待1s，1s没拿到令牌就返回被限流的逻辑
	defer timer.Stop()
	select {
	case <-ch:
		fmt.Println(fmt.Sprintf("我【协程ID:%d】拿到令牌桶啦，开始接下来的任务。。。", i))
	case <-timer.C:
		fmt.Println(fmt.Sprintf("我【协程ID:%d】被限流了", i))
	}
}
