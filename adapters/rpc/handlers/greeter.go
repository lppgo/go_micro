package handlers

import (
	"context"

	proto "lianxi/microex/adapters/rpc/proto"
)

type Greeter struct{}

// 调用方法
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	// 注意 ！！！
	rsp.Greeting = "Hello " + req.Name
	return nil
}
