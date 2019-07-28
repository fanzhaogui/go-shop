package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

import (
	_ "shop/router"
)

type ShopConf struct {
	redisAddr string
}

var shopConf = &ShopConf{}

func initAppConf() {
	redisAddr := beego.AppConfig.String("redis_addr")
	logs.Debug("redis addr is : ", redisAddr)
	shopConf.redisAddr = redisAddr
}

// 关于直接运用go run 运行项目，会报错的问题，即使配置文件设置了 autorendor=false
// 请先 go build，然后在运行
//
func main() {
	// init app config
	initAppConf()

	// run application
	beego.Run()
}
