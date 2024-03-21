package main

import (
	"fmt"
	"sync"
	"time"
)

//使sync.WaitGroup中的Wait函数支持WaitTimeout功能。

type MyWaitGroup struct {
	wg sync.WaitGroup
	t  time.Duration
}

func (m *MyWaitGroup) Wait() { //主要判断那个条件先执行
	rst := make(chan struct{})
	go func() {
		timer := time.NewTimer(m.t)
		select {
		case <-timer.C:
			rst <- struct{}{}
		}
	}()

	go func() {
		m.wg.Wait()
		rst <- struct{}{}
	}()
	<-rst
}
func (m *MyWaitGroup) Add(i int) {
	m.wg.Add(i)
}
func (m *MyWaitGroup) Done() {
	m.wg.Done()
}

func main() {
	var mw = &MyWaitGroup{
		wg: sync.WaitGroup{},
		t:  1 * time.Second,
	}
	for i := 0; i < 10; i++ {
		mw.Add(1)
		go print1(i, mw)
	}

	mw.Wait()
}

func print1(i int, mw *MyWaitGroup) {
	defer mw.Done()
	time.Sleep(2 * time.Second)
	fmt.Println(i)
}
