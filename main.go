package main

import (
	"github.com/astaxie/beego"
	_ "resbox_go/database"
	_ "resbox_go/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//启动服务
	beego.Run()
}
