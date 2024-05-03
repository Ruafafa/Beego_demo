package main

import (
	_ "Beego_demo/routers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 设置日志输出到文件
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/beego.log"}`)
	if err != nil {
		return
	}

	// 如果需要，可以设置日志级别
	logs.SetLevel(logs.LevelInformational)
	beego.Run()
}
