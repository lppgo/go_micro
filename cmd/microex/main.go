package main

import (
	"log"
	"net/http"

	"time"

	"lianxi/microex/adapters/rpc"

	limiter "github.com/juju/ratelimit"
	"github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	grpcserver "github.com/micro/go-plugins/server/grpc"

	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	"github.com/spf13/viper"
)

const (
	Name    = "abcd"
	Version = "1.0.0"
)

// init 初始化配置
func init() {

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("dev")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
func handleProtobufRPC(w http.ResponseWriter, r *http.Request) {

	http.Error(w, "handleProtobufRPC", http.StatusMethodNotAllowed)
	return
}

//-------------server-------------
func main() {
	srv := grpc.NewService(

		micro.Name(Name),       // 服务名称
		micro.Version(Version), // 服务版本，可以根据版本选择

		// 服务注册配置
		micro.Registry(
			registry.NewRegistry(
				registry.Addrs(viper.GetStringSlice("registry.addr")...),
				registry.Timeout(5*time.Second), // 注册服务超时时间

			),
		),
		micro.RegisterInterval(3*time.Second), // 注册时间差
		micro.RegisterTTL(6*time.Second),      // 注册有效期

		// 服务监听
		micro.Server(
			grpcserver.NewServer(
				server.Name(Name),
				server.Version(Version),
				server.Address(viper.GetString("rpc.server")),
				server.Wait(true), // 等待请求完成退出
			),
		),

		// 辅助配置
		micro.WrapHandler(
			ratelimit.NewHandlerWrapper(
				limiter.NewBucket(time.Second, 500),
				true,
			), // 限速策略
		),
	)

	// Go Web 注册
	// websrv := web.NewService(
	// 	web.Name("test.web"),
	// 	web.Version("1.0"),
	// )

	// websrv.HandleFunc("/", handleProtobufRPC)

	// if err := websrv.Init(); err != nil {
	// 	log.Fatal(err)
	// }
	// go func() {
	// 	if err := websrv.Run(); err != nil {
	// 		log.Fatal(err)
	// 	}

	// }()
	srv.Init()

	rpc.Register(srv)

	if err := srv.Run(); err != nil {
		log.Println(err)
	}

}
