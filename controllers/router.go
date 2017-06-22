package controllers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &User{})
	beego.Router("/user/logon", &User{}, `post:Login`)
	beego.Router("/tomato", &Tomato{})
}
