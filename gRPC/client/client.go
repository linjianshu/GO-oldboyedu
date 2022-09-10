package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "src/code.oldboyedu.com/gRPC/proto"
)

func main() {

	//1.连接服务器
	dial, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常 , %s \n", err)
	}
	defer dial.Close()
	//2.实例grpc客户端
	client := pb.NewUserInfoServiceClient(dial)
	//3.调用接口
	request := pb.UserRequest{Name: "zs"}
	response, err := client.GetUserInfo(context.Background(), &request)
	if err != nil {
		fmt.Println("响应异常 %s\n", err)
		return
	}
	fmt.Printf("响应结果: %v\n", response)
}
