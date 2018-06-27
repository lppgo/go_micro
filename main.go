package main

import (
	"context"
	"fmt"

	hello "lianxi/microex/adapters/rpc/proto"

	"github.com/micro/go-grpc"
)

//-------------client-------------
func main() {

	service := grpc.NewService()
	service.Init()

	cl := hello.NewGreeterService("abcd", service.Client())

	rsp, err := cl.Hello(context.TODO(), &hello.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
