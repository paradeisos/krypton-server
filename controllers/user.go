package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"krypton-server/models"
	"net/http"
)

type User struct {
	beego.Controller
}

func (c *User) Login() {
}

// register
func (c *User) Post() {
	var params *UserRegisterParams
	resp := &Response{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		beego.Error(err)

	}

	user := models.User.NewUserModel(params.Username, params.Email, params.Password, "")
	err = user.Save()
	if err != nil {
		beego.Error(err)
	}

	resp.Status = http.StatusOK
	c.Data["json"] = resp
	c.ServeJSON()
}
