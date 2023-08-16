package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	obj := new(MyQueue)
	go func() {
		for i := 0; i < 10000; i++ {
			obj.Push(i)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			pop, err := obj.Pop()
			if err != nil {
				continue
			}
			fmt.Println(pop)
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println("打印实时len：", obj.Len())
			time.Sleep(3 * time.Second)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)
	fmt.Println("start..")
	s := <-c
	fmt.Println("End...", s)
}

type MyQueue []int

func (q MyQueue) Len() int {
	return len(q)
}

func (q *MyQueue) Push(x int) {
	*q = append(*q, x)
}

func (q *MyQueue) Pop() (int, error) {
	if q.Len() == 0 {
		return 0, errors.New("empty")
	}
	old, n := *q, len(*q)
	x := old[0]
	*q = old[1:n]
	return x, nil
}
