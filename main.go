package main

import (
	_ "go-firstweb/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//静态文件处理目录
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
