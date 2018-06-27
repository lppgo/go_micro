package rpc

import (
	"lianxi/microex/adapters/rpc/handlers"
	proto "lianxi/microex/adapters/rpc/proto"

	"github.com/micro/go-micro"
)

// Register 注册RPC服务
func Register(srv micro.Service) {
	// TODO: do your works here.
	proto.RegisterGreeterHandler(srv.Server(), new(handlers.Greeter))
}
