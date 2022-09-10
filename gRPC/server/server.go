package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "src/code.oldboyedu.com/gRPC/proto"
)

func main() {
	//1.需要监听
	addr := "127.0.0.1:8082"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	fmt.Println("监听正常")
	//2.需要实例化grpc
	server := grpc.NewServer()
	//3.在grpc注册微服务
	pb.RegisterUserInfoServiceServer(server, &u)
	//4.启动服务端
	server.Serve(listen)

}

// UserInfoService 定义空接口
type UserInfoService struct {
}

var u = UserInfoService{}

// GetUserInfo 实现方法
func (u *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	//通过用户名查询用户信息
	name := req.Name
	//数据库里查询用户信息
	if name == "zs" {
		res = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"singing", "run"},
		}
	}
	return
}
