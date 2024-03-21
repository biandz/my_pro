package main

import (
	"context"
	"fmt"
	"time"
)

// 探索context的用法
func main() {
	ctx := context.TODO()
	ch := make(chan struct{})
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go func() {
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("wait timeout")
	case <-ch:
		fmt.Println("exec end!!!")
	}
}
