package main

import (
	_ "krypton-server/routers"

	"github.com/astaxie/beego"
	"krypton-server/controllers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	controllers.InitEnv()
	beego.Run()
}
