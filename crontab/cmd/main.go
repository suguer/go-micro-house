package main

import (
	"context"
	"crontab/config"
	"crontab/internal/handler"
	"crontab/internal/model"
	"crontab/internal/service"
	"crontab/routes"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user/discovery"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {

	config.InitConfig()
	model.InitDB()
	routes.InitConfig()
	go startListen() //转载路由

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
}

func startListen() {
	// etcd注册
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	arr := viper.GetStringMapString("domain")
	for ServiceName, ServicePath := range arr {
		// RPC 连接
		conn, err := RPCConnect(ctx, ServicePath, etcdRegister)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		Service := GetService(ServiceName, conn)
		routes.RData.AddService(ServiceName, Service)

	}
	c := cron.New()
	handler.InitHandle(c, ctx)
	// c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
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
