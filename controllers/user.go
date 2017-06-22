package controllers

import (
	"encoding/json"
	"krypton-server/errors"
	"krypton-server/models"
	"krypton-server/utils/jwt"
	"net/http"

	"strings"

	"github.com/astaxie/beego"
)

type User struct {
	beego.Controller
}

//TODO: session
func (c *User) Login() {
	var params *UserLoginParams
	if json.Unmarshal(c.Ctx.Input.RequestBody, &params) != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	var (
		user *models.UserModel
		err  error
	)

	if strings.Contains(params.Name, "@") {
		user, err = models.User.FindByEmail(params.Name)
	} else {
		user, err = models.User.FindByUsername(params.Name)
	}
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	if user.Password != params.Password {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	token := jwt.GenToken(user.Id.Hex(), user.Username, 86400)
	cookie := http.Cookie{
		Name:   "Authorization",
		Value:  token,
		Path:   "/",
		MaxAge: 86400,
	}

	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	c.Data["json"] = Newresponse(http.StatusOK, "", user)
	c.ServeJSON()
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
