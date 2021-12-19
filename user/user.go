/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 16:54:22
 * @Email: 17719495105@163.com
 */
package main

import (
	"flag"
	"fmt"

	"github.com/smartgreeting/store-rpc/user/internal/config"
	"github.com/smartgreeting/store-rpc/user/internal/server"
	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
		// 注册grpcui

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
