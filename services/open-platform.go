package services

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform"
	"power-wechat-tutorial/config"
)

var OpenPlatformApp *openPlatform.OpenPlatform

func NewOpenPlatformAppService(conf *config.Configuration) (*openPlatform.OpenPlatform, error) {

	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{conf.MiniProgram.RedisAddr},
		})
	}

	app, err := openPlatform.NewOpenPlatform(&openPlatform.UserConfig{

		AppID:  conf.OpenPlatform.AppID,
		Secret: conf.OpenPlatform.AppSecret,

		Token:  conf.OpenPlatform.MessageToken,
		AESKey: conf.OpenPlatform.MessageAesKey,

		//Log: openPlatform.Log{
		//	Level: "debug",
		//	File:  "./wechat.log",
		//},
		Cache:     cache,
		HttpDebug: false,
		Debug:     false,
		Http: openPlatform.Http{
			Timeout: 30,
		},
		//"sandbox": true,
	})

	return app, err
}
