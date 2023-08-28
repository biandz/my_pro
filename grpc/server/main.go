package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"my_grpc/pb"
	"net"
)

const (
	port = ":50000"
)

// grpc server
type server struct {
	pb.UnimplementedUserServiceServer
}

// 实现gRPC接口
func (s *server) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		User: &pb.User{
			Name: "test_user",
			Id:   int32(in.Id),
		},
	}, nil
}

// 拦截器，简单打印下日志
func LogUnaryInterceptorMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (r interface{}, err error) {
		r, err = handler(ctx, req)

		fmt.Printf("fullMethod(%s), errCode(%v)\n", info.FullMethod, err)
		return r, err
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 拦截器
	options := grpc.UnaryInterceptor(LogUnaryInterceptorMiddleware())
	s := grpc.NewServer(options)

	// 注册服务器实现
	pb.RegisterUserServiceServer(s, &server{})

	// 注册服务端反射
	reflection.Register(s)

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
