package main

import (
	"api-gateway/config"
	"api-gateway/discovery"
	"api-gateway/internal/crontab"
	"api-gateway/internal/service"
	"api-gateway/pkg/utils"
	"api-gateway/routes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	config.InitConfig()
	crontab.InitCrontab()
	utils.InitRedis()
	go startListen() //转载路由

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	// etcd注册
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ginRouter := routes.NewRouter()
	arr := viper.GetStringMapString("domain")
	for ServiceName, ServicePath := range arr {
		fmt.Printf("ServiceName: %v\n", ServiceName)
		// RPC 连接
		conn, err := RPCConnect(ctx, ServicePath, etcdRegister)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		Service := GetService(ServiceName, conn)
		// Service := service.NewHouseServiceClient(conn)
		// wrapper.NewServiceWrapper(ServiceName)
		ginRouter.AddService(ServiceName, Service)
	}
	ginEngine := ginRouter.Start()
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足")
		fmt.Println(err)
	}
	go func() {
		// 优雅关闭
		utils.GracefullyShutdown(server)
	}()
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
}

func GetService(serviceName string, conn *grpc.ClientConn) interface{} {
	switch serviceName {
	case "user":
		return service.NewUserServiceClient(conn)
	case "house":
		return service.NewHouseServiceClient(conn)
	case "house_group":
		return service.NewHouseGroupServiceClient(conn)
	case "sms":
		return service.NewSmsServiceClient(conn)
	}
	return nil
}

func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}
