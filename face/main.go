package main

import (
	"fmt"
	"time"
)

type Task struct {
	ch chan *Worker
}

func (t *Task) Run() {
	for ch := range t.ch {
		ch.Run()
	}
}

func (t *Task) AddWorker(worker *Worker) {
	t.ch <- worker
}

func NewTask() *Task {
	return &Task{
		ch: make(chan *Worker, 20),
	}
}

type Worker struct {
	F func(i int)
	T int
}

func (w *Worker) Run() {
	go func() {
		timer := time.NewTimer(time.Duration(w.T) * time.Second)
		defer timer.Stop()
		<-timer.C
		w.F(w.T)
	}()
}

func main() {
	//task := NewTask()
	//go task.Run()
	////添加任务
	//for i := 1; i < 100; i++ {
	//	w := &Worker{
	//		F: t1,
	//		T: i,
	//	}
	//	task.AddWorker(w)
	//}
	//fmt.Println("任务添加完成！！！！")
	////模拟阻塞
	//for {
	//
	//}
	fmt.Println(reverse("bdzc"))
}

func t1(i int) {
	fmt.Println(fmt.Sprintf("延时 %d s执行", i))
}

func reverse(str string) string {
	b := []byte(str)
	i := 0
	j := len(b) - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
