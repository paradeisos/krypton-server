package routers

import (
	"krypton-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &controllers.User{})
}
