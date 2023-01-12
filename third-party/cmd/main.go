package main

import (
	"third-party/config"
	"third-party/internal/handler"
	"third-party/pkg/utils"
)

func main() {
	config.InitConfig()
	utils.InitWechat()
	service := handler.NewWechatService()
	service.GetQRCode()
}
