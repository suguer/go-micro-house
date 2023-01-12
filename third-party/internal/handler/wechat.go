package handler

import (
	"third-party/pkg/utils"

	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
)

type WechatService struct {
}

func NewWechatService() *WechatService {
	return &WechatService{}
}

func (*WechatService) GetQRCode() {
	qr := utils.MiniProgram.GetQRCode()
	qr.GetWXACodeUnlimit(qrcode.QRCoder{
		Scene: "123",
	})
	// response, err :=
}
