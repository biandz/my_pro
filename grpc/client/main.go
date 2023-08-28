package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"my_grpc/pb"
	"time"
)

var (
	conn    *grpc.ClientConn
	address = "localhost:50000"
)

func main() {
	conn = Init()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUserInfo(ctx, &pb.UserRequest{Id: 18})
	if err != nil {
		log.Fatalf("fatal: %v", err)
	}
	fmt.Println(fmt.Sprintf("打印返回结果：%v", r.User))
}

func Init() *grpc.ClientConn {
	connect, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return connect
}
