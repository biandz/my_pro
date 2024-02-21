package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"my_grpc/pb"
)

var (
	address = "localhost:50000"
)

func main() {
	conn := Init()
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	r, err := c.GetUserInfo(context.Background(), &pb.UserRequest{Id: 18})
	if err != nil {
		log.Fatalf("fatal1111111: %v", err)
	}
	fmt.Println(fmt.Sprintf("打印返回结果：%v", r.User))
}

func Init() *grpc.ClientConn {
	//证书
	c, err := credentials.NewClientTLSFromFile("../conf/custom.pem", "")
	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}
	connect, err := grpc.Dial(address, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return connect
}
