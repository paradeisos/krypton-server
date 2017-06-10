package controllers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &User{})
	beego.Router("/tomato", &Tomato{})
}
