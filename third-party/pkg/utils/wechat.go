package utils

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/spf13/viper"
)

var Wechater *wechat.Wechat
var MiniProgram *miniprogram.MiniProgram

func InitWechat() *wechat.Wechat {
	Wechater = wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     viper.GetString("wechat.AppID"),
		AppSecret: viper.GetString("wechat.AppSecret"),
		Cache:     memory,
	}
	MiniProgram = Wechater.GetMiniProgram(cfg)
	// Wechater.GetOfficialAccount(&config.Config{}).GetAccessToken()
	return Wechater
}
