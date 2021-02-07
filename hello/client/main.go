package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "go_micro/hello/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("Greeter.client"),
	)
	service.Init()

	greeter := pb.NewGreeterService("Greeter",service.Client())

	rsp,err := greeter.Hello(context.TODO(),&pb.HelloReq{Name:"jdx"})
	if err != nil{
		fmt.Println("call hello err:",err)
		return
	}
	fmt.Println(rsp.Greeting)
}
