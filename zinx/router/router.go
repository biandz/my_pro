package router

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"my_zinx/msg"
)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) PreHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("PreHandle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
func (r *PingRouter) Handle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("Handle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	request.GetConnection().SendMsg(msg.PongID, []byte("pong...pong...pong...[FromServer]"))
}
func (r *PingRouter) PostHandle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("PostHandle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}

type PongRouter struct {
	znet.BaseRouter
}

func (r *PongRouter) Handle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("Handle: recv from sever : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
