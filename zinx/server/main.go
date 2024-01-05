package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"my_zinx/msg"
	"my_zinx/router"
)

func main() {
	server := znet.NewServer()
	server.SetOnConnStart(connBegin)
	server.SetOnConnStop(connLost)
	server.AddRouter(msg.PingID, &router.PingRouter{})
	server.Serve()
}

func connBegin(conn ziface.IConnection) {
	fmt.Println("链接建立执行！！")
}

func connLost(conn ziface.IConnection) {
	fmt.Println("链接断开执行！！")
}
