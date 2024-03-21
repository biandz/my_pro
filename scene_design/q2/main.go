package main

import (
	"fmt"
	"sync"
	"time"
)

//题目一场景：某一高并发web服务器需限制IP的频繁访问，现模拟100个IP同时并发访问该服务器，每个IP需访问1000次。
//
//题目一要求：每个IP三分钟内只能访问一次。
//
//题目一内容：修改以下代码完成上述题目要求，要求能成功输出【success100】。

//type Ban struct {
//	visitIPs map[string]time.Time
//}
//func NewBan() *Ban {
//	return &Ban{
//		visitIPs: make(map[string]time.Time),
//	}
//}
//func (o *Ban) visit(ip string) bool {
//	if _, ok := o.visitIPs[ip]; ok {
//		return true
//	}
//	o.visitIPs[ip] = time.Now()
//	return false
//}
//func main() {
//	success := 0
//	ban := NewBan()
//	for i := 0; i < 1000; i++ {
//		for j := 0; j < 100; j++ {
//			go func() {
//				ip := fmt.Sprintf("192.168.1.%d", j)
//				if !ban.visit(ip) {
//					success++
//				}
//			}()
//		}
//	}
//	fmt.Println("success:", success)
//}

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan() *Ban {
	return &Ban{
		visitIPs: make(map[string]time.Time),
	}
}
func (o *Ban) visit(ip string) bool {
	defer o.lockAndUnlock()()
	if v, ok := o.visitIPs[ip]; ok {
		currentNow := time.Now()
		if currentNow.Unix()-v.Unix() > 180 { //超过三分钟，允许访问
			o.visitIPs[ip] = currentNow
			return false
		}
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func (o *Ban) lockAndUnlock() func() {
	o.lock.Lock()
	return func() {
		o.lock.Unlock()
	}
}

func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}(j)
		}
	}
	fmt.Println("success:", success)
}
