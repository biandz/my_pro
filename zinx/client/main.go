package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"my_zinx/msg"
	"my_zinx/router"
	"time"
)

func main() {
	//创建链接
	client := znet.NewClient("127.0.0.1", 8999)

	//设置创建链接成功后的钩子函数
	client.SetOnConnStart(onClientStart)

	//读取路由消息
	client.AddRouter(msg.PongID, &router.PongRouter{})

	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}

func onClientStart(conn ziface.IConnection) {
	fmt.Println("客户端链接成功！！！")
	go loopPing(conn)
}

func loopPing(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(msg.PingID, []byte("ping...ping...ping..."))
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
