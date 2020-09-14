package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "go_micro/hello/proto"
)

//实现greeter.pb.micro.go GreeterHandler 接口
type GreeterServiceHandler struct {
}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *pb.HelloReq, rsp *pb.HelloRsp) error {
	rsp.Greeting = " 你好 " + req.Name
	return nil
}

func main() {
	//创建新服务
	service := micro.NewService(
		micro.Name("Greeter"),
	)
	//初始化
	service.Init()
	//注册处理器，调用 Greeter 服务接口处理请求
	if err := pb.RegisterGreeterHandler(service.Server(),new(GreeterServiceHandler));err != nil{
		fmt.Println("register err:",err)
		return
	}
	//启动服务
	//默认使用mdns,若要指定etcd 则 go run main.go --registry=etcd
	if err := service.Run(); err != nil{
		fmt.Println("run err:",err)
	}
}
