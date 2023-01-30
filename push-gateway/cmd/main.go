package main

import (
	"fmt"
	"net"
	"push-gateway/config"
	"push-gateway/discovery"
	"push-gateway/internal/handler"
	"push-gateway/internal/model"
	"push-gateway/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.InitConfig()
	model.InitDB()

	// etcd 地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()
	Node := discovery.Server{
		Name:    viper.GetString("server.domain"),
		Addr:    grpcAddress,
		Author:  "Ken",
		Version: "v1",
	}
	server := grpc.NewServer()
	defer server.Stop()
	// 绑定service
	service.RegisterSmsServiceServer(server, handler.NewSmsService())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(Node, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
