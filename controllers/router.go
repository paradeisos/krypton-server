package controllers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/register", &User{}, `post:Register`)
	beego.Router("/user/logon", &User{}, `post:Login`)
	beego.Router("/tomato", &Tomato{})
	beego.Router("/todo", &Todo{})
}
